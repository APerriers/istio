package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
	"istio/greeter/proto/hello"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong v2")
	})

	r.GET("/hello", func(c *gin.Context) {
		// create a new service
		service := micro.NewService()

		// parse command line flags
		service.Init()

		// Use the generated client stub
		cl := go_micro_srv_greeter.NewSayService("go.micro.srv.greeter", service.Client())

		// Make request
		rsp, err := cl.Hello(context.Background(), &go_micro_srv_greeter.Request{
			Name: "John",
		})
		if err != nil {
			fmt.Println(err)
			c.String(200, err.Error())
			return
		}

		c.String(200, rsp.Msg)
	})

	log.Fatal(r.Run(":8080"))
}
