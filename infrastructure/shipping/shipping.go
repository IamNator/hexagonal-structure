package shipping

type ShippingService interface {
	Ship(orderID int) error
}

type DefaultShippingService struct{}

func NewDefaultShippingService() *DefaultShippingService {
	return &DefaultShippingService{}
}

func (s *DefaultShippingService) Ship(orderID int) error {
	// Ship the order
	return nil
}