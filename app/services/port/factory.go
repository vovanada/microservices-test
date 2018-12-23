package port

import (
	"github.com/globalsign/mgo"
	"github.com/vovanada/microservices-test/app"
	"github.com/vovanada/microservices-test/app/services/port/internal"
	"github.com/vovanada/microservices-test/app/services/port/internal/config"
	"github.com/vovanada/microservices-test/app/services/port/persistence/repository/mongo"
)

const (
	ServiceName = "port-service"
	MongoDB     = "port"
)

func Factory() (app.Service, error) {

	c, err := config.InitConfig()

	if err != nil {
		return nil, err
	}

	session, err := mgo.Dial(c.MongoURL())
	if err != nil {
		return nil, err
	}

	portRepository := mongo.NewPortRepository(session.DB(MongoDB))

	return internal.NewPortService(portRepository), nil
}
