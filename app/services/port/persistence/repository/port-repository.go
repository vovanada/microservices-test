package repository

import (
	"github.com/vovanada/microservices-test/app/services/port/persistence/model"
)

type PortRepository interface {
	Save(*model.Port) (*model.Port, error)
	GetWithPagination(limit, page int) (ports []*model.Port, total int, err error)
}
