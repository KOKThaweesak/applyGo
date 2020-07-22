package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Configs) PingNotFound(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		if c.EnablePing {
			return next(context)
		}
		return context.JSON(http.StatusNotFound, map[string]interface{}{"api": "not found"})
	}
}
