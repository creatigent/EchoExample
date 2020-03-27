package api

import (
    "github.com/labstack/echo"
    "net/http"
)

func SetAdminGroup(g *echo.Group){
    g.GET("/main", mainAdmin)
}


func mainAdmin(c echo.Context) error {
    return c.String(http.StatusOK, "you are on the main admin page")
}