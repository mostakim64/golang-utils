package monitor

import "github.com/labstack/echo-contrib/prometheus"

func NewEchoPrometheusClient() *prometheus.Prometheus{
	return prometheus.NewPrometheus("klikit", nil)
}
