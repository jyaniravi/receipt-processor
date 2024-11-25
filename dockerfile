# Step 1: Use the official Go image as a builder
FROM golang:1.20 as builder

# Step 2: Set the working directory
WORKDIR /app

# Step 3: Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

# Step 4: Copy the application code
COPY . .

# Step 5: Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o receipt-processor .

# Step 6: Use a lightweight image for deployment
FROM alpine:latest

# Step 7: Set the working directory in the smaller image
WORKDIR /app

# Step 8: Copy the built executable from the builder
COPY --from=builder /app/receipt-processor .

# Step 9: Expose the application port
EXPOSE 8080

# Step 10: Command to run the application
CMD ["./receipt-processor"]