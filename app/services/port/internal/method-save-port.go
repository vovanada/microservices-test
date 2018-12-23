package internal

import (
	"context"
	"github.com/vovanada/microservices-test/app/services/port/pb/port"
	"github.com/vovanada/microservices-test/app/services/port/persistence/model"
)

func (rcv *portService) SavePort(ctx context.Context, in *port.SavePortRequest) (*port.SavePortResponse, error) {

	item := in.GetItem()

	_, err := rcv.portRepository.Save(&model.Port{
		PortID:      item.GetPortID(),
		Unlocs:      item.GetUnlocs(),
		Timezone:    item.GetTimezone(),
		Regions:     item.GetRegions(),
		Province:    item.GetProvince(),
		Coordinates: item.GetCoordinates(),
		Code:        item.GetCode(),
		City:        item.GetCity(),
		Alias:       item.GetAlias(),
		Country:     item.GetCountry(),
		Name:        item.GetName(),
	})

	if err != nil {
		return nil, err
	}

	return &port.SavePortResponse{}, nil
}
