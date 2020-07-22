package app

import (
	"applyGo/internal/ping"
	"applyGo/internal/student"
	"net/http"

	"github.com/labstack/echo/v4"
)

type route struct {
	Group      string
	Path       string
	Method     string
	Endpoint   echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

func InitRoute(e *echo.Echo, cv *Configs) error {

	pingEndpoint := ping.NewEndpoint()
	studentEndpoint := student.NewEndpoint()
	routes := []route{
		route{
			Group:      "",
			Path:       "/health_check",
			Method:     http.MethodGet,
			Endpoint:   pingEndpoint.HealthCheck,
			Middleware: []echo.MiddlewareFunc{cv.PingNotFound},
			//Middleware: nil,
		},
		route{
			Group:      "student",
			Path:       "/number",
			Method:     http.MethodGet,
			Endpoint:   studentEndpoint.NumberOfStudent,
			Middleware: nil,
		},
		route{
			Group:      "student",
			Path:       "/calgrade",
			Method:     http.MethodGet,
			Endpoint:   studentEndpoint.CalulateGrade,
			Middleware: nil,
		},
	}

	for _, rt := range routes {
		e.Group(rt.Group).Add(rt.Method, rt.Path, rt.Endpoint, rt.Middleware...)
	}
	return nil
}
