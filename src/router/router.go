package router

import (
    "api"
    "github.com/labstack/echo"
    "api/middleware"
)

func New() *echo.Echo{
    e := echo.New()

    // add some group
    adminGroup := e.Group("/admin")
    cookieGroup := e.Group("/cookie")
    jwtGroup := e.Group("/jwt")

    // add MainMiddleware
    middleware.SetMainMiddleWare(e)
    middleware.SetAdminMiddleware(adminGroup)
    middleware.SetCookieMiddleware(cookieGroup)
    middleware.SetJwtMiddleware(jwtGroup)

    // set mainGroup router
    api.SetMainGroup(e)

    // set Group router
    api.SetAdminGroup(adminGroup)
    api.SetCookieGroup(cookieGroup)
    api.SetJwtGroup(jwtGroup)

    return e

}
