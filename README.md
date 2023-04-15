# hexagonal-structure
The best implementation of the haxagonal structure I have found


```
myapp/
├── cmd/
│   └── myapp/
│       └── main.go
├── domain/
│   ├── product.go
│   └── product_repository.go
├── infrastructure/
│   ├── database/
│   │   ├── mysql/
│   │   │   └── product_repository.go
│   │   └── postgres/
│   │       └── product_repository.go
│   └── http/
│       └── product_handler.go
└── usecase/
    └── product_usecase.go

```
