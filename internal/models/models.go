package models

type PaymentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

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
