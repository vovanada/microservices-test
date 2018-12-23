package main

import (
	"flag"
	"github.com/vovanada/microservices-test/app"
	portGateway "github.com/vovanada/microservices-test/app/gateways/port"
	portService "github.com/vovanada/microservices-test/app/services/port"
	"log"
)

var services = map[string]func() (app.Service, error){}

func main() {
	var kind, addr string
	flag.StringVar(&kind, "kind", "", "Service name")
	flag.StringVar(&addr, "addr", ":8080", "Service address")

	flag.Parse()

	//add services

	services[portGateway.ServiceName] = portGateway.Factory
	services[portService.ServiceName] = portService.Factory

	//get service

	serviceFactory, ok := services[kind]

	if !ok {
		log.Fatalf("Service [%s] not found", kind)
	}

	service, err := serviceFactory()

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(service.Start(addr))
}
