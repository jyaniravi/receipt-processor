package process

import (
	"receipt-processor/types"
	"testing"
)

func TestGetWholeDollarTotalPoints(t *testing.T) {
	tests := []struct {
		name        string
		totalAmount float64
		expected    int
	}{
		{"Whole dollar amount", 50.0, 50},
		{"Non-whole dollar amount", 50.25, 0},
		{"Zero amount", 5.0, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := getWholeDollarTotalPoints(tt.totalAmount)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}

func TestGetAlphanumericRetailerPoints(t *testing.T) {
	tests := []struct {
		name     string
		retailer string
		expected int
	}{
		{"Alphanumeric characters", "SuperMart123", 12},
		{"No alphanumeric characters", "!@#$%^&*", 0},
		{"Mixed characters", "123ABC!@#", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := getAlphanumericRetailerPoints(tt.retailer)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}

func TestGetMultipleOf25Points(t *testing.T) {
	tests := []struct {
		name        string
		totalAmount float64
		expected    int
	}{
		{"Multiple of 0.25", 50.25, 25},
		{"Not a multiple of 0.25", 50.10, 0},
		{"Zero amount", 0.0, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := getMultipleOf25Points(tt.totalAmount)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}

func TestCountPointsForItems(t *testing.T) {
	tests := []struct {
		name       string
		numOfItems int
		expected   int
	}{
		{"Even number of items", 4, 10},
		{"Odd number of items", 5, 10},
		{"No items", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := countPointsForItems(tt.numOfItems)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}

func TestCountItemDescriptionLengthPoints(t *testing.T) {
	tests := []struct {
		name      string
		itemsList []types.Item
		expected  int
	}{
		{
			"Valid descriptions divisible by 3",
			[]types.Item{{Name: "Banana", Price: "2.99"}, {Name: "Apple", Price: "3.75"}},
			1,
		},
		{
			"Invalid descriptions not divisible by 3",
			[]types.Item{{Name: "Bread", Price: "1.99"}, {Name: "Milk", Price: "2.50"}},
			0,
		},
		{
			"No items",
			[]types.Item{},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := countItemDescriiptionLengthPoints(tt.itemsList)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}

func TestGetOddDatePoints(t *testing.T) {
	tests := []struct {
		name         string
		purchaseDate string
		expected     int
	}{
		{"Odd date", "2024-11-15", 6},
		{"Even date", "2024-11-16", 0},
		{"Invalid date", "invalid-date", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := getOddDatePoints(tt.purchaseDate)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}

func TestGetPurchaseTimePoints(t *testing.T) {
	tests := []struct {
		name         string
		purchaseTime string
		expected     int
	}{
		{"Within time range", "14:30", 10},
		{"Outside time range", "12:00", 0},
		{"Invalid time", "invalid-time", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := getPurchaseTimePoints(tt.purchaseTime)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}
