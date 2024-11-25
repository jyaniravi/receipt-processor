package process

import (
	"math"
	"receipt-processor/types"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	WholeDollarPoints  = 50
	MultipleOf25Points = 25
	OddDatePoints      = 6
	PurchaseTimePoints = 10
)

// Calculate computes the total reward points based on receipt data.
// It aggregates points from different criteria like total amount, item descriptions, and timestamps.
func Calculate(receipt types.Receipt) float64 {
	totalPoints := 0

	totalAmount, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		totalAmount = 0
	}

	totalPoints += getWholeDollarTotalPoints(totalAmount)

	totalPoints += getAlphanumericRetailerPoints(receipt.Retailer)

	totalPoints += getMultipleOf25Points(totalAmount)

	totalPoints += countPointsForItems(len(receipt.Items))

	totalPoints += countItemDescriiptionLengthPoints(receipt.Items)

	totalPoints += getOddDatePoints(receipt.PurchaseDate)

	totalPoints += getPurchaseTimePoints(receipt.PurchaseTime)

	return float64(totalPoints)
}

func getWholeDollarTotalPoints(totalAmount float64) int {
	if totalAmount == float64(int(totalAmount)) {
		return WholeDollarPoints
	}

	return 0
}

func getAlphanumericRetailerPoints(retailer string) int {
	alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)

	return len(alphanumeric.FindAllString(retailer, -1))
}

func getMultipleOf25Points(totalAmount float64) int {
	if math.Mod(totalAmount, 0.25) == 0 {
		return MultipleOf25Points
	}

	return 0
}

func countPointsForItems(numOfItems int) int {
	return (numOfItems / 2) * 5
}

func countItemDescriiptionLengthPoints(itemsList []types.Item) int {
	pointsCount := 0

	for _, item := range itemsList {
		if len(strings.TrimSpace(item.Name))%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				continue
			}
			pointsCount += int(math.Ceil(price * 0.2))
		}
	}

	return pointsCount
}

func getOddDatePoints(purchaseDate string) int {
	date, _ := time.Parse("2006-01-02", purchaseDate)

	if date.Day()%2 != 0 {
		return OddDatePoints
	}

	return 0
}

func getPurchaseTimePoints(purchaseTime string) int {
	timePurchased, _ := time.Parse("15:04", purchaseTime)

	if timePurchased.Hour() >= 14 && timePurchased.Hour() <= 16 {
		return PurchaseTimePoints
	}

	return 0
}
