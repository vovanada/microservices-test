package internal

import (
	"context"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/vovanada/microservices-test/app/gateways/port/internal/config"
	"github.com/vovanada/microservices-test/app/gateways/port/representation"
	"github.com/vovanada/microservices-test/app/services/port/pb/port"
	"github.com/vovanada/microservices-test/shared/gateway/errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type clientAPI struct {
	apiContainer *restful.Container
	config       config.Contract
	portClient   port.PortDomainServiceClient
}

func NewPortGateway(config config.Contract, portClient port.PortDomainServiceClient) *clientAPI {
	gateway := &clientAPI{
		apiContainer: restful.NewContainer(),
		config:       config,
		portClient:   portClient,
	}

	gateway.initRestful()

	return gateway
}

func (rcv *clientAPI) initRestful() {
	ws := new(restful.WebService)

	ws.Path("/api").
		Consumes(restful.MIME_JSON, "multipart/form-data").
		Produces(restful.MIME_JSON).Doc("Ports API").
		HeaderParameter("Access-Control-Allow-Origin", "*")

	ws.Route(ws.POST("/ports/import").
		To(rcv.handlerImportPorts).
		Doc("Import ports").
		Param(ws.FormParameter("file", "JSON file").DataType("file")).
		Returns(http.StatusNoContent, "OK", nil).
		Returns(http.StatusBadRequest, "Bad request", errors.Error{}).
		Returns(http.StatusInternalServerError, "Internal error", errors.Error{}),
	)

	ws.Route(ws.GET("/ports").
		To(rcv.handlerGetPorts).
		Doc("Get ports with pagination").
		Param(ws.QueryParameter("limit", "Limit").DefaultValue("10")).
		Param(ws.QueryParameter("page", "Page number").DefaultValue("1")).
		Returns(http.StatusOK, "OK", representation.PortList{}).
		Returns(http.StatusBadRequest, "Bad request", errors.Error{}).
		Returns(http.StatusInternalServerError, "Internal error", errors.Error{}),
	)

	rcv.apiContainer.Add(ws)

	swaggerConfig := restfulspec.Config{
		WebServices: rcv.apiContainer.RegisteredWebServices(),
		APIPath:     "/apidocs.json",
		DisableCORS: true,
	}
	rcv.apiContainer.Add(restfulspec.NewOpenAPIService(swaggerConfig))
}

func (rcv *clientAPI) Start(addr string) error {
	return rcv.start(addr)
}

func (rcv *clientAPI) start(addr string) error {
	server := http.Server{
		Addr:    addr,
		Handler: rcv.apiContainer,
	}

	go func() {
		log.Printf("Start api server: %s", addr)
		err := server.ListenAndServe()

		if err != nil {
			log.Printf("Server stopped: %s", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Printf("Server interrupted")

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	return server.Shutdown(ctx)
}
