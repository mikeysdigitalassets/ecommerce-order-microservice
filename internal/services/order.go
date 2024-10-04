package services

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type OrderRequest struct {
	UserID          int    `json:"user_id"`
	ShippingAddress string `json:"shipping_address"`
	City            string `json:"city"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	State           string `json:"state"`
	PostalCode      string `json:"postal_code"`
	OrderDate       string `json:"order_date"`
	TotalAmount     int64  `json:"total_amount"`
}

func InsertOrder(order OrderRequest, transactionID string) error {
	db, err := sql.Open("postgres", "postgres://postgres:kali@localhost:5432/ecommerce_db?sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
	INSERT INTO orders (user_id, shipping_address, city, first_name, last_name, 
	state, postal_code, order_date, total_amount, transaction_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	log.Printf("Inserting into DB: %s", order.OrderDate)

	_, err = db.Exec(query, order.UserID, order.ShippingAddress, order.City,
		order.FirstName, order.LastName, order.State, order.PostalCode, order.OrderDate,
		order.TotalAmount, transactionID)
	if err != nil {
		return err
	}
	return nil
}
