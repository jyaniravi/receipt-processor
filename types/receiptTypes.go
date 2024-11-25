package types

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	Name  string `json:"shortDescription"`
	Price string `json:"price"`
}

type ReceiptProcessResponse struct {
	ReceiptID string `json:"id"`
}

type GetPointsResponse struct {
	Points float64 `json:"points"`
}
