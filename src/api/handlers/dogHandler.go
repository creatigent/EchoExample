package handlers

import (
    "encoding/json"
    "github.com/labstack/echo"
    "net/http"
    "log"
)

type Dog struct {
    Name string `json:"name"`
    Type string `json:"type"`
}


func AddDogs(c echo.Context) error {
    // 对于body的处理 ，方法1
    dog := Dog{}
    defer c.Request().Body.Close()

    err := json.NewDecoder(c.Request().Body).Decode(&dog)
    if err != nil {
        log.Printf("failed processint addDog request %s", err)
        return echo.NewHTTPError(http.StatusInternalServerError)
    }

    log.Printf("this is your dog %s", dog)
    return c.String(http.StatusOK, "we got your dog")
}
