package echo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"bitbucket.org/shadowchef/utils/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

// TokenIdentifier function
type TokenIdentifier func(c echo.Context) (string, error)

// ByEmailToken allows email as token bucket identifier
func ByEmailToken(c echo.Context) (string, error) {
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		logger.Error(err)
		return "", errors.New("failed to parse request body")
	}

	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var body emailReqData
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		logger.Error(err)
		return "", errors.New("failed to parse request body")
	}

	if err := body.Validate(); err != nil {
		return "", errors.New("invalid email")
	}
	return body.Email, nil
}

// ByRemoteIPToken allows RemoteIP as token bucket identifier
func ByRemoteIPToken(c echo.Context) (string, error) {
	return c.RealIP(), nil
}

// RateLimiter, a echo middleware to rate limiting an endpoint
func RateLimiter(tokenIdentifier TokenIdentifier, unit string, ratePerUnit int) echo.MiddlewareFunc {
	rateLimiter := configureRateLimiterStore(unit, ratePerUnit)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			identifier, err := tokenIdentifier(c)
			if err != nil {
				logger.Error(err)
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"data": err.Error()})
			}
			isAllowed, err := rateLimiter.Allow(identifier)
			if err != nil {
				logger.Error(err)
				return c.JSON(http.StatusInternalServerError, err)
			}
			if !isAllowed {
				logger.Info(fmt.Sprintf("route: [%s] is rate limited for user identifier: [%s]", c.Request().URL.Path, identifier))
				return c.JSON(http.StatusTooManyRequests, map[string]interface{}{"data": "Too many requests"})
			}
			return next(c)
		}
	}
}

func configureRateLimiterStore(unit string, ratePerUnit int) middleware.RateLimiterStore {
	if ratePerUnit == 0 {
		logger.Warn("ratePerUnit can't be zero")
		ratePerUnit = 1
	}
	return middleware.NewRateLimiterMemoryStoreWithConfig(
		middleware.RateLimiterMemoryStoreConfig{
			Rate:      calculateRateLimit(unit, ratePerUnit),
			Burst:     ratePerUnit,
			ExpiresIn: calculateRateLimiterCleanupTime(unit),
		},
	)
}

func calculateRateLimiterCleanupTime(unit string) time.Duration {
	switch strings.ToUpper(unit) {
	case timeUnitMinute:
		return 1 * time.Minute
	case timeUnitSecond:
		return 1 * time.Hour
	default:
		return 1 * time.Second
	}
}

func calculateRateLimit(unit string, ratePerUnit int) rate.Limit {
	switch strings.ToUpper(unit) {
	case timeUnitMinute:
		return rate.Every(1 * time.Minute / time.Duration(ratePerUnit))
	case timeUnitHour:
		return rate.Every(1 * time.Hour / time.Duration(ratePerUnit))
	default:
		return rate.Every(1 * time.Second / time.Duration(ratePerUnit))
	}
}
