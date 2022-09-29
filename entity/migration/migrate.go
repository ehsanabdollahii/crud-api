package migration

import (
	"crud-api/entity"
	"crud-api/services"
	log "github.com/sirupsen/logrus"
)

func MigrateDB() error {
	db := services.GetOrmService()

	log.Debug("migrating database ...")

	err := db.AutoMigrate(
		&entity.City{},
	)

	if err != nil {
		log.Infof("migration failed: %s", err.Error())
		return err
	}

	log.Info("migration successful")

	return nil
}
