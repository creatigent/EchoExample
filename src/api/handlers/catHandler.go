package handlers

import (
    "encoding/json"
    "fmt"
    "github.com/labstack/echo"
    "io/ioutil"
    "net/http"
    "log"
)

type Cat struct {
    Name string `json:"name"`
    Type string `json:"type"`
}

func GetCats(c echo.Context) error {
    /**
      @Des: 对于参数的处理
    */
    catName := c.QueryParam("name")
    catType := c.QueryParam("type")

    data := c.Param("data")

    //return c.String(http.StatusOK, fmt.Sprintf("catName is %s and catType is %s, data is %s !", catName, catType, data))

    if data == "string" {
        return c.String(http.StatusOK, fmt.Sprintf("catName is %s and catType is %s, data is %s !", catName, catType, data))
    }

    if data == "json" {
        // 返回JSON数据
        return c.JSON(http.StatusOK, map[string]string{
            "name":catName,
            "type":catType,
        })
    }

    return c.JSON(http.StatusBadRequest, map[string]string{
        "error":"you need let's us know if you want json or string!",
    })
}

func AddCats(c echo.Context) error {
    // 对于body的处理 ，方法1
    cat := Cat{}

    defer c.Request().Body.Close()

    b, err := ioutil.ReadAll(c.Request().Body)
    if err != nil {
        log.Printf("failed reading the request body: %s", err)
        return c.String(http.StatusInternalServerError, "")
    }

    err = json.Unmarshal(b, &cat)
    if err != nil {
        log.Printf("failed unmarshaling in addCats : %s", err)
        return c.String(http.StatusInternalServerError, "")
    }

    log.Printf("this is your cat: %s", cat)
    return c.String(http.StatusOK, "we got you cat!")
}
