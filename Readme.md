# E-commerce API Service

A RESTful API service for managing products and brands built with Go, Fiber, and PostgreSQL.

## 📋 Prerequisites

Before running this application, make sure you have the following installed:

- **Go** (version 1.19 or later)
- **PostgreSQL** (version 12 or later)
- **Git**

## 🛠️ Installation

### 1. Clone the Repository

```bash
git clone https://github.com/rezajo220/ecommerce.git
cd ecommerce
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Install Swagger CLI (Optional)

For generating Swagger documentation:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## ⚙️ Configuration

### 1. Environment Variables

Create a `.env` file in the root directory:

```env
# Server Configuration
SERVER_PORT=8000
SERVER_READ_TIMEOUT=5000
SERVER_WRITE_TIMEOUT=5000

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ecommerce
DB_SSL_MODE=disable
```

### 2. Database Setup

1. **Create Database:**
   ```bash
   psql -h localhost -p 5432 -U postgres
   ```
   
   ```sql
   CREATE DATABASE ecommerce;
   \q
   ```

2. **Run Database Migrations:**
   ```bash
   psql -h localhost -p 5432 -U postgres -d ecommerce
   ```
   
   ```sql
   -- Enable UUID extension
   CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
   
   -- Create brands table
   CREATE TABLE brands (
       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
       brand_name TEXT NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   
   -- Create products table
   CREATE TABLE products (
       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
       product_name TEXT NOT NULL,
       price NUMERIC NOT NULL,
       qty NUMERIC NOT NULL DEFAULT 0,
       brand_id UUID NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       FOREIGN KEY (brand_id) REFERENCES brands(id) ON DELETE RESTRICT
   );
   
   -- Insert sample brands
   INSERT INTO brands (brand_name) VALUES 
   ('Samsung'),
   ('Apple'),
   ('Xiaomi'),
   ('Oppo'),
   ('Vivo');
   ```

## 🚀 Running the Application

### Development Mode

#### Option 1: Using Go Commands

```bash
# Run from project root
go run ./cmd

# Or build and run
go build -o bin/ecommerce.exe ./cmd
./bin/ecommerce.exe
```

#### Option 2: Using Make (if installed)

```bash
make run    # Run the application
make build  # Build the application
make clean  # Clean build artifacts
```

## 📚 API Documentation

### Interactive Documentation

Once the application is running, access the Swagger UI at:

```
http://localhost:8000/swagger/
```

**Response:**
```json
{
  "status": "ok",
  "service": "e-commerce-api"
}
```

## 🔗 API Endpoints

### Products

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/v1/products/` | Create a new product |
| `GET` | `/api/v1/products/` | Get products with pagination |
| `PUT` | `/api/v1/products/{id}` | Update a product |
| `DELETE` | `/api/v1/products/{id}` | Delete a product |

### Brands

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/v1/brands/` | Create a new brand |
| `GET` | `/api/v1/brands/` | Get all brands |
| `DELETE` | `/api/v1/brands/{id}` | Delete a brand |

## 📝 API Usage Examples

### Create a Brand

```bash
curl -X POST http://localhost:8000/api/v1/brands/ \
  -H "Content-Type: application/json" \
  -d '{
    "brand_name": "Samsung"
  }'
```

**Response:**
```json
{
  "message": "Brand created successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "brand_name": "Samsung",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### Create a Product

```bash
curl -X POST http://localhost:8000/api/v1/products/ \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Galaxy S24",
    "price": 12000000.50,
    "qty": 50.0,
    "brand_id": "550e8400-e29b-41d4-a716-446655440000"
  }'
```

### Get Products with Pagination

```bash
curl "http://localhost:8000/api/v1/products/?page=1&limit=10"
```

**Response:**
```json
{
  "message": "Products retrieved successfully",
  "data": {
    "products": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440001",
        "product_name": "Galaxy S24",
        "price": 12000000.50,
        "qty": 50.0,
        "brand_id": "550e8400-e29b-41d4-a716-446655440000",
        "brand_name": "Samsung",
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "limit": 10,
    "total_pages": 1
  }
}
```

### Update a Product

```bash
curl -X PUT http://localhost:8000/api/v1/products/550e8400-e29b-41d4-a716-446655440001 \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Galaxy S24 Ultra",
    "price": 15000000.00,
    "qty": 30.0
  }'
```

### Delete a Product

```bash
curl -X DELETE http://localhost:8000/api/v1/products/550e8400-e29b-41d4-a716-446655440001
```

## 📁 Project Structure

```
ecommerce/
├── cmd/                          # Application entrypoint
│   ├── main.go                   # Main application
│   ├── config.go                 # Configuration management
│   └── bootstrap.go              # Database connection
├── internal/                     # Private application code
│   ├── domain/                   # Domain models and DTOs
│   │   ├── product.go
│   │   ├── brand.go
│   │   └── responses.go
│   ├── repository/               # Data access layer
│   │   ├── product_repository.go
│   │   └── brand_repository.go
│   ├── service/                  # Business logic layer
│   │   ├── product_service.go
│   │   └── brand_service.go
│   └── handler/                  # HTTP handlers
│       ├── product_handler.go
│       ├── brand_handler.go
│       └── routes/               # Route definitions
│           ├── product_routes.go
│           └── brand_routes.go
├── docs/                         # Generated Swagger documentation
├── .env                          # Environment variables
├── go.mod                        # Go modules
├── go.sum                        # Go dependencies
├── README.md                     # This file
```

## 🚧 Troubleshooting

### Common Issues

#### 1. Database Connection Error

```
Failed to connect to database: pq: password authentication failed
```

**Solution:** Check your `.env` file and ensure PostgreSQL is running.

#### 2. Table Does Not Exist Error

```
pq: relation "brands" does not exist
```

**Solution:** Run the database migrations as described in the Database Setup section.

#### 3. Port Already in Use

```
listen tcp :8000: bind: address already in use
```

**Solution:** Change the `SERVER_PORT` in your `.env` file or stop the process using port 8000.

#### 4. Swagger Documentation Not Loading

**Solution:** 
1. Ensure Swagger dependencies are installed
2. Generate docs: `swag init -g cmd/main.go -o docs/`
3. Restart the application