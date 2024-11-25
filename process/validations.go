package process

import (
	"fmt"
	"receipt-processor/types"
	"strconv"
	"time"
)

func ValidateReceipt(receipt types.Receipt) error {
	// Validate total amount
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return fmt.Errorf("total amount is invalid: %w", err)
	}
	if total < 0 {
		return fmt.Errorf("total amount cannot be negative")
	}

	// Validate each item's price
	for _, item := range receipt.Items {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return fmt.Errorf("invalid price for item '%s': %w", item.Name, err)
		}
		if price < 0 {
			return fmt.Errorf("item price cannot be negative:'%s'", item.Name)
		}

	}

	// Validate purchase date format
	_, err = time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return fmt.Errorf("purchase date is invalid: %w", err)
	}

	// Validate purchase time format
	_, err = time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return fmt.Errorf("purchase time is invalid: %w", err)
	}

	return nil
}
