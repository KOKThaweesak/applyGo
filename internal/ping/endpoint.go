package ping

import "github.com/labstack/echo/v4"

type Endpoint struct{}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (e Endpoint) HealthCheck(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{"api": "success"})
}
