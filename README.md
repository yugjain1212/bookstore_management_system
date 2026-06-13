# Book Management System

A robust RESTful API for a Bookstore Management System built with **Go (Golang)**, **GORM**, and **MySQL**.

---

## 🚀 Features

- **CRUD Operations**: Complete control over bookstore inventory (Create, Read, Update, Delete books).
- **ORM (GORM)**: Structured database interactions and automatic schema migrations.
- **Router (Gorilla Mux)**: Dynamic request routing with clean, REST-compliant URL paths.

---

## 🛠️ Tech Stack

- **Language**: Go (1.24+)
- **HTTP Router**: Gorilla Mux
- **ORM**: GORM (v1)
- **Database**: MySQL

---

## 📁 Project Structure

```text
├── cmd
│   └── main
│       └── main.go         # Entry point of the application
├── pkg
│   ├── config
│   │   └── app.go          # Database connection configuration
│   ├── controllers
│   │   └── book-controller.go # Request handlers/business logic
│   ├── models
│   │   └── book.go         # GORM Book model and DB methods
│   ├── routes
│   │   └── bookstore-routes.go # API route definitions
│   └── utils
│       └── utils.go        # JSON body parsing helper
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Configuration & Installation

### 1. Prerequisites
- **Go** installed on your system.
- **MySQL Server** running locally or remotely.

### 2. Database Setup
Create a new database named `go-project` in MySQL:
```sql
CREATE DATABASE `go-project`;
```

### 3. Application Config
Open `pkg/config/app.go` and update the database connection string with your MySQL username, password, and port:
```go
// Replace with your credentials: username:password@tcp(host:port)/dbname
d, err := gorm.Open("mysql", "<username>:<password>@tcp(127.0.0.1:3306)/go-project?charset=utf8mb4&parseTime=True&loc=Local")
```

### 4. Running the Application
Download dependencies and start the server:
```bash
# Get dependencies
go mod tidy

# Run the project
go run ./cmd/main
```
The server will start listening on **`http://localhost:9010`**.

---

## 🔌 API Endpoints

| Method | Endpoint | Description | Request Body |
| :--- | :--- | :--- | :--- |
| **POST** | `/bookstore` | Add a new book | JSON (Name, Author, Publication) |
| **GET** | `/bookStores` | Get list of all books | *None* |
| **GET** | `/bookstore/{id}` | Get book by ID | *None* |
| **PUT** | `/bookstore/{id}` | Update book by ID | JSON (fields to update) |
| **DELETE** | `/bookstore/{id}` | Delete book by ID | *None* |

### Example Request Body (Create Book)
`POST /bookstore`
```json
{
  "name": "The Go Programming Language",
  "author": "Alan A. A. Donovan",
  "publication": "Addison-Wesley"
}
```
