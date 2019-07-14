package configuration

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Arquivei map[string]string `json:"arquivei"`
	Database map[string]string `json:"database"`
}


var configuration Configuration

func LoadConfiguration() {
	file, _ := os.Open("configuration.json")
	defer file.Close()
	err := json.NewDecoder(file).Decode(&configuration)

	if err != nil {
		log.Fatalf("Error while loading configuration: %s\n", err)
	}
}

func GetConfiguration() Configuration {
	return configuration
}
