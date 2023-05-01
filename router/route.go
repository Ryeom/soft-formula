package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

func Initialize(e *echo.Echo) {
	apis := e.Group("/soft-formula")
	{
		route(apis)
	}
}

func route(g *echo.Group) {
	g.GET("/health-check", healthCheck)
}

func healthCheck(c echo.Context) error {
	result := HttpResult{
		Code:    200,
		Message: "I'm still alive.",
	}

	j, err := json.Marshal(c.Request().Header)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(j))
	rawBody, _ := ioutil.ReadAll(c.Request().Body)
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))

	fmt.Println(string(rawBody))
	fmt.Println(string(rawBody))
	return c.JSON(http.StatusOK, result)
}

type HttpResult struct {
	Code    interface{} `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
