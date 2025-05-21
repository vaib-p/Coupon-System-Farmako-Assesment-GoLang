# 🏷️ Coupon Management System

A Go-based RESTful API service to manage discount coupons for an eCommerce or healthcare platform. This service allows creating, retrieving, and validating coupons with multiple constraints like usage type, applicable time windows, order value thresholds, and more.

---

## 🌐 Hosted URL

Live Project: [https://coupon-system-farmako-assesment-golang-production.up.railway.app](https://coupon-system-farmako-assesment-golang-production.up.railway.app)  
Swagger Docs: [https://coupon-system-farmako-assesment-golang-production.up.railway.app/swagger/index.html](https://coupon-system-farmako-assesment-golang-production.up.railway.app/swagger/index.html)

---


## 📁 Project Structure

```
.
├── main.go
├── models/
│   └── coupon.go
├── handlers/
│   └── coupon_handler.go
├── repository/
│   └── db.go
├── cache/
│   └── coupon_cache.go
├── docs/
│   └── swagger.json / swagger.yaml
└── go.mod
```

---

## 🔧 Technologies Used

- Go (Golang)
- Gin Web Framework
- GORM (Go ORM)
- PostgreSQL
- Swagger (Swaggo)
- Go-Cache (In-memory caching)
- Docker (optional)

---

## 🧩 Features

- Create new coupons
- List all available coupons
- Validate coupons with:
  - Expiry date check
  - Minimum order value
  - Valid time window
- Support for different:
  - **Usage types** (`one_time`, `multi_use`, `time_based`)
  - **Discount types** (`flat`, `percentage`)
- Caching validation responses for performance

---

## 📦 Coupon Model

```go
type Coupon struct {
    CouponCode             string         `gorm:"primaryKey"`
    ExpiryDate             time.Time
    UsageType              UsageType
    ApplicableMedicineIDs  pq.StringArray `gorm:"type:text[]"`
    ApplicableCategories   pq.StringArray `gorm:"type:text[]"`
    MinOrderValue          float64
    ValidTimeWindow        pq.StringArray `gorm:"type:text[]"` // ["09:00", "21:00"]
    TermsAndConditions     string
    DiscountType           DiscountType
    DiscountValue          float64
    MaxUsagePerUser        int
}
```

---

## 🔌 API Endpoints

### ✅ Create Coupon

- **URL:** `POST /coupons`
- **Request Body:** `models.Coupon`
- **Success Response:** `201 Created`
- **Failure Response:** `400 Bad Request | 500 Internal Server Error`

### 📋 Get All Coupons

- **URL:** `GET /coupons`
- **Response:** `200 OK` — Returns list of coupons

### 🎯 Validate Coupon

- **URL:** `POST /coupons/validate`
- **Request Body:**

```json
{
  "coupon_code": "SAVE20",
  "total": 150.75,
  "time": "14:30"
}
```

- **Response:**

```json
{
  "is_valid": true,
  "message": "Coupon is valid",
  "discount": {
    "type": "percentage",
    "value": 20
  }
}
```

---

## 🚀 Running Locally

### 1️⃣ Clone Repository

```bash
git clone https://github.com/yourusername/coupon-system.git
cd coupon-system
```

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

### 3️⃣ Run Application

```bash
go run main.go
```
---
## 🐳 Run with Docker

### 1️⃣ Build the image

```bash
docker build -t coupon-system .
```

### 2️⃣Run the container
```bash
docker run -d -p 8080:8080 --env-file .env coupon-system
```


## 🧪 Swagger Documentation

Swagger annotations are added for all routes.

### 📚 Generate Swagger

```bash
swag init
```

Then access via:  
`http://localhost:8080/swagger/index.html`

---
Also Available at:
---

`https://coupon-system-farmako-assesment-golang-production.up.railway.app/swagger/index.html`

---
## 🧠 Caching

Validation responses are cached using `go-cache` with a default expiration time and automatic cleanup.  
This reduces DB hits and improves performance.

---



---

## 👤 Author

**Your Name**  
Backend Developer | Go & Java Enthusiast  
[LinkedIn](https://linkedin.com/in/vaib-p) • [GitHub](https://github.com/vaib-p)

---

## 📜 License

This project is licensed under the MIT License.
