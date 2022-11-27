package main

import (
	"exval/pkg"
	postgresPersistence "exval/pkg/persistence/postgres"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	pkg.SetupConfig()

	err := execute()
	if err != nil {
		log.Panic(err)
	}
	os.Exit(0)
}

func execute() error {
	options, err := pkg.ParseOptions()
	if err != nil {
		return err
	}

	db, err := gorm.Open(postgres.Open(options.PostgresConnectionString))

	if err != nil {
		return err
	}

	err = postgresPersistence.AutoMigration(db)
	if err != nil {
		return err
	}

	return nil
}
