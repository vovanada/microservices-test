package internal

import (
	"context"
	"github.com/vovanada/microservices-test/app/services/port/pb/port"
)

func (rcv *portService) GetPorts(ctx context.Context, in *port.GetPortsRequest) (*port.GetPortsResponse, error) {

	ports, total, err := rcv.portRepository.GetWithPagination(int(in.GetLimit()), int(in.GetPage()))

	if err != nil {
		return nil, err
	}

	resp := &port.GetPortsResponse{
		Items: make([]*port.Port, len(ports)),
		Total: int64(total),
	}

	for k := range ports {
		resp.Items[k] = &port.Port{
			Name:        ports[k].Name,
			Country:     ports[k].Country,
			Alias:       ports[k].Alias,
			City:        ports[k].City,
			Code:        ports[k].Code,
			Coordinates: ports[k].Coordinates,
			Province:    ports[k].Province,
			Regions:     ports[k].Regions,
			Timezone:    ports[k].Timezone,
			Unlocs:      ports[k].Unlocs,
			PortID:      ports[k].PortID,
		}
	}

	return resp, nil
}
