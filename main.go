package main

import (
	"github.com/ebelanja/micro-examples/helloworld/handler"
	"github.com/ebelanja/micro-examples/helloworld/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	helloworld "github.com/ebelanja/micro-examples/helloworld/proto"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.helloworld"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	helloworld.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.helloworld", service.Server(), new(subscriber.Helloworld))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
