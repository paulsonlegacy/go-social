# Go-Social API - Scalable Social Backend in Pure Go

## 🚀 Overview

This is a high-performance **social API** built in **pure Go**, without using backend frameworks like Gin. Designed as a **learning project** to explore **advanced backend engineering**, it covers a wide range of essential backend topics, including security, middleware, database engineering, caching, and cloud deployment.

## 📌 Features

### ✅ Core Backend Engineering

- **Custom Routing & Middleware** using `chi`
- **Request Validation & Authentication** (JWT, OAuth, CSRF, CORS)
- **Rate Limiting & Security** best practices

### 🗄️ Database Engineering

- **PostgreSQL/MySQL** database management
- **Database Migrations** with tools like `migrate`
- **Indexing & Query Optimization** for performance

### ⚡ Performance & Caching

- **Redis for caching** frequently accessed data
- **Rate limiting** to prevent abuse
- **Optimized database queries**

### 📦 DevOps & Deployment

- **Containerization with Docker**
- **Cloud Deployment on Google Cloud**
- **CI/CD pipelines** for automated deployment

### 📑 API Documentation

- **Swagger (OpenAPI 3.0)** for interactive API docs
- **Postman Collection** for easy testing

## 🛠️ Tech Stack

| Technology           | Purpose                           |
| -------------------- | --------------------------------- |
| **Go (Golang)**      | Backend programming               |
| **Chi Router**       | Lightweight routing               |
| **PostgreSQL/MySQL** | Database management               |
| **Redis**            | Caching and real-time performance |
| **Docker**           | Containerization                  |
| **Google Cloud**     | Scalable deployment               |
| **Swagger/OpenAPI**  | API documentation                 |

## 🔧 Getting Started

### Prerequisites

Make sure you have the following installed:

- **Go 1.21+**
- **PostgreSQL/MySQL**
- **Redis** (optional, for caching)
- **Docker** (for containerization)

### Installation & Setup

```sh
# Clone the repository
git clone https://github.com/paulsonlegacy/social-api.git
cd social-api

# Install dependencies
go mod tidy

# Run the application
go run .
```

### Running with Docker

```sh
docker-compose up --build
```

## 📖 API Documentation

- Swagger UI: `http://localhost:8080/docs`
- Postman Collection: [Download here](https://github.com/paulsonlegacy/go-social/postman_collection.json)

## 🤝 Contributing

We welcome contributions! Please fork the repository and submit a pull request.

---