package test

import (
	"go-nfe-repeater/server"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestIntegratedSuite(t *testing.T) {
	go server.Init()
	waitServerStartup()

	t.Run("Consumer and Database Synchronization test", testConsumerDatabaseSynchronization)
	t.Run("Controller empty query parameter error handling validation", testControllerMissingParameterErrorHandling)
	t.Run("Controller nfe not found error handling validation", testControllerNfeNotFoundErrorHandling)
	t.Run("Migration validation", testMigration)
}

func TestUnitSuite(t *testing.T) {
	t.Run("Testing Adapter mapping", testAdapterMapping)
	t.Run("Testing repository", testDatabasePersistence)
}

func waitServerStartup() {
	for {
		time.Sleep(time.Second)
		resp, err := http.Get("http://localhost:8080")

		if err != nil {
			log.Println("Server is STARTING: ", err)
			continue
		}

		resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Println("Server is STARTING: ", resp.StatusCode)
			continue
		}

		break
	}
	log.Println("Server is RUNNING")
}