// Package main
package main

import (
	"context"
	"go-micro.dev/v4"
	"istio/greeter/proto/hello"
	"log"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *go_micro_srv_greeter.Request, rsp *go_micro_srv_greeter.Response) error {
	log.Println("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	go_micro_srv_greeter.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
