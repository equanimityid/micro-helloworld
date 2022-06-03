package main

import (
	"github.com/equanimityid/micro-helloworld/handler"
	"github.com/equanimityid/micro-helloworld/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	helloworld "github.com/equanimityid/micro-helloworld/proto"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.helloworld"),
		micro.Version("0.1.0"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	helloworld.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.svc.helloworld", service.Server(), new(subscriber.Helloworld))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
