package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"

	"receipt-processor/types"
)

// "receipt-processor/points"

var (
	// TODO: Mutex it so multiple requests are processed acccordingly
	receiptData = make(map[string]types.Receipt)
)

func main() {
	http.HandleFunc("/receipts/process", processReceiptsHandler)
	// http.HandleFunc("/receipts/", getPointsHandler)

	log.Println("!!!!! Starting Receipt-Processor !!!!!")

	log.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func processReceiptsHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Request is not of POST method", http.StatusMethodNotAllowed)
		return
	}

	var receipt types.Receipt
	err := json.NewDecoder(request.Body).Decode(&receipt)
	if err != nil {
		http.Error(response, "Invalid receipt posted", http.StatusBadRequest)
	}

	id := uuid.New().String()
	receiptData[id] = receipt

	jsonResponse := map[string]string{"receiptID": id}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(jsonResponse)
}

// func getPointsHandler(request *http.Request, response http.ResponseWriter) {

// }
