# 🍽️ Restaurant Reservation System (Go + Gin)

## 📖 Overview
The **Restaurant Reservation System** is a RESTful API built using the **Go Gin Framework** and **PostgreSQL**.  
It provides endpoints to manage users, restaurants, menus, and orders with role-based authentication (Admin & User).

This project aims to create a simple, reliable, and easy-to-use backend for restaurant reservation and management.

---

## ⚙️ Features

### 🧑‍💼 Admin
- Create, update, and delete restaurants  
- Manage menu items  
- Assign admin roles  
- View and manage all orders  

### 👤 User
- Register and log in with JWT authentication  
- View restaurant lists and menus  
- Create and view their own orders  

### 🌐 Public Access
- Anyone can view restaurants and menus without logging in  

---

## 🧱 Tech Stack

| Component | Technology |
|------------|-------------|
| Language | Go (Golang) |
| Framework | Gin Web Framework |
| Database | PostgreSQL |
| Authentication | JWT (JSON Web Token) |
| Containerization | Docker & Docker Compose |
| Base OS | Alpine Linux |

---

## 🧩 API Endpoints Summary

| Endpoint | Method | Description | Access |
|-----------|---------|-------------|---------|
| `/api/users/register` | **POST** | Register a new user | Public |
| `/api/users/login` | **POST** | Login and get JWT token | Public |
| `/api/restaurants` | **GET** | Get list of all restaurants | Public |
| `/api/restaurants/:id` | **GET** | Get restaurant details | Public |
| `/api/restaurants/:id/menu` | **GET** | Get menu items for a restaurant | Public |
| `/api/admin/restaurants` | **POST** | Create restaurant | Admin |
| `/api/admin/restaurants/:id` | **PUT** | Update restaurant | Admin |
| `/api/admin/restaurants/:id` | **DELETE** | Delete restaurant | Admin |
| `/api/admin/menu` | **POST** | Create menu item | Admin |
| `/api/admin/menu/:id` | **PUT** | Update menu item | Admin |
| `/api/admin/menu/:id` | **DELETE** | Delete menu item | Admin |
| `/api/orders` | **POST** | Create order | User |
| `/api/orders/mine` | **GET** | View user’s orders | User |
| `/api/admin/orders/:id/status` | **PUT** | Update order status | Admin |

---

## 🧰 Environment Variables

| Variable | Description | Example |
|-----------|-------------|----------|
| `DB_HOST` | Database host | postgres |
| `DB_PORT` | Database port | 5432 |
| `DB_USER` | Database user | postgres |
| `DB_PASSWORD` | Database password | ggwp |
| `DB_NAME` | Database name | fooddb |
| `PORT` | App port | 8080 |
| `JWT_SECRET` | Secret key for JWT | abc |

---

## 🐳 Docker Setup

### 1️⃣ Build & Run
```bash
docker compose up --build
```

This will:
- Build your Go app using the included Dockerfile  
- Start PostgreSQL 15  
- Run the REST API on port **8080**

### 2️⃣ Access
- API: [http://localhost:8080](http://localhost:8080)  
- Postgres: `localhost:5432`

---

## 🗂️ Project Structure

```
FINAL-PROJECT/
├── config/                 # Configuration loader
├── controllers/            # Gin route controllers (Auth, Restaurant, Menu, Orders)
├── db/
│   └── sql_migrations/     # SQL migration files
├── repository/             # Database query logic
├── structs/                # Data models and DTOs
├── Dockerfile              # Multi-stage build for Go app
├── docker-compose.yml      # Defines app + Postgres services
├── go.mod / go.sum         # Go module dependencies
├── main.go                 # Application entry point
└── README.md
```

---

## 🧾 Example Usage (with curl)

### Register a user
```bash
curl -X POST http://localhost:8080/api/users/register   -H "Content-Type: application/json"   -d '{"name":"Budi","email":"budi@example.com","password":"secret123"}'
```

### Login as admin
```bash
curl -X POST http://localhost:8080/api/users/login   -H "Content-Type: application/json"   -d '{"email":"admin@food.local","password":"admin123"}'
```

### Create a restaurant (Admin only)
```bash
curl -X POST http://localhost:8080/api/admin/restaurants   -H "Authorization: Bearer <ADMIN_TOKEN>"   -H "Content-Type: application/json"   -d '{"name":"Bakso Mantap","address":"Jl. Sudirman No.10","phone":"021-123456"}'
```

### Create an order (User only)
```bash
curl -X POST http://localhost:8080/api/orders   -H "Authorization: Bearer <USER_TOKEN>"   -H "Content-Type: application/json"   -d '{"restaurant_id":1,"items":[{"menu_item_id":1,"quantity":2}]}'
```

---

## 🚀 Running Locally (without Docker)

```bash
go mod download
go run main.go
```

Make sure PostgreSQL is running locally with the same environment variables.

---

## 🔒 Authentication
- JWT tokens are generated on login.  
- Include in headers:
  ```
  Authorization: Bearer <token>
  ```

---

## 🧾 Migration Notes
Place your migration SQL at:
```
db/sql_migrations/migrate.sql
```
Example:
```sql
-- +migrate Up
CREATE TABLE restaurants (...);

-- +migrate Down
DROP TABLE restaurants;
```

---
