package domain

type Vendor struct {
    ID      int
    Name    string
    Address string
}

type VendorRepository interface {
    GetVendorByID(id int) (*Vendor, error)
    GetAllVendors() ([]*Vendor, error)
    CreateVendor(v *Vendor) error
    UpdateVendor(v *Vendor) error
    DeleteVendor(id int) error
}
