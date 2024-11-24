package points

import (
	"math"
	"receipt-processor/types"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Calculate(receipt types.Receipt) int {
	totalPoints := 0

	totalAmount, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0
	}

	totalPoints += getWholeDollarTotalPoints(totalAmount)

	totalPoints += getAlphanumericRetailerPoints(receipt.Retailer)

	totalPoints += getMultipleOf25Points(totalAmount)

	totalPoints += countPointsForItems(len(receipt.Items))

	totalPoints += countItemDescriiptionLengthPoints(receipt.Items)

	totalPoints += getOddDatePoints(receipt.PurchaseDate)

	totalPoints += getPurchaseTimePoints(receipt.PurchaseTime)

	return totalPoints
}

func getWholeDollarTotalPoints(totalAmount float64) int {
	if totalAmount == float64(int(totalAmount)) {
		return 50
	}

	return 0
}

func getAlphanumericRetailerPoints(retailer string) int {
	alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)

	return len(alphanumeric.FindAllString(retailer, -1))
}

func getMultipleOf25Points(totalAmount float64) int {
	if math.Mod(totalAmount, 0.25) == 0 {
		return 25
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
	date, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		return 0
	}

	if date.Day()%2 != 0 {
		return 6
	}

	return 0
}

func getPurchaseTimePoints(purchaseTime string) int {
	timePurchased, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		return 0
	}

	if timePurchased.Hour() >= 14 && timePurchased.Hour() <= 16 {
		return 10
	}

	return 0
}
