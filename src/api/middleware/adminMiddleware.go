package middleware

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)



func SetAdminMiddleware(g *echo.Group) {
    g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Skipper:          nil,
        Format:           `${time_rfc3339}  ${status}  ${method} ${host}  ${path}  ${latency_human}` + "\n",
        CustomTimeFormat: "",
        Output:           nil,
    }))

    g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (b bool, err error) {
        if username == "jack" && password == "1234"{
            return true, nil
        }
        return false, nil
    }))
}