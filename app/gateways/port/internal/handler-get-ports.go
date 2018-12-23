package internal

import (
	"github.com/emicklei/go-restful"
	"github.com/vovanada/microservices-test/app/gateways/port/representation"
	"github.com/vovanada/microservices-test/app/services/port/pb/port"
	"github.com/vovanada/microservices-test/shared/gateway/errors"
	"github.com/vovanada/microservices-test/shared/gateway/helpers"
	"log"
)

func (rcv *clientAPI) handlerGetPorts(req *restful.Request, resp *restful.Response) {
	limit, err := helpers.GetIntParam(req, limitQueryParamName, defaultLimit)

	if err != nil {
		errors.BadRequest(resp)
		return
	}

	page, err := helpers.GetIntParam(req, pageQueryParamName, defaultPage)

	if err != nil {
		errors.BadRequest(resp)
		return
	}

	portsServiceResp, err := rcv.portClient.GetPorts(req.Request.Context(), &port.GetPortsRequest{
		Limit: limit,
		Page:  page,
	})

	if err != nil {
		errors.Internal(resp, err)
		return
	}

	portsResp := representation.PortList{
		Items: make([]*representation.Port, len(portsServiceResp.GetItems())),
		Total: portsServiceResp.GetTotal(),
	}

	for k := range portsServiceResp.GetItems() {
		portsResp.Items[k] = &representation.Port{
			PortID:      portsServiceResp.GetItems()[k].GetPortID(),
			Unlocs:      portsServiceResp.GetItems()[k].GetUnlocs(),
			Timezone:    portsServiceResp.GetItems()[k].GetTimezone(),
			Regions:     portsServiceResp.GetItems()[k].GetRegions(),
			Province:    portsServiceResp.GetItems()[k].GetProvince(),
			Coordinates: portsServiceResp.GetItems()[k].GetCoordinates(),
			Code:        portsServiceResp.GetItems()[k].GetCode(),
			City:        portsServiceResp.GetItems()[k].GetCity(),
			Alias:       portsServiceResp.GetItems()[k].GetAlias(),
			Country:     portsServiceResp.GetItems()[k].GetCountry(),
			Name:        portsServiceResp.GetItems()[k].GetName(),
		}
	}

	err = resp.WriteEntity(portsResp)

	if err != nil {
		log.Printf("Error while writing entity to response: %s", err)
	}
}
