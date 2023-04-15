package stripe

type PaymentGateway interface {
	Charge(creditCardNumber string, amount float64) error
	Refund(transactionID string, amount float64) error
}

type StripePaymentGateway struct {
	apiKey string
}

func NewStripePaymentGateway(apiKey string) *StripePaymentGateway {
	return &StripePaymentGateway{
		apiKey: apiKey,
	}
}

func (s *StripePaymentGateway) Charge(creditCardNumber string, amount float64) error {
	// Charge the credit card using the Stripe API
	return nil
}

func (s *StripePaymentGateway) Refund(transactionID string, amount float64) error {
	// Refund the transaction using the Stripe API
	return nil
}