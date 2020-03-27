package api

import (
    "github.com/labstack/echo"
    "net/http"
)

func SetCookieGroup(g *echo.Group){
    g.GET("/main", mainCookie)
}

func mainCookie(c echo.Context) error {
    return c.String(http.StatusOK, "your are on the secret cookie page")
}
