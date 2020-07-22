package student

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoint struct{ srv *Service }

func NewEndpoint() *Endpoint {
	return &Endpoint{
		srv: Newservice(),
	}
}

func (e Endpoint) NumberOfStudent(c echo.Context) error {
	num, err := e.srv.NumberOfStudents("1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"api": ""})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"count": num})
}

type GradeInput struct {
	Grade int `query:"grade"`
}

func (e Endpoint) CalulateGrade(c echo.Context) error {
	grade := &GradeInput{}
	if err := c.Bind(grade); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{})
	}

	gradeString, err := e.srv.CalulateGrade(grade.Grade)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"api": ""})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"grade": gradeString})
}
