package application

import (
    "context"
    "errors"
    "strconv"
    "github.com/example/vendor/domain"
    "github.com/example/vendor/infrastructure/db/postgres"
)

type VendorService interface {
    GetVendorByID(ctx context.Context, vendorID string) (*domain.Vendor, error)
    CreateVendor(ctx context.Context, vendor *domain.Vendor) error
    UpdateVendor(ctx context.Context, vendor *domain.Vendor) error
    DeleteVendor(ctx context.Context, vendorID string) error
}

type vendorService struct {
    vendorRepo postgres.VendorRepository
}

func NewVendorService(vendorRepo postgres.VendorRepository) VendorService {
    return &vendorService{vendorRepo: vendorRepo}
}

func (s *vendorService) GetVendorByID(ctx context.Context, vendorID string) (*domain.Vendor, error) {
    id, err := strconv.Atoi(vendorID)
    if err != nil {
        return nil, errors.New("invalid vendor id")
    }

    vendor, err := s.vendorRepo.GetVendorByID(ctx, id)
    if err != nil {
        return nil, err
    }

    return vendor, nil
}

func (s *vendorService) CreateVendor(ctx context.Context, vendor *domain.Vendor) error {
    err := s.vendorRepo.CreateVendor(ctx, vendor)
    if err != nil {
        return err
    }

    return nil
}

func (s *vendorService) UpdateVendor(ctx context.Context, vendor *domain.Vendor) error {
    err := s.vendorRepo.UpdateVendor(ctx, vendor)
    if err != nil {
        return err
    }

    return nil
}

func (s *vendorService) DeleteVendor(ctx context.Context, vendorID string) error {
    id, err := strconv.Atoi(vendorID)
    if err != nil {
        return errors.New("invalid vendor id")
    }

    err = s.vendorRepo.DeleteVendor(ctx, id)
    if err != nil {
        return err
    }

    return nil
}
