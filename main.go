package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloWorld struct {
	SubType  string                   `json:"subtype"`
	Children []map[string]interface{} `json:"children"`
}

func main() {
	e := echo.New()

	e.GET("/hello-world", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface {
		}{
			"_children": []map[string]interface{}{
				{
					"subtype":     "Label",
					"text":        "Hello world",
					"translation": "[0, 0]",
				},
				{
					"subtype":     "Label",
					"text":        "How are you doing",
					"translation": "[0, 20]",
				},
			},
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
