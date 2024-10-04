package handlers

import (
	"go-micro/internal/models"
	"go-micro/internal/services"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/paymentintent"
)

func HandlePayment(c *gin.Context) {
	var req models.PaymentRequest

	// binds incoming json request to PaymentRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if stripe.Key == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Stripe API key not set"})
		return
	}

	// creates the paymentintent with stripe api
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(req.Amount),
		Currency: stripe.String(req.Currency),
	}
	pi, err := paymentintent.New(params)

	if err != nil {
		log.Fatalf("Error processing payment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "payment failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client_secret": pi.ClientSecret})
}

func HandleTransaction(c *gin.Context) {
	var req services.OrderRequest

	// binds the incoming json orderrequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// generates the transaction ID
	transactionID := uuid.New().String()

	// insert the order into the db
	err := services.InsertOrder(req, transactionID)
	if err != nil {
		log.Printf("error inserting oder: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process oder"})
		return
	}

}
