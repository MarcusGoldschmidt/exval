package main

import (
	"context"
	"exval/cmd/api/docs"
	"exval/pkg"
	"exval/pkg/api"
	"exval/pkg/controllers"
	postgresPersistence "exval/pkg/persistence/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	"os"
)

// @title           Exval API
// @version         1.0
// @description     Exval is a system that allows you to evaluate expressions
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func main() {
	pkg.SetupConfig()

	err := execute()
	if err != nil {
		log.Panic(err)
	}
	os.Exit(0)
}

func execute() error {
	ctx := context.Background()

	options, err := pkg.ParseOptions()
	if err != nil {
		return err
	}

	log.Infof(options.PostgresConnectionString)

	db, err := gorm.Open(postgres.Open(options.PostgresConnectionString))

	if options.AutoMigration {
		err = postgresPersistence.AutoMigration(db)
		if err != nil {
			return err
		}
	}

	expressionRepository := postgresPersistence.NewExpressionRepositoryImpl(db)

	expressionController := controllers.NewExpressionController(expressionRepository)

	exvalApi := api.NewExvalApi(options, expressionController, docs.SwaggerInfo)

	go func() {
		log.Infof("Starting server on %s:%d", options.Address, options.Port)
		err := exvalApi.Start(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	err = exvalApi.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
