package middleware

import (
    "github.com/labstack/echo"
    "log"
    "net/http"
    "strings"
)

func SetCookieMiddleware(g *echo.Group){
    g.Use(checkCookie)
}


func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cookie, err := c.Cookie("sessionID")

        if err != nil {
            log.Println(err)
            if strings.Contains(err.Error(), "http: named cookie not present"){
                return c.String(http.StatusUnauthorized, "you don't have any cookid")
            }
            return err
        }
        if cookie.Value == "some_string"{
            return next(c)
        }
        return c.String(http.StatusUnauthorized, "your don't have the right cookie")
    }
}