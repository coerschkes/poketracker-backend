package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type LoggerMiddleware struct {
	LogHeaders bool
}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{false}
}

func (i *LoggerMiddleware) Chain(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("REQUEST: %v %v\n", c.Request().Method, c.Request().RequestURI)
		if i.LogHeaders {
			log.Printf("HEADERS: %v\n", i.mapHeadersToString(c.Request().Header))
		}
		return next(c)
	}
}

func (i *LoggerMiddleware) mapHeadersToString(headers http.Header) string {
	headersString := "["
	for header := range headers {
		headersString = headersString + header + ": " + headers.Get(header) + " | "
	}
	headersString = headersString[:len(headersString)-2]
	headersString = headersString + "]"
	return headersString
}
