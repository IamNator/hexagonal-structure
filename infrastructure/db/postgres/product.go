package postgres

import (
    "database/sql"
    "errors"

    "path/to/domain"
)

type ProductRepository struct {
    db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
    return &ProductRepository{
        db: db,
    }
}

func (r *ProductRepository) GetProductByID(id int) (*domain.Product, error) {
    var p domain.Product
    err := r.db.QueryRow("SELECT id, name, description, price, vendor_id FROM products WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.VendorID)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("product not found")
        }
        return nil, err
    }
    return &p, nil
}

func (r *ProductRepository) GetAllProducts() ([]*domain.Product, error) {
    rows, err := r.db.Query("SELECT id, name, description, price, vendor_id FROM products")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []*domain.Product
    for rows.Next() {
        var p domain.Product
        err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.VendorID)
        if err != nil {
            return nil, err
        }
        products = append(products, &p)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return products, nil
}

func (r *ProductRepository) CreateProduct(p *domain.Product) error {
    err := r.db.QueryRow("INSERT INTO products(name, description, price, vendor_id) VALUES ($1, $2, $3, $4) RETURNING id", p.Name, p.Description, p.Price, p.VendorID).Scan(&p.ID)
    if err != nil {
        return err
    }
    return nil
}

func (r *ProductRepository) UpdateProduct(p *domain.Product) error {
    res, err := r.db.Exec("UPDATE products SET name = $1, description = $2, price = $3, vendor_id = $4 WHERE id = $5", p.Name, p.Description, p.Price, p.VendorID, p.ID)
    if err != nil {
        return err
    }
    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return err
    }
    if rowsAffected == 0 {
        return errors.New("product not found")
    }
    return nil
}

func (r *ProductRepository) DeleteProduct(id int) error {
    res, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
    if err != nil {
        return err
    }
    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return err
    }
    if rowsAffected == 0 {
        return errors.New("product not found")
    }
    return nil
}
