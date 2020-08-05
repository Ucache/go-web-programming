// Package classification Helloworld API.
//
//
// Terms Of Service:
//
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 1.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
//
// swagger:meta

package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Req struct {
	Name string `json:"name"`
}

// swagger:parameters idofHelloworld
type ReqWapper struct {
	//name
	//in: body
	Body Req
}

type Resp struct {
	Msg string `json:"msg"`
}

// swagger:response Resp
type RespWapper struct {
	//in: body
	Body Resp
}

// swagger:route GET /helloworld idofHelloworld
// responses:
//    200: Resp
func Helloworld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", Helloworld)
	e.Logger.Fatal(e.Start(":1323"))
}
