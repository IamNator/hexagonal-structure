package domain

type Customer struct {
    ID       int
    Name     string
    Email    string
    Password string
}

type CustomerRepository interface {
    GetCustomerByID(id int) (*Customer, error)
    GetCustomerByEmail(email string) (*Customer, error)
    CreateCustomer(customer *Customer) error
    UpdateCustomer(customer *Customer) error
    DeleteCustomer(id int) error
}

type CustomerService interface {
    GetCustomerByID(id int) (*Customer, error)
    GetCustomerByEmail(email string) (*Customer, error)
    CreateCustomer(name, email, password string) error
    UpdateCustomer(id int, name, email, password string) error
    DeleteCustomer(id int) error
}
