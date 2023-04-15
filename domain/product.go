package domain

type Product struct {
    ID          int
    Name        string
    Description string
    Price       float64
    VendorID    int
}

type ProductRepository interface {
    GetProductByID(id int) (*Product, error)
    GetAllProducts() ([]*Product, error)
    CreateProduct(p *Product) error
    UpdateProduct(p *Product) error
    DeleteProduct(id int) error
}
