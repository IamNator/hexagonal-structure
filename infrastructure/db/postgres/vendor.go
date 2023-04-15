package postgres

import (
    "database/sql"
    "errors"
    "fmt"
    "log"

    "your-project/domain"
)

type vendorRepository struct {
    db *sql.DB
}

func NewVendorRepository(db *sql.DB) domain.VendorRepository {
    return &vendorRepository{db}
}

func (r *vendorRepository) GetVendorByID(id int) (*domain.Vendor, error) {
    v := &domain.Vendor{}
    query := `SELECT id, name, address FROM vendors WHERE id=$1`
    err := r.db.QueryRow(query, id).Scan(&v.ID, &v.Name, &v.Address)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("Vendor not found")
        }
        log.Printf("Error while getting vendor: %v", err)
        return nil, fmt.Errorf("Error while getting vendor: %w", err)
    }
    return v, nil
}

func (r *vendorRepository) GetAllVendors() ([]*domain.Vendor, error) {
    vendors := []*domain.Vendor{}
    query := `SELECT id, name, address FROM vendors`
    rows, err := r.db.Query(query)
    if err != nil {
        log.Printf("Error while getting all vendors: %v", err)
        return nil, fmt.Errorf("Error while getting all vendors: %w", err)
    }
    defer rows.Close()
    for rows.Next() {
        v := &domain.Vendor{}
        err = rows.Scan(&v.ID, &v.Name, &v.Address)
        if err != nil {
            log.Printf("Error while scanning vendor row: %v", err)
            return nil, fmt.Errorf("Error while scanning vendor row: %w", err)
        }
        vendors = append(vendors, v)
    }
    if err = rows.Err(); err != nil {
        log.Printf("Error while iterating vendor rows: %v", err)
        return nil, fmt.Errorf("Error while iterating vendor rows: %w", err)
    }
    return vendors, nil
}

func (r *vendorRepository) CreateVendor(v *domain.Vendor) error {
    query := `INSERT INTO vendors (name, address) VALUES ($1, $2) RETURNING id`
    err := r.db.QueryRow(query, v.Name, v.Address).Scan(&v.ID)
    if err != nil {
        log.Printf("Error while creating vendor: %v", err)
        return fmt.Errorf("Error while creating vendor: %w", err)
    }
    return nil
}

func (r *vendorRepository) UpdateVendor(v *domain.Vendor) error {
    query := `UPDATE vendors SET name=$2, address=$3 WHERE id=$1`
    res, err := r.db.Exec(query, v.ID, v.Name, v.Address)
    if err != nil {
        log.Printf("Error while updating vendor: %v", err)
        return fmt.Errorf("Error while updating vendor: %w", err)
    }
    rowsAffected, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error while getting affected rows after vendor update: %v", err)
        return fmt.Errorf("Error while getting affected rows after vendor update: %w", err)
    }
    if rowsAffected == 0 {
        return errors.New("Vendor not found")
    }
    return nil
}

func (r *vendorRepository) DeleteVendor(id int) error {
    query := `DELETE FROM vendors WHERE id=$1`
    res, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error while deleting vendor: %v", err)
        return fmt.Errorf("Error while deleting vendor: %w", err)
    }
    rowsAffected, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error while getting affected rows after vendor deletion: %v", err)
        return fmt.Errorf("Error while getting affected rows after vendor deletion: %w", err)
    }
    if rowsAffected == 0 {
        return errors.New("Vendor not found")
    }
    return nil
}
