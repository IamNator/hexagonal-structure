# hexagonal-structure
The best implementation of the haxagonal structure I have found


Let's say you're building an e-commerce platform that allows vendors to manage their own products and customers to search and purchase them. Here's how you could structure your code using the hexagonal architecture:

Domain layer: This layer contains the core domain logic for the product application. It might include things like creating and updating products, managing product inventory, handling customer orders, and processing payments. You would define interfaces for any external dependencies that the domain layer needs, such as a payment gateway interface or a shipping interface.

Infrastructure layer: This layer contains the implementation of the interfaces defined in the domain layer. For example, you might have a PostgreSQL implementation of the database interface, or a Stripe implementation of the payment gateway interface.

Application layer: This layer contains the business logic for the product application. It acts as the glue between the domain and infrastructure layers, coordinating the interaction between them. For example, the application layer might retrieve product data from the database using the database interface defined in the domain layer, then pass that data to the user interface layer for display.

Interfaces layer: This layer contains the implementation of the user interfaces for the product application. For example, you might have a web interface that allows vendors to manage their products and customers to search and purchase them. You would define interfaces for the application layer to interact with the user interface layer.


```
├── cmd
│   ├── main.go
│   └── vendor
├── domain
│   ├── customer.go
│   ├── order.go
│   ├── product.go
│   ├── interfaces.go
│   └── vendor.go
├── infrastructure
│   ├── db
│   │   └── postgres
│   │       ├── order.go
│   │       ├── product.go
│   │       └── vendor.go
│   ├── payment
│   │   └── stripe
│   │       └── payment.go
│   └── shipping
│       └── shipping.go
├── interfaces
│   ├── api
│   │   ├── handlers.go
│   │   ├── middleware.go
│   │   └── routes.go
│   └── web
│       ├── handlers.go
│       └── templates
└── vendor
    ├── application
    │   ├── customer_service.go
    │   ├── order_service.go
    │   ├── product_service.go
    │   └── vendor_service.go
    └── config
        └── config.go


```
