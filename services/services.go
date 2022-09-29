package services

import (
	"crud-api/services/registry"
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

var container di.Container

func SetupServices(services ...*di.Def) {
	builder, _ := di.NewBuilder()

	for _, service := range services {
		err := builder.Add(*service)
		if err != nil {
			panic("error on add definition to the container")
		}
	}

	container = builder.Build()
}

func GetOrmService() *gorm.DB {
	return container.Get(registry.OrmServiceDefinition).(*gorm.DB)
}
