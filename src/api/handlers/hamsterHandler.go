package handlers

import (
    "github.com/labstack/echo"
    "net/http"
    "log"
)

type Hamster struct {
    Name string `json:"name"`
    Type string `json:"type"`
}

func AddHamster(c echo.Context) error {
    // 对于body的处理 ，方法1
    hamster := Hamster{}

    err := c.Bind(&hamster)
    if err != nil {
        log.Printf("failed processint addhamster request %s", err)
        return echo.NewHTTPError(http.StatusInternalServerError)
    }
    log.Printf("this is your hamster :%s", hamster)
    return c.String(http.StatusOK, "we got your hamster")
}
