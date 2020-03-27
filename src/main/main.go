package main

import (
    "router"
)

func main() {
    e := router.New()

    e.Start(":1323")
}
