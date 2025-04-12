# Microservices Project
### Clean Architecture-Based Microservices E-Commerce Platform

## 📌 Objective

This project focuses on designing and implementing a basic **e-commerce platform** using **microservices architecture**. The solution adheres to **clean architecture** principles and leverages **Go with Gin** for HTTP handling and any preferred database for persistence. 
This project is a collection of microservices designed for handling **Orders** and **Inventory**. Each service has its own database, REST APIs, and distinct configurations.


## Project Structure

```
├── order_service                                    # Order service handling order logic
├── inventory_service                                # Inventory service for managing inventory
├── api-gateway                                      # API gateway for single access point for different services.
├── MicroserviceExample.postman_collection.json      # Postman collection. Use this to test the application.
```

## 📦 Microservices Overview

The system consists of three services:

### 1. 🚪 API Gateway (Gin)
- Acts as the single entry point.
- Handles:
  - Routing requests to appropriate services

### 2. 📦 Inventory Service (Gin + DB)
- Manages productdata.
- Responsible for:
  - Storing product information
  - Tracking stock levels and prices
- Provides **CRUD** operations.

**Endpoints:**
- `POST /products` – Create a new product  
- `GET /products/:id` – Retrieve product by ID  
- `PATCH /products/:id` – Update product details  
- `DELETE /products/:id` – Remove a product  
- `GET /products` – List all products (supports filtering and pagination)

### 3. 🧾 Order Service (Gin + DB)
- Handles:
  - Order creation
  - Order status updates
- Stores:
  - Each order information

**Endpoints:**
- `POST /orders` – Create a new order  
- `GET /orders/:id` – Retrieve order by ID  
- `PATCH /orders/:id` – Update order status (pending, completed, cancelled)  
- `GET /orders` – List all user orders  

---


## 👨‍🎓 Author

- **Name:** Temutjin Koszhanov  
- **Course:** Advanced Programming II  
