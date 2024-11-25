package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"receipt-processor/process"
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

	// Initializing the Server
	log.Println("!!!!! Starting Receipt-Processor !!!!!")
	log.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", muxHandler))

}

func processReceiptsHandler(response http.ResponseWriter, request *http.Request) {
	var receipt types.Receipt

	// Validate the receipt
	err := json.NewDecoder(request.Body).Decode(&receipt)
	if err != nil {
		http.Error(response, "The receipt is invalid", http.StatusBadRequest)
	}

	err = process.ValidateReceipt(receipt)
	if err != nil {
		http.Error(response, fmt.Sprintf("Invalid receipt data: %v", err), http.StatusBadRequest)
		return
	}

	// Generating new uuid
	id := uuid.New()
	mutex.Lock()
	receiptData[id] = receipt
	mutex.Unlock()

	// Responding to processing request
	receiptResponse := types.ReceiptProcessResponse{
		ReceiptID: id.String(),
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(receiptResponse)
}

func getPointsHandler(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	id := variables["id"]

	uuid, err := uuid.Parse(id)
	if err != nil {
		http.Error(response, "No receipt found for that id", http.StatusNotFound)
		return
	}

	// Searching receipt in cache - receiptData
	mutex.Lock()
	receipt, exists := receiptData[uuid]
	mutex.Unlock()

	if !exists {
		http.Error(response, "No receipt found for that id", http.StatusNotFound)
		return
	}

	// Calculate the points for the receipt
	points := process.Calculate(receipt)

	// Responding to get request
	getPointsResponse := types.GetPointsResponse{
		Points: points,
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(getPointsResponse)
}
