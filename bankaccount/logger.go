package main

import (
	"log"
	"time"
)

func logOperation(accountID, operation string, amount float64) {
	log.Printf("[%s] Account ID: %s, Operation: %s, Amount: %.2f", time.Now().Format(time.RFC3339), accountID, operation, amount)
}
