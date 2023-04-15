package domain

type CustomerService interface {
    GetCustomerByID(id int) (*Customer, error)
    GetAllCustomers() ([]*Customer, error)
    CreateCustomer(c *Customer) error
    UpdateCustomer(c *Customer) error
    DeleteCustomer(id int) error
}

type OrderService interface {
    GetOrderByID(id int) (*Order, error)
    GetAllOrders() ([]*Order, error)
    CreateOrder(o *Order) error
    UpdateOrder(o *Order) error
    DeleteOrder(id int) error
}

type ProductService interface {
    GetProductByID(id int) (*Product, error)
    GetAllProducts() ([]*Product, error)
    CreateProduct(p *Product) error
    UpdateProduct(p *Product) error
    DeleteProduct(id int) error
}

type VendorService interface {
    GetVendorByID(id int) (*Vendor, error)
    GetAllVendors() ([]*Vendor, error)
    CreateVendor(v *Vendor) error
    UpdateVendor(v *Vendor) error
    DeleteVendor(id int) error
}
