package internal

import (
	"fmt"
	"github.com/vovanada/microservices-test/app"
	portPb "github.com/vovanada/microservices-test/app/services/port/pb/port"
	"github.com/vovanada/microservices-test/app/services/port/persistence/repository"
	"google.golang.org/grpc"
	"log"
	"net"
)

type portService struct {
	portRepository repository.PortRepository
}

func NewPortService(portRepository repository.PortRepository) app.Service {
	return &portService{
		portRepository: portRepository,
	}
}

func (rcv *portService) Start(addr string) error {
	return rcv.start(addr)
}

func (rcv *portService) start(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen address [%s]: %s", addr, err)
	}

	grpcServer := grpc.NewServer()
	portPb.RegisterPortDomainServiceServer(grpcServer, rcv)

	log.Printf("Start grpc service: %s", addr)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to start: %v", err)
	}

	return nil
}
