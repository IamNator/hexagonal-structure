package domain

type Order struct {
    ID          int
    CustomerID  int
    ProductID   int
    Quantity    int
    TotalAmount float64
}

type OrderRepository interface {
    GetOrderByID(id int) (*Order, error)
    GetAllOrders() ([]*Order, error)
    CreateOrder(o *Order) error
    UpdateOrder(o *Order) error
    DeleteOrder(id int) error
}
