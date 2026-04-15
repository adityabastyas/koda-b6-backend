# Koda B6 Backend

Backend API untuk aplikasi Koda B6 - Platform E-Commerce yang dibangun menggunakan **Go** dengan framework **Gin**, **PostgreSQL** sebagai database, dan **Redis** untuk caching.

## 📋 Daftar Isi

- [Fitur Utama](#fitur-utama)
- [Tech Stack](#tech-stack)
- [Prasyarat](#prasyarat)
- [Instalasi](#instalasi)
- [Konfigurasi Environment](#konfigurasi-environment)
- [Menjalankan Aplikasi](#menjalankan-aplikasi)
- [Struktur Proyek](#struktur-proyek)
- [API Endpoints](#api-endpoints)
- [Database Schema](#database-schema)
- [Docker Deployment](#docker-deployment)
- [Dokumentasi API](#dokumentasi-api)
- [Kontribusi](#kontribusi)
- [Lisensi](#lisensi)

---

## ✨ Fitur Utama

### Authentication & Authorization
- ✅ User Registration dan Login
- ✅ JWT Token Authentication
- ✅ Role-Based Access Control (RBAC) untuk Admin
- ✅ Forgot Password dengan Email Verification
- ✅ User Profile Management

### Product Management
- ✅ CRUD Operations untuk Products
- ✅ Multiple Product Images per Product
- ✅ Product Variants dan Sizes
- ✅ Product Categories
- ✅ Stock Management

### Shopping Features
- ✅ Shopping Cart Management
- ✅ Discount & Promo System
- ✅ Transactions & Order History
- ✅ Order Details dengan Product Information
- ✅ Product Reviews & Ratings

### Additional Features
- ✅ CORS Support
- ✅ File Upload (Product Images, User Photos)
- ✅ Redis Caching
- ✅ Swagger API Documentation
- ✅ Database Migrations

---

## 🛠 Tech Stack

| Component | Technology | Version |
|-----------|-----------|---------|
| **Language** | Go | 1.25.0 |
| **Web Framework** | Gin | 1.12.0 |
| **Database** | PostgreSQL | 18-alpine |
| **Cache** | Redis | 8.6.1 |
| **Authentication** | JWT | 5.3.1 |
| **Password Hashing** | Argon2 | 1.5.0 |
| **API Documentation** | Swagger/OpenAPI | 1.16.6 |
| **Containerization** | Docker | Latest |

---

## 📋 Prasyarat

Sebelum memulai, pastikan Anda telah menginstal:

- **Go** >= 1.25.0 ([Download](https://golang.org/dl/))
- **PostgreSQL** >= 18 ([Download](https://www.postgresql.org/download/))
- **Redis** >= 8.0 ([Download](https://redis.io/download))
- **Docker & Docker Compose** (Optional untuk containerization)
- **Git**
- **curl** atau **Postman** (untuk testing API)

---

## 🚀 Instalasi

### 1. Clone Repository

```bash
git clone https://github.com/adityabastyas/koda-b6-backend.git
cd koda-b6-backend
```

### 2. Install Go Dependencies

```bash
go mod tidy
```

### 3. Setup Database

Buat database PostgreSQL:

```bash
# Masuk ke PostgreSQL
psql -U postgres

# Buat database
CREATE DATABASE koda_b6;

# Exit
\q
```

### 4. Run Database Migrations

Pastikan Anda memiliki `migrate` CLI atau jalankan manual:

```bash
# Menggunakan migrate CLI
migrate -path migrations -database "postgresql://username:password@localhost:5432/koda_b6?sslmode=disable" up

# Atau jalankan queries secara manual dari file migrations/
```

### 5. Setup Environment Variables

Copy file `.env.example` ke `.env`:

```bash
cp .env.example .env
```

Atau buat file `.env` baru dengan konfigurasi berikut:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=koda_b6

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_DB=0

# Server Configuration
PORT=8888
APP_ENV=development

# JWT Configuration
JWT_SECRET=your_jwt_secret_key_here
JWT_EXPIRATION=24h

# Email Configuration (untuk Forgot Password)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your_email@gmail.com
SMTP_PASSWORD=your_app_password

# File Upload
UPLOAD_PATH=./uploads
MAX_FILE_SIZE=5242880  # 5MB
```

---

## ⚙️ Konfigurasi Environment

### Database Configuration

```env
DB_HOST=localhost         # Host database
DB_PORT=5432             # Port PostgreSQL
DB_USER=postgres         # Username database
DB_PASSWORD=password     # Password database
DB_NAME=koda_b6          # Nama database
DB_SSLMODE=disable       # SSL mode untuk connection
```

### Redis Configuration

```env
REDIS_HOST=localhost     # Host Redis
REDIS_PORT=6379         # Port Redis
REDIS_DB=0              # Redis database number
REDIS_PASSWORD=         # Password Redis (optional)
```

### JWT & Authentication

```env
JWT_SECRET=your_secret_key    # Secret key untuk signing JWT
JWT_EXPIRATION=24h            # Token expiration time
```

### Server Configuration

```env
PORT=8888              # Port untuk menjalankan server
APP_ENV=development    # Environment (development/production)
```

---

## 🏃 Menjalankan Aplikasi

### Development Mode (Local)

```bash
# Jalankan aplikasi
go run cmd/main.go

# Output:
# [GIN-debug] Loaded HTML Templates...
# [GIN-debug] POST   /auth/register
# [GIN-debug] POST   /auth/login
# ...
# [GIN-debug] Listening and serving HTTP on :8888
```

### Production Mode (Build)

```bash
# Build aplikasi
go build -o backend cmd/main.go

# Jalankan binary
./backend
```

### Docker Compose

```bash
# Jalankan semua services (PostgreSQL, Redis, Backend)
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs -f backend
```

---

## 📁 Struktur Proyek

```
koda-b6-backend/
├── cmd/
│   └── main.go                 # Entry point aplikasi
├── internal/
│   ├── di/
│   │   └── container.go        # Dependency Injection setup
│   ├── handlers/
│   │   ├── auth.go             # Auth endpoints
│   │   ├── user.go             # User endpoints
│   │   ├── product.go          # Product endpoints
│   │   ├── kategory.go         # Category endpoints
│   │   ├── promo.go            # Promo endpoints
│   │   ├── discount.go         # Discount endpoints
│   │   ├── cart.go             # Cart endpoints
│   │   ├── transaction.go      # Transaction endpoints
│   │   ├── reviews.go          # Reviews endpoints
│   │   └── ...                 # Handler lainnya
│   ├── models/
│   │   ├── user.go             # User model
│   │   ├── product.go          # Product model
│   │   ├── cart.go             # Cart model
│   │   └── ...                 # Model lainnya
│   ├── repository/
│   │   ├── user_repo.go        # User data access
│   │   ├── product_repo.go     # Product data access
│   │   └── ...                 # Repository lainnya
│   ├── service/
│   │   ├── user.go             # User business logic
│   │   ├── product.go          # Product business logic
│   │   └── ...                 # Service lainnya
│   ├── routes/
│   │   └── routes.go           # API routes configuration
│   ├── lib/
│   │   ├── db.go               # Database connection
│   │   ├── redis.go            # Redis connection
│   │   ├── middleware.go       # Middleware definitions
│   │   └── ...                 # Utility functions
│   └── config/
│       └── config.go           # Configuration management
├── migrations/
│   ├── 000001_init_db.up.sql
│   ├── 000001_init_db.down.sql
│   ├── 000002_create_forgot_password.up.sql
│   └── 000002_create_forgot_password.down.sql
├── docs/
│   ├── swagger.json            # Swagger API documentation
│   ├── swagger.yaml            # Swagger YAML format
│   └── docs.go                 # Swagger generated file
├── uploads/
│   └── [image files]           # Uploaded files directory
├── go.mod                      # Go module definition
├── go.sum                      # Go dependencies hash
├── docker-compose.yml          # Docker Compose configuration
├── Dockerfile                  # Docker build configuration
├── .env.example                # Environment variables template
├── .gitignore                  # Git ignore rules
└── README.md                   # Project documentation
```

---

## 🔌 API Endpoints

### Authentication
```
POST   /auth/register           # Register user baru
POST   /auth/login              # Login user
POST   /auth/forgot-password    # Request forgot password
PATCH  /auth/forgot-password    # Reset password
```

### Users
```
GET    /users                   # Get all users (Protected)
POST   /users/upload            # Upload user photo (Protected)
PATCH  /users/profile           # Update user profile (Protected)
```

### Products
```
GET    /products                # Get all products
GET    /products/:id            # Get product by ID
POST   /products                # Create product (Admin)
PATCH  /products/:id            # Update product (Admin)
DELETE /products/:id            # Delete product (Admin)
```

### Categories
```
GET    /kategorys               # Get all categories
GET    /kategorys/:id           # Get category by ID
POST   /kategorys               # Create category (Admin)
PATCH  /kategorys/:id           # Update category (Admin)
DELETE /kategorys/:id           # Delete category (Admin)
```

### Promos
```
GET    /promos                  # Get all promos
GET    /promos/:id              # Get promo by ID
POST   /promos                  # Create promo (Admin)
PATCH  /promos/:id              # Update promo (Admin)
DELETE /promos/:id              # Delete promo (Admin)
```

### Discounts
```
GET    /discounts               # Get all discounts
GET    /discounts/:id           # Get discount by ID
POST   /discounts               # Create discount (Protected)
PATCH  /discounts/:id           # Update discount (Protected)
DELETE /discounts/:id           # Delete discount (Protected)
```

### Cart
```
GET    /carts                   # Get all carts (Protected)
GET    /carts/:user_id          # Get cart by user ID (Protected)
```

### Cart Items
```
GET    /cart-items/:user_id     # Get cart items by user ID (Protected)
POST   /cart-items/:user_id     # Add item to cart (Protected)
PATCH  /cart-items/:id          # Update cart item (Protected)
DELETE /cart-items/:id          # Delete cart item (Protected)
```

### Transactions
```
GET    /transactions            # Get all transactions (Protected)
GET    /transactions/:id        # Get transaction by ID (Protected)
GET    /transactions/user/:user_id  # Get user transactions (Protected)
POST   /transactions            # Create transaction (Protected)
DELETE /transactions/:id        # Delete transaction (Protected)
```

### Transaction Products
```
GET    /transaction-products/:transaction_id  # Get products (Protected)
POST   /transaction-products/:transaction_id  # Add product (Protected)
DELETE /transaction-products/:id              # Delete product (Protected)
```

### Product Variants
```
GET    /product-variant/:product_id           # Get variants by product
GET    /product-variant/detail/:id            # Get variant detail
POST   /product-variant                       # Create variant (Protected)
PATCH  /product-variant/:id                   # Update variant (Protected)
DELETE /product-variant/:id                   # Delete variant (Protected)
```

### Product Sizes
```
GET    /product-sizes/:product_id             # Get sizes by product
GET    /product-sizes/detail/:id              # Get size detail
POST   /product-sizes                         # Create size (Protected)
PATCH  /product-sizes/:id                     # Update size (Protected)
DELETE /product-sizes/:id                     # Delete size (Protected)
```

### Product Images
```
GET    /product-images/:product_id            # Get images by product
POST   /product-images                        # Upload image (Protected)
DELETE /product-images/:id                    # Delete image (Protected)
```

### Reviews
```
GET    /reviews                               # Get all reviews
GET    /reviews/product/:product_id           # Get reviews by product
GET    /reviews/user/:user_id                 # Get user reviews (Protected)
POST   /reviews/:user_id                      # Create review (Protected)
DELETE /reviews/:id                           # Delete review (Protected)
```

---

## 🗄️ Database Schema

### Database Diagram

```
┌─────────────────┐     ┌──────────────────┐
│     users       │     │   categories     │
├─────────────────┤     ├──────────────────┤
│ id (PK)         │     │ id (PK)          │
│ email           │     │ name             │
│ password        │     │ description      │
│ name            │     │ created_at       │
│ photo_url       │     │ updated_at       │
│ role            │     └──────────────────┘
│ created_at      │              │
│ updated_at      │              │ 1:N
│ deleted_at      │              │
└─────────────────┘              │
        │                         │
        │ 1:N                     │
        │                    ┌──────────────┐
        │                    │   products   │
        │                    ├──────────────┤
        │                    │ id (PK)      │
        │                    │ name         │
        │                    │ description  │
        │                    │ price        │
        │                    │ category_id  │
        │                    │ stock        │
        │                    │ created_at   │
        │                    │ updated_at   │
        │                    │ deleted_at   │
        │                    └──────────────┘
        │                           │
        ├──────────────┬────────────┤
        │              │            │
        │         1:N  │            │ 1:N
        │              │            │
   ┌────────────┐  ┌────────────────────┐  ┌─────────────────┐
   │   carts    │  │ product_variants   │  │ product_images  │
   ├────────────┤  ├────────────────────┤  ├─────────────────┤
   │ id (PK)    │  │ id (PK)            │  │ id (PK)         │
   │ user_id(FK)│  │ product_id (FK)    │  │ product_id (FK) │
   │ total      │  │ color              │  │ url             │
   │ created_at │  │ size_id (FK)       │  │ created_at      │
   │ updated_at │  │ created_at         │  │ updated_at      │
   └────────────┘  │ updated_at         │  └─────────────────┘
        │          │ deleted_at         │
        │          └────────────────────┘
        │ 1:N              │
        │                  │ 1:N
   ┌─────────────┐    ┌──────────────────┐
   │ cart_items  │    │ product_sizes    │
   ├─────────────┤    ├──────────────────┤
   │ id (PK)     │    │ id (PK)          │
   │ cart_id(FK) │    │ product_id (FK)  │
   │ product_id  │    │ size             │
   │ quantity    │    │ stock            │
   │ price       │    │ created_at       │
   │ created_at  │    │ updated_at       │
   │ updated_at  │    │ deleted_at       │
   └─────────────┘    └──────────────────┘

┌──────────────────┐     ┌──────────────────────┐
│  transactions    │     │ transaction_products │
├──────────────────┤     ├──────────────────────┤
│ id (PK)          │     │ id (PK)              │
│ user_id (FK)     │     │ transaction_id (FK)  │
│ total_amount     │     │ product_id (FK)      │
│ status           │     │ quantity             │
│ created_at       │     │ price                │
│ updated_at       │     │ created_at           │
│ deleted_at       │     │ updated_at           │
└──────────────────┘     └──────────────────────┘
        │                        │
        │ 1:N                    │ 1:N
        │                        │
┌───────────────────┐      ┌──────────────┐
│     reviews       │      │   discounts  │
├───────────────────┤      ├──────────────┤
│ id (PK)           │      │ id (PK)      │
│ user_id (FK)      │      │ name         │
│ product_id (FK)   │      │ code         │
│ rating            │      │ discount_pct │
│ comment           │      │ active       │
│ created_at        │      │ created_at   │
│ updated_at        │      │ updated_at   │
│ deleted_at        │      │ deleted_at   │
└───────────────────┘      └──────────────┘

┌──────────────┐
│   promos     │
├──────────────┤
│ id (PK)      │
│ title        │
│ description  │
│ discount_pct │
│ active       │
│ start_date   │
│ end_date     │
│ created_at   │
│ updated_at   │
│ deleted_at   │
└──────────────┘
```

---

## 🐳 Docker Deployment

### Build Docker Image

```bash
# Build image
docker build -t koda-b6-backend:latest .

# Run container
docker run -p 8888:8888 --env-file .env koda-b6-backend:latest
```

### Docker Compose (Recommended)

```bash
# Start all services (PostgreSQL, Redis, Backend)
docker-compose up -d

# Check status
docker-compose ps

# View logs
docker-compose logs -f backend

# Stop services
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

### Docker Compose Services

- **PostgreSQL**: Port 5432
- **Redis**: Port 6379
- **Backend**: Port 20500 (mapped to 8888)
- **Frontend**: Port 20501 (Nginx)

---

## 📚 Dokumentasi API

API documentation tersedia melalui Swagger/OpenAPI. Setelah aplikasi berjalan, akses:

```
http://localhost:8888/swagger/index.html
```

### Mengupdate Swagger Documentation

Jika Anda mengubah handler atau menambahkan endpoint baru, update Swagger:

```bash
# Install swag jika belum
go install github.com/swaggo/swag/cmd/swag@latest

# Generate documentation
swag init -g cmd/main.go

# Documentation akan di-generate ke folder docs/
```

### Testing API dengan cURL

```bash
# Register user
curl -X POST http://localhost:8888/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123","name":"John Doe"}'

# Login
curl -X POST http://localhost:8888/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'

# Get all products
curl -X GET http://localhost:8888/products

# Get user profile (requires token)
curl -X GET http://localhost:8888/users \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## 🧪 Testing

### Unit Tests

```bash
# Run all tests
go test ./...

# Run tests dengan coverage
go test -cover ./...

# Run specific test
go test -run TestUserHandler ./internal/handlers
```

### Integration Tests

```bash
# Jalankan dengan test database
DB_NAME=test_koda_b6 go test -v ./...
```

---

## 🔐 Security Best Practices

1. **Never commit `.env` file** - Gunakan `.env.example` sebagai template
2. **JWT Secret** - Gunakan secret yang kuat dan random
3. **Password Hashing** - Semua password di-hash menggunakan Argon2
4. **CORS Configuration** - Restrict origin sesuai kebutuhan
5. **Database Credentials** - Gunakan environment variables, jangan hardcode
6. **HTTPS** - Gunakan HTTPS di production
7. **Rate Limiting** - Implementasikan rate limiting untuk authentication endpoints
8. **SQL Injection Prevention** - Gunakan parameterized queries (sudah built-in di pgx)

---

## 📝 Database Migration Guide

### Membuat Migration Baru

```bash
# Install migrate CLI
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Buat migration baru
migrate create -ext sql -dir migrations -seq create_table_name

# Contoh:
migrate create -ext sql -dir migrations -seq add_column_to_users
```

### Menjalankan Migrations

```bash
# Up (apply all pending migrations)
migrate -path migrations \
  -database "postgresql://username:password@localhost:5432/koda_b6?sslmode=disable" up

# Down (rollback 1 migration)
migrate -path migrations \
  -database "postgresql://username:password@localhost:5432/koda_b6?sslmode=disable" down

# Goto specific version
migrate -path migrations \
  -database "postgresql://username:password@localhost:5432/koda_b6?sslmode=disable" goto 2
```

---

## 🛠️ Development Tips

### Hot Reload dengan Nodemon (Optional)

Untuk development yang lebih cepat, gunakan nodemon:

```bash
# Install nodemon
npm install -g nodemon

# Jalankan dengan auto-reload
nodemon --exec go run cmd/main.go --ext go
```

### Debugging

```bash
# Gunakan Delve debugger
go install github.com/go-delve/delve/cmd/dlv@latest

# Start debugging
dlv debug cmd/main.go
```

### Code Quality

```bash
# Format code
go fmt ./...

# Lint code
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
golangci-lint run ./...

# Vet code
go vet ./...
```

---

## 📦 Dependencies Management

### Update Dependencies

```bash
# Check outdated packages
go list -u -m all

# Update specific package
go get -u github.com/gin-gonic/gin

# Update all packages
go get -u ./...

# Tidy modules
go mod tidy
```

---

## 🚨 Troubleshooting

### Port Already in Use

```bash
# Linux/Mac - Find and kill process on port 8888
lsof -i :8888
kill -9 <PID>

# Windows - Find and kill process
netstat -ano | findstr :8888
taskkill /PID <PID> /F
```

### Database Connection Error

```
error: connection refused
```

**Solution**: Pastikan PostgreSQL running dan credentials di `.env` benar.

### Redis Connection Error

```
error: cannot get address
```

**Solution**: Pastikan Redis running pada host dan port yang benar.

### Migration Issues

```bash
# Reset database (hati-hati!)
DROP DATABASE koda_b6;
CREATE DATABASE koda_b6;

# Re-run migrations
migrate -path migrations -database "postgresql://..." up
```

---

## 📄 Project Information

- **Repository**: [GitHub - koda-b6-backend](https://github.com/adityabastyas/koda-b6-backend)
- **Author**: Aditya Bastyas
- **License**: MIT
- **Status**: Active Development

---

## 🤝 Kontribusi

Kami terbuka untuk kontribusi! Silakan:

1. Fork repository
2. Buat branch untuk feature Anda (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Open Pull Request

### Coding Standards

- Follow Go conventions dan best practices
- Gunakan meaningful variable names
- Add comments untuk complex logic
- Write unit tests untuk new features
- Update documentation

---

## 📞 Support & Contact

Untuk questions atau issues:

- **GitHub Issues**: [Create Issue](https://github.com/adityabastyas/koda-b6-backend/issues)
- **Email**: Hubungi maintainer melalui GitHub profile

---

## 📄 Lisensi

Project ini dilisensikan di bawah MIT License - lihat file [LICENSE](LICENSE) untuk details.

---

## 🎯 Roadmap

- [ ] Implement caching strategy for frequently accessed data
- [ ] Add payment gateway integration
- [ ] Implement email notifications
- [ ] Add advanced search and filtering
- [ ] Implement real-time notifications dengan WebSocket
- [ ] Add batch operations for admin
- [ ] Implement audit logging
- [ ] Add metrics dan monitoring

---

## 📊 Project Stats

```
Language: Go
Framework: Gin
Database: PostgreSQL + Redis
API Style: RESTful
Authentication: JWT
Documentation: Swagger/OpenAPI
Container: Docker
```

---

**Last Updated**: 2024
**Go Version**: 1.25.0
**Status**: ✅ Production Ready

---

*Untuk pertanyaan atau saran, silakan buat issue di repository GitHub kami. Terima kasih telah menggunakan Koda B6 Backend!*