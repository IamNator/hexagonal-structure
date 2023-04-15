package application

import (
	"errors"
	"time"

	"github.com/example/vendor/domain"
)

type OrderService interface {
	CreateOrder(customer *domain.Customer, items []*domain.OrderItem) (*domain.Order, error)
	GetOrderById(id int) (*domain.Order, error)
	GetOrdersByCustomer(customer *domain.Customer) ([]*domain.Order, error)
}

type orderService struct {
	orderRepo     domain.OrderRepository
	productRepo   domain.ProductRepository
	customerRepo  domain.CustomerRepository
	shippingSvc   domain.ShippingService
	paymentSvc    domain.PaymentService
	taxCalculator domain.TaxCalculator
}

func NewOrderService(
	orderRepo domain.OrderRepository,
	productRepo domain.ProductRepository,
	customerRepo domain.CustomerRepository,
	shippingSvc domain.ShippingService,
	paymentSvc domain.PaymentService,
	taxCalculator domain.TaxCalculator,
) OrderService {
	return &orderService{
		orderRepo:     orderRepo,
		productRepo:   productRepo,
		customerRepo:  customerRepo,
		shippingSvc:   shippingSvc,
		paymentSvc:    paymentSvc,
		taxCalculator: taxCalculator,
	}
}

func (s *orderService) CreateOrder(customer *domain.Customer, items []*domain.OrderItem) (*domain.Order, error) {
	// calculate subtotal and tax
	subtotal := 0.0
	for _, item := range items {
		product, err := s.productRepo.GetById(item.ProductId)
		if err != nil {
			return nil, err
		}
		if product.Stock < item.Quantity {
			return nil, errors.New("product is out of stock")
		}
		item.Price = product.Price
		subtotal += product.Price * float64(item.Quantity)
	}

	tax := s.taxCalculator.CalculateTax(subtotal)

	// calculate shipping cost
	shippingCost, err := s.shippingSvc.CalculateShippingCost(customer, items)
	if err != nil {
		return nil, err
	}

	// create order
	order := &domain.Order{
		CustomerId:   customer.Id,
		Items:        items,
		Subtotal:     subtotal,
		Tax:          tax,
		ShippingCost: shippingCost,
		Total:        subtotal + tax + shippingCost,
		CreatedAt:    time.Now(),
	}

	// charge payment
	if err := s.paymentSvc.Charge(order.Total); err != nil {
		return nil, err
	}

	// reduce product stock
	for _, item := range items {
		if err := s.productRepo.ReduceStock(item.ProductId, item.Quantity); err != nil {
			return nil, err
		}
	}

	// save order
	if err := s.orderRepo.Save(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) GetOrderById(id int) (*domain.Order, error) {
	return s.orderRepo.GetById(id)
}

func (s *orderService) GetOrdersByCustomer(customer *domain.Customer) ([]*domain.Order, error) {
	return s.orderRepo.GetByCustomer(customer)
}
