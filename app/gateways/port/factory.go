package port

import (
	"github.com/vovanada/microservices-test/app"
	"github.com/vovanada/microservices-test/app/gateways/port/internal"
	"github.com/vovanada/microservices-test/app/gateways/port/internal/config"
	"github.com/vovanada/microservices-test/app/services/port/pb/port"
	"google.golang.org/grpc"
)

const ServiceName = "port-gateway"

func Factory() (app.Service, error) {

	c, err := config.InitConfig()

	if err != nil {
		return nil, err
	}

	s, err := getPortServiceClient(c.PortServiceAddr())

	if err != nil {
		return nil, err
	}

	return internal.NewPortGateway(c, s), nil
}

func getPortServiceClient(addr string) (port.PortDomainServiceClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return port.NewPortDomainServiceClient(conn), nil
}
