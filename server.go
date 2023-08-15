package main

import (
	"github.com/kyoruni/echo-crud/route"
)

func main() {
	e := route.Route()
	e.Logger.Fatal(e.Start(":8080"))
}
