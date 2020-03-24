package main

import (
	"fmt"
	"go-crawler/controller"
	"go-crawler/schema"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	argsWithProg := os.Args
	e := echo.New()

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &schema.CustomContext{Context: c}
			return h(cc)
		}
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Match([]string{echo.POST, echo.GET}, "/getContent", controller.GetContent)
	fmt.Println("argsWithProg", argsWithProg)
	if len(argsWithProg) > 1 {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", argsWithProg[1])))
	} else {
		e.Logger.Fatal(e.Start(":1323"))
	}

}
