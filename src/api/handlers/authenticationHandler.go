package handlers

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo"
    "log"
    "net/http"
    "time"
)

type JwtClaims struct {
    Name string `json:"name"`
    jwt.StandardClaims
}

func Login(c echo.Context) error {
    username := c.QueryParam("username")
    password := c.QueryParam("password")

    if username == "jack" && password == "1234" {
        cookie := &http.Cookie{}

        cookie.Name = "sessionID"
        cookie.Value = "some_string"
        cookie.Expires = time.Now().Add(48 * time.Hour)

        c.SetCookie(cookie)

        token, err := createJwtToken()
        if err != nil {
            log.Println("error creating JWT token", err)
            return c.String(http.StatusInternalServerError, "something went wrong")
        }

        return c.JSON(http.StatusOK, map[string]string{
            "message": "you were loggen  in!",
            "token": token,
        })
    }
    return c.String(http.StatusUnauthorized, "your user name or password was wrong")
}

func createJwtToken() (string, error ){
    claims := JwtClaims{
        Name: "jack",
        StandardClaims:jwt.StandardClaims{
            ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
            Id:        "main_user_id",
        },
    }

    rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

    token, err := rawToken.SignedString([]byte("mySecret"))
    if err != nil {
        return "", nil
    }
    return token, nil
}