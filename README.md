# Mercado Livre Challenge - Go API

This project is a RESTful API built in Go as part of a Mercado Livre coding challenge.
The API provides endpoints to manage products, sellers, and product reviews, loading data from JSON files.

## API Endpoints

**Products**
| Method | Endpoint                        | Description                     |
| ------ | ------------------------------- | ------------------------------- |
| GET    | `/api/v1/products`              | List products (with pagination) |
| GET    | `/api/v1/products/{id}`         | Get product by ID               |
| GET    | `/api/v1/products/{id}/reviews` | Get reviews for a product       |

**Sellers**
| Method | Endpoint               | Description      |
| ------ | ---------------------- | ---------------- |
| GET    | `/api/v1/sellers/{id}` | Get seller by ID |

## Project Structure

- router/ - Router and HTTP handlers (Gin controllers)
- service/ - Business logic and data loading services
- model/ - Data models (Product, Seller, Review)
- utils/ - Utility functions (JSON loading, etc)
- response/ - Data model for response
- mocks/ - mocks structures
- logger/ - zap logger
- docs/ - swagger doc
- data/ - generated json files for the project
- cmd/main.go - Application entrypoint and router setup

## Run Instructions

For detailed setup and execution steps, see [RUN.md](RUN.md).