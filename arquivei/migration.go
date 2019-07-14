package arquivei

import (
	"go-nfe-repeater/nfe"
	"log"
)

type DatabaseMigrationHandler struct {
	Repository nfe.INfeRepository
}

func NewDatabaseMigrationHandler(repository nfe.INfeRepository) DatabaseMigrationHandler {
	return DatabaseMigrationHandler{Repository: repository}
}

func (handler DatabaseMigrationHandler) MigrateData() {
	data := RetrieveNfeData()
	for _ , nfe := range data {
		err := handler.Repository.CreateOrUpdate(nfe)
		if err != nil {
			log.Fatal(err)
		}
	}
}