package monitor

import "github.com/labstack/echo-contrib/prometheus"

func NewMonitorClient() *prometheus.Prometheus{
	return prometheus.NewPrometheus("klikit", nil)
}
