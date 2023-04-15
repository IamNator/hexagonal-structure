package main

import (
    "github.com/your-repo-name/vendor/application"
    "github.com/your-repo-name/vendor/infrastructure/db/postgres"
    "github.com/your-repo-name/vendor/infrastructure/payment/stripe"
    "github.com/your-repo-name/vendor/infrastructure/shipping"
    "github.com/your-repo-name/vendor/interfaces/api"
    "github.com/your-repo-name/vendor/interfaces/web"
)

func main() {
    // Initialize database
    db := postgres.NewPostgresDB()

    // Initialize payment gateway
    paymentGateway := stripe.NewStripePaymentGateway()

    // Initialize shipping service
    shippingService := shipping.NewShippingService()

    // Initialize application services
    customerService := application.NewCustomerService(db)
    productService := application.NewProductService(db, shippingService)
    orderService := application.NewOrderService(db, productService, paymentGateway)
    vendorService := application.NewVendorService(db)

    // Initialize API and web interfaces
    apiInterface := api.NewAPIInterface(customerService, orderService, productService, vendorService)
    webInterface := web.NewWebInterface(customerService, orderService, productService, vendorService)

    // Start the server
    apiInterface.StartServer()
    webInterface.StartServer()
}
