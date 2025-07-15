package api

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// DefaultStructuredLogger logs a gin HTTP request in JSON format. Uses the
// default logger from rs/zerolog.
func DefaultStructuredLogger() fiber.Handler {
	return StructuredLogger(&log.Logger)
}

// StructuredLogger logs a gin HTTP request in JSON format. Allows to set the
// logger for testing purposes.
func StructuredLogger(logger *zerolog.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()
		path := c.Path()
		raw := c.Queries()
		c.Next()

		timeStamp := time.Now()
		latency := timeStamp.Sub(start)
		if latency > time.Minute {
			latency = latency.Truncate(time.Second)
		}

		method := c.Method()
		statusCode := c.Response().StatusCode()
		if len(raw) != 0 {
			for k, v := range raw {
				path = path + "?" + k + "=" + v
			}
		}

		var logEvent *zerolog.Event
		if c.Response().StatusCode() >= 500 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}

		logEvent.Str("protocol", string(c.Request().Header.Protocol())).
			Str("method", method).
			Int("status_code", statusCode).
			Str("path", path).
			Str("duration", latency.String()).
			Msg("Request Recieved")
		return nil
	}
}
