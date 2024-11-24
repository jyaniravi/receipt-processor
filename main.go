package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"receipt-processor/points"
	"receipt-processor/types"
)

var (
	receiptData = make(map[uuid.UUID]types.Receipt)
	mutex       sync.Mutex
)

func main() {
	muxHandler := mux.NewRouter()

	muxHandler.HandleFunc("/receipts/process", processReceiptsHandler).Methods("POST")
	muxHandler.HandleFunc("/receipts/{id}/points", getPointsHandler).Methods("GET")

	log.Println("!!!!! Starting Receipt-Processor !!!!!")

	log.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", muxHandler))

}

func processReceiptsHandler(response http.ResponseWriter, request *http.Request) {

	var receipt types.Receipt

	err := json.NewDecoder(request.Body).Decode(&receipt)
	if err != nil {
		http.Error(response, "Invalid receipt posted", http.StatusBadRequest)
	}

	id := uuid.New()
	mutex.Lock()
	receiptData[id] = receipt
	mutex.Unlock()

	jsonResponse := map[string]string{"receiptID": id.String()}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(jsonResponse)
}

func getPointsHandler(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	id := variables["id"]

	uuid, err := uuid.Parse(id)
	if err != nil {
		http.Error(response, "No receipt found for that id", http.StatusNotFound)
		return
	}

	mutex.Lock()
	receipt, exists := receiptData[uuid]
	mutex.Unlock()

	if !exists {
		http.Error(response, "No receipt found for that id", http.StatusNotFound)
		return
	}

	points := points.Calculate(receipt)
	jsonResponse := map[string]int{"points": points}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(jsonResponse)
}
