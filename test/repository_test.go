package test

import (
	"go-nfe-repeater/database"
	"go-nfe-repeater/nfe"
	"testing"
)

func testDatabasePersistence(t *testing.T) {
	db := database.NewEmbeddedDatabaseConnection()
	repository := nfe.NewNfeRepository(db)

	nfeEntry := nfe.Nfe{AccessKey:"1", XmlValue:"<value>1</value>"}

	err := repository.CreateOrUpdate(nfeEntry)

	if err != nil {
		t.Fatalf("Error while persisting Nfe into the database")
	}

	nfeResult, err := repository.FindByAccessKey(nfeEntry.AccessKey)

	if err != nil {
		t.Fatalf("Error while retrieving Nfe from the database")
	}

	if nfeEntry.XmlValue != nfeResult.XmlValue || nfeEntry.AccessKey != nfeResult.AccessKey {
		t.Fatalf("The database is generating unsynchronized data")
	}
}