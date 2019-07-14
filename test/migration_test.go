package test

import (
	"go-nfe-repeater/arquivei"
	"go-nfe-repeater/database"
	"go-nfe-repeater/nfe"
	"testing"
)

func testMigration(t *testing.T) {
	db := database.NewEmbeddedDatabaseConnection()
	repo := nfe.NewNfeRepository(db)
	handler := arquivei.NewDatabaseMigrationHandler(repo)

	handler.MigrateData()

	nfeData := arquivei.RetrieveNfeData()

	for _, data := range nfeData {
		nfeEntry, err := repo.FindByAccessKey(data.AccessKey)

		if err != nil {
			t.Fatalf("An error occured while retrieving nfe from the repository: %s", err)
		}

		if data.XmlValue != nfeEntry.XmlValue {
			t.Errorf("Migration is not synchroning data properly")
		}
	}
}