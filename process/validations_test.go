package process

import (
	"receipt-processor/types"
	"testing"
)

func TestValidateReceipt(t *testing.T) {
	tests := []struct {
		name    string
		receipt types.Receipt
		wantErr bool
	}{
		{
			name: "Valid receipt",
			receipt: types.Receipt{
				Retailer:     "SuperMart",
				PurchaseDate: "2024-11-24",
				PurchaseTime: "15:30",
				Items: []types.Item{
					{Name: "Bananas", Price: "2.99"},
					{Name: "Apples", Price: "3.75"},
				},
				Total: "6.74",
			},
			wantErr: false,
		},
		{
			name: "Negative total amount",
			receipt: types.Receipt{
				Retailer:     "SuperMart",
				PurchaseDate: "2024-11-24",
				PurchaseTime: "15:30",
				Items: []types.Item{
					{Name: "Bananas", Price: "2.99"},
					{Name: "Apples", Price: "3.75"},
				},
				Total: "-6.74",
			},
			wantErr: true,
		},
		{
			name: "Negative item price",
			receipt: types.Receipt{
				Retailer:     "SuperMart",
				PurchaseDate: "2024-11-24",
				PurchaseTime: "15:30",
				Items: []types.Item{
					{Name: "Bananas", Price: "-2.99"},
					{Name: "Apples", Price: "3.75"},
				},
				Total: "6.74",
			},
			wantErr: true,
		},
		{
			name: "Invalid total amount",
			receipt: types.Receipt{
				Retailer:     "SuperMart",
				PurchaseDate: "2024-11-24",
				PurchaseTime: "15:30",
				Items: []types.Item{
					{Name: "Bananas", Price: "2.99"},
					{Name: "Apples", Price: "3.75"},
				},
				Total: "invalid",
			},
			wantErr: true,
		},
		{
			name: "Invalid item price",
			receipt: types.Receipt{
				Retailer:     "SuperMart",
				PurchaseDate: "2024-11-24",
				PurchaseTime: "15:30",
				Items: []types.Item{
					{Name: "Bananas", Price: "invalid"},
					{Name: "Apples", Price: "3.75"},
				},
				Total: "6.74",
			},
			wantErr: true,
		},
		{
			name: "Invalid purchase date",
			receipt: types.Receipt{
				Retailer:     "SuperMart",
				PurchaseDate: "invalid-date",
				PurchaseTime: "15:30",
				Items: []types.Item{
					{Name: "Bananas", Price: "2.99"},
					{Name: "Apples", Price: "3.75"},
				},
				Total: "6.74",
			},
			wantErr: true,
		},
		{
			name: "Invalid purchase time",
			receipt: types.Receipt{
				Retailer:     "SuperMart",
				PurchaseDate: "2024-11-24",
				PurchaseTime: "invalid-time",
				Items: []types.Item{
					{Name: "Bananas", Price: "2.99"},
					{Name: "Apples", Price: "3.75"},
				},
				Total: "6.74",
			},
			wantErr: true,
		},
		{
			name: "No items in receipt",
			receipt: types.Receipt{
				Retailer:     "SuperMart",
				PurchaseDate: "2024-11-24",
				PurchaseTime: "15:30",
				Items:        []types.Item{},
				Total:        "0.00",
			},
			wantErr: false, // Assuming empty items are valid, but can be changed based on requirements
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateReceipt(tt.receipt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateReceipt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
