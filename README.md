# Receipt Processor Challenge

## Overview

This project is a solution to the Fetch Rewards Receipt Processor Challenge. The goal of [receipt-processor-challenge](https://github.com/fetch-rewards/receipt-processor-challenge) challenge is to create a system that processes receipts and calculates points based on specific rules. The application provides two main APIs: one for processing receipts and another for retrieving the points associated with a receipt.

---

## Author

- **Name:** Ravi Jyani    
- **GitHub:** [Receipt Processor (GitHub)](https://github.com/jyaniravi/receipt-processor)

---

## What This Project Does

1. **Processes receipts:** Parses a JSON receipt to generate a unique receipt ID.  
2. **Calculates points:** Implements business rules to calculate points based on receipt details.  
3. **Stores and retrieves points:** Associates the calculated points with the receipt ID for later retrieval.  
4. **Exposes APIs:** Provides two endpoints to process receipts and fetch points.

---

## Rules for Calculating Points

Points are calculated based on the following rules:
1. **Retailer Name:**  
   - Earn 1 point for every alphanumeric character in the retailer's name.

2. **Total Amount:**  
   - If the total is a round dollar amount (e.g., `10.00`), add 50 points.
   - If the total is a multiple of 0.25, add 25 points.

3. **Items:**  
   - Earn 5 points for every two items on the receipt.
   - If an item's description has a length that is a multiple of 3, multiply the item's price by 0.2 and round up to the nearest integer. Add that many points.

4. **Purchase Date:**  
   - If the purchase date is odd, add 6 points.

5. **Purchase Time:**  
   - If the purchase time is between 2:00 PM and 4:00 PM (inclusive), add 10 points.

---

## Docker Instructions

### Build the Docker Image
Run the following command in the root directory of the project:
```bash
docker build -t receipt-processor .
```

Start the container and expose it on port 8080:
```bash
docker run -p 8080:8080 receipt-processor
```
---
## API Usage

### Endpoints

1. Process Receipts

	- URL: POST /receipts/process 
	- Description: Processes a receipt and returns a unique receipt ID.
	- Request Body Example:
    ```JSON
    {
        "retailer": "Target",
        "purchaseDate": "2022-01-01",
        "purchaseTime": "13:01",
        "total": "35.35",
        "items": [
            { "shortDescription": "Mountain Dew 12PK", "price": "6.49" },
            { "shortDescription": "Emils Cheese Pizza", "price": "12.25" }
        ]
    }
    ```
    - Response Example:
    ```JSON
    {
        "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
    }
    ```
2. 
	- URL: GET /receipts/{id}/points
	- Description: Retrieves the points associated with a processed receipt.
	- Request Example:
    ```bash
        GET /receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points
    ```
    - Response Example:
    ```JSON
    {
        "points": 28
    }
    ```
---

## Guide to the Specification Document

The OpenAPI specification (api.yml) is located at `specs/receipt-processor.yaml` and contains all details about the endpoints, request/response formats, and examples. You can load it into tools like Swagger Editor to view interactive API documentation.
