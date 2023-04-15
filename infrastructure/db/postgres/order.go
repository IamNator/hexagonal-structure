package postgres

import (
    "database/sql"
    "fmt"

    "path/to/domain"
)

type OrderRepository struct {
    db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
    return &OrderRepository{db: db}
}

func (repo *OrderRepository) GetOrderByID(id int) (*domain.Order, error) {
    query := "SELECT id, customer_id, product_id, quantity, total_amount FROM orders WHERE id = $1"
    row := repo.db.QueryRow(query, id)

    o := &domain.Order{}
    err := row.Scan(&o.ID, &o.CustomerID, &o.ProductID, &o.Quantity, &o.TotalAmount)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("order not found")
        }
        return nil, err
    }

    return o, nil
}

func (repo *OrderRepository) GetAllOrders() ([]*domain.Order, error) {
    query := "SELECT id, customer_id, product_id, quantity, total_amount FROM orders"
    rows, err := repo.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    orders := []*domain.Order{}
    for rows.Next() {
        o := &domain.Order{}
        err := rows.Scan(&o.ID, &o.CustomerID, &o.ProductID, &o.Quantity, &o.TotalAmount)
        if err != nil {
            return nil, err
        }
        orders = append(orders, o)
    }

    return orders, nil
}

func (repo *OrderRepository) CreateOrder(o *domain.Order) error {
    query := "INSERT INTO orders (customer_id, product_id, quantity, total_amount) VALUES ($1, $2, $3, $4) RETURNING id"
    row := repo.db.QueryRow(query, o.CustomerID, o.ProductID, o.Quantity, o.TotalAmount)

    err := row.Scan(&o.ID)
    if err != nil {
        return err
    }

    return nil
}

func (repo *OrderRepository) UpdateOrder(o *domain.Order) error {
    query := "UPDATE orders SET customer_id = $1, product_id = $2, quantity = $3, total_amount = $4 WHERE id = $5"
    _, err := repo.db.Exec(query, o.CustomerID, o.ProductID, o.Quantity, o.TotalAmount, o.ID)
    if err != nil {
        return err
    }

    return nil
}

func (repo *OrderRepository) DeleteOrder(id int) error {
    query := "DELETE FROM orders WHERE id = $1"
    _, err := repo.db.Exec(query, id)
    if err != nil {
        return err
    }

    return nil
}
