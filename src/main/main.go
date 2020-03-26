package main

import (
    "github.com/labstack/echo"
    "net/http"
)

func hello(c echo.Context) error {
    return c.String(http.StatusOK, "hello")
}

func main() {
    e := echo.New()
    e.GET("/", hello)

    e.Start(":1323")
}
