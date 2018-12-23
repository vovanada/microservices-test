package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/vovanada/microservices-test/app/gateways/port/representation"
	portPb "github.com/vovanada/microservices-test/app/services/port/pb/port"
	"github.com/vovanada/microservices-test/shared/gateway/errors"
	"io"
	"log"
	"net/http"
)

func (rcv *clientAPI) handlerImportPorts(req *restful.Request, resp *restful.Response) {
	var err error

	file, _, err := req.Request.FormFile(fileParamName)

	if err != nil {
		errors.BadRequest(resp)
		return
	}

	ports := make(chan *representation.Port)
	done := make(chan struct{})

	go func() {
		err = rcv.streamPortLoad(file, ports, done)
		if err != nil {
			log.Printf("Error while stream port loading, err: %s", err)
		}
	}()

	go rcv.savePort(ports, done)

	<-done

	if err != nil {
		errors.BadRequest(resp)
		return
	}

	resp.WriteHeader(http.StatusNoContent)
}

func (rcv *clientAPI) streamPortLoad(r io.Reader, ports chan *representation.Port, done chan struct{}) error {
	defer func() {
		close(done)
	}()

	dec := json.NewDecoder(r)

	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if _, ok := t.(json.Delim); ok {
			continue
		}

		if dec.More() {
			row := &representation.Port{}
			key, ok := t.(string)
			if !ok {
				return fmt.Errorf("invalid JSON")
			}

			if err := dec.Decode(&row); err != nil {
				return fmt.Errorf("error while decoding port: %s", err)
			}

			row.PortID = key

			ports <- row
		}
	}

	return nil
}

func (rcv *clientAPI) savePort(ports chan *representation.Port, done chan struct{}) {
	for {
		select {
		case port, ok := <-ports:
			if !ok {
				return
			}

			portPbItem := &portPb.Port{
				PortID:      port.PortID,
				Name:        port.Name,
				Country:     port.Country,
				Alias:       port.Alias,
				City:        port.City,
				Code:        port.Code,
				Coordinates: port.Coordinates,
				Province:    port.Province,
				Regions:     port.Regions,
				Timezone:    port.Timezone,
				Unlocs:      port.Unlocs,
			}

			_, err := rcv.portClient.SavePort(context.Background(), &portPb.SavePortRequest{
				Item: portPbItem,
			})

			if err != nil {
				log.Printf("Failed to save port [%s], err: %s", portPbItem.PortID, err)
			}
		case <-done:
			return
		}

	}
}
