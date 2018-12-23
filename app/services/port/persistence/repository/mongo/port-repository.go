package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/vovanada/microservices-test/app/services/port/persistence/model"
	"github.com/vovanada/microservices-test/app/services/port/persistence/repository"
)

const portsCollectionName = "ports"

type PortRepository struct {
	db *mgo.Database
}

func NewPortRepository(db *mgo.Database) repository.PortRepository {
	return &PortRepository{db: db}
}

func (r *PortRepository) Save(model *model.Port) (*model.Port, error) {
	_, err := r.db.C(portsCollectionName).Upsert(bson.M{"_id": model.PortID}, model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
func (r *PortRepository) GetWithPagination(limit, page int) ([]*model.Port, int, error) {
	var (
		models []*model.Port
		count  int
	)

	query := r.db.C(portsCollectionName).Find(bson.M{})

	count, err := query.Count()

	if err != nil {
		return models, count, err
	}

	if count == 0 {
		return models, count, err
	}

	err = query.Skip((page - 1) * limit).Limit(limit).All(&models)
	if err != nil {
		return models, count, err
	}

	return models, count, nil
}
