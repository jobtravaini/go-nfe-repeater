package server

import (
	"github.com/gin-gonic/gin"
	"go-nfe-repeater/arquivei"
	"go-nfe-repeater/configuration"
	"go-nfe-repeater/database"
	"go-nfe-repeater/nfe"
	"net/http"
)

func ServerStatus(c *gin.Context) {
	c.Status(http.StatusOK)
}

func Init() {
	router := gin.Default()
	controller, migration := initializeComponents()
	router.GET("/", ServerStatus)
	router.GET("/nfe", controller.RetrieveNfe)

	migration.MigrateData()

	router.Run()
}

func initializeComponents() (nfe.NfeController, arquivei.DatabaseMigrationHandler) {
	configuration.LoadConfiguration()

	db := database.NewDatabaseConnection()
	nfeRepository := nfe.NewNfeRepository(db)
	nfeService := nfe.NewNfeService(nfeRepository)
	nfeController := nfe.NewNfeController(nfeService)

	migrationHandler := arquivei.NewDatabaseMigrationHandler(nfeRepository)

	return nfeController, migrationHandler
}