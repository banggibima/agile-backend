package middleware

import (
	"net/http"

	"github.com/banggibima/backend-agile/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type LoggerMiddleware struct {
	Config *config.Config
	Logger *logrus.Logger
}

func NewLoggerMiddleware(
	config *config.Config,
	logger *logrus.Logger,
) LoggerMiddleware {
	return LoggerMiddleware{
		Config: config,
		Logger: logger,
	}
}

func (l LoggerMiddleware) WithConfig() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:  true,
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			method := values.Method
			uri := values.URI
			status := values.Status
			latency := values.Latency.Nanoseconds()

			switch {
			case status >= http.StatusOK && status < http.StatusMultipleChoices:
				l.Logger.Infof("request: method=%s uri=%s status=%d latency=%dns", method, uri, status, latency)
			case status >= http.StatusMultipleChoices && status < http.StatusBadRequest:
				l.Logger.Infof("request [redirect]: method=%s uri=%s status=%d latency=%dns", method, uri, status, latency)
			case status >= http.StatusBadRequest && status < http.StatusInternalServerError:
				l.Logger.Warnf("client error: method=%s uri=%s status=%d latency=%dns", method, uri, status, latency)
			case status >= http.StatusInternalServerError:
				l.Logger.Errorf("server error: method=%s uri=%s status=%d latency=%dns", method, uri, status, latency)
			default:
				l.Logger.Infof("request: method=%s uri=%s status=%d latency=%dns", method, uri, status, latency)
			}
			return nil
		},
	})
}
