package test

import (
	"encoding/json"
	"go-nfe-repeater/arquivei"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

func testConsumerDatabaseSynchronization(t *testing.T) {
	client := &http.Client{}
	consumerData := arquivei.RetrieveNfeData()

	for _, data := range consumerData {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/nfe", nil)

		query := req.URL.Query()
		query.Add("key", data.AccessKey)
		req.URL.RawQuery = query.Encode()

		response, err := client.Do(req)

		if err != nil || response.StatusCode != 200 {
			t.Fatalf("The Controller is not answering successfully")
		}

		var responseData map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&responseData)

		if err != nil {
			t.Fatalf("Error while unmarshalling service response: %s", err)
		}

		if responseData["xml"] != data.XmlValue {
			t.Fatalf("Database and Service are not synchonized")
		}

		response.Body.Close()
	}
}

func testControllerMissingParameterErrorHandling(t *testing.T) {
	client := &http.Client{}

	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/nfe", nil)

	response, err := client.Do(request)

	if err != nil {
		t.Fatalf("The controller is throwing an unexpected error")
	}

	var responseData map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseData)

	if err != nil {
		t.Fatalf("Error while unmarshalling the error response: %s", err)
	}

	errorMessage := responseData["error"].(string)

	if response.StatusCode != 400 || !strings.Contains(errorMessage, "query parameter") {
		t.Fatalf("The controller is not handling the error correctly")
	}
}

func testControllerNfeNotFoundErrorHandling(t *testing.T) {
	client := &http.Client{}
	rand.Seed(time.Now().UnixNano())
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/nfe", nil)

	query := request.URL.Query()
	query.Add("key", strconv.Itoa(rand.Intn(1000)))
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)

	if err != nil {
		t.Fatalf("The controller is throwing an unexpected error")
	}

	var responseData map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseData)

	if err != nil {
		t.Fatalf("Error while unmarshalling the error response: %s", err)
	}

	errorMessage := responseData["error"].(string)

	if response.StatusCode != 400 || !strings.Contains(errorMessage, "The nfe requested was not found") {
		t.Fatalf("The controller is not handling the error correctly")
	}
}
