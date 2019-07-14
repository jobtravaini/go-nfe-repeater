package arquivei

import (
	"encoding/json"
	"go-nfe-repeater/configuration"
	"go-nfe-repeater/nfe"
	"log"
	"net/http"
)

type NfeData struct {
	AccessKey string `json:"access_key"`
	Xml string `json:"xml"`
}

type StatusData struct {
	Code uint8 `json:"code"`
	Message string `json:"message"`
}

type NfeReceivedResponse struct {
	Status StatusData `json:"status"`
	NfeDataArray []NfeData `json:"data"`
	Body json.RawMessage
}

func RetrieveNfeData() []nfe.Nfe {
	request, _ := http.NewRequest(http.MethodGet, configuration.GetConfiguration().Arquivei["api-url"], nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-api-id", configuration.GetConfiguration().Arquivei["api-id"])
	request.Header.Set("x-api-key", configuration.GetConfiguration().Arquivei["api-key"])

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Printf("An error occurred while calling nfe_received api: %s\n", err)
	}

	defer response.Body.Close()

	var data NfeReceivedResponse

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Println(err)
	}

	nfeArray := make([]nfe.Nfe, len(data.NfeDataArray))

	for i, nfeData := range data.NfeDataArray  {
		nfeArray[i] = MapNfe(nfeData)
	}

	return nfeArray
}