
# Shadowchef Utils

Some common utils functions and log functions have been written here for common usage on other services. 

## Installation
Configure `.gitconfig` to include the below code
```
[url "git@bitbucket.org:"]
        insteadOf = https://bitbucket.org/

# this will enable the use of ssh version of the repo url instead of https.
```

To setup private bitbucket cloud
```
git config --global url."git@bitbucket.org:shadowchef".insteadOf "https://bitbucket.org/shadowchef"
```

Install Utils with `go get`.

```bash
GOPRIVATE=bitbucket.org/shadowchef/utils go get bitbucket.org/shadowchef/utils
```

## Usage
### Methods package
```go
package main
import (
	"fmt"

	"github.com/klikit/utils/methods"
)

func main() {
	fmt.Println(methods.GenerateRandomStringOfLength(20))
}
```

### Logger package
```go
package main
import (
	"fmt"

	"github.com/klikit/utils/logger"
)

func main() {
	// FOR sending Error/Panic/Fatal Log to slack channel
	webhookUrl := "webhook url"
	service := "service name"
	logger.SetSlackLogger(webhookUrl, service)
	logger.Error("Error occurred ", e.Error(), nil)
	logger.Info("put your message here...")
}
```

### Generic logger package
```go
package main

import "github.com/klikit/utils/logger"

func main() {
	kLogger := logger.NewLoggerClient()
	kLogger.Info("log some info...")
}

```

### slackit package
```go
package main

import (
	"github.com/klikit/utils/slackit"
	"fmt"
)

func main() {
	webhookUrl := "https://hooks.slack.com/services/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"
	header := "Alert"
	summary := "Alert Summary"
	details := "Details"
	serviceName := "Storage"
	metadata := "metadata"
	mentions := []string{"@here", "@there"}
	
	slackitClient := slackit.NewSlackitClient(webhookUrl)

	clientReq := slackit.ClientRequest{
		Header: header,
		ServiceName: serviceName,
		Summary: summary,
		Metadata: metadata,
		Details: details,
		Status: slackit.Alert,
		Mentions:    mentions,
	}

	err := slackitClient.Send(clientReq)
	if err != nil {
		fmt.Print("Alert occurred sending message to slack: ", err)
	}
}
```

### monitor package

```go
package main

import (
	"github.com/klikit/utils/monitor"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	monitor.NewEchoPrometheusClient(e, "/metrics")
}
```

### redisutil package
```go
package main

import (
	"github.com/klikit/utils/redisutil"
	"fmt"
)

var redissutil *redisutil.Redis
//This method should be called once at the bootstrapping of service
func ConnectRedis() {
	host := "127.0.0.1"
	port := "6379"
	pass := ""
	db := 1
	prefix := "map:"
	redissutil = redisutil.Connect(host, port, pass, db, prefix)
}

func Redis() *redisutil.Redis {
	return redissutil
}
func main() {
	ConnectRedis()
	err := Redis().Set("test_key", "test_value", 60)
	if err != nil {
		fmt.Println("Failed to set redis value")
    }
	value, err := Redis().Get("test_key")
	if err != nil {
		fmt.Println("Failed to get redis value")
    }
	fmt.Println("Value found: ", value)
}

```


To run tests, run the following command

```bash
go run main.go
```

