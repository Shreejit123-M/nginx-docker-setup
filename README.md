# 🚀 NGINX Reverse Proxy with Docker Compose (Go + Python Apps)

This project demonstrates how to use **Docker Compose** to run:
- A **Go (Golang)** web application (Service 1)
- A **Python Flask** web application (Service 2)
- An **NGINX Reverse Proxy** container that routes incoming traffic based on URL path

All services run inside containers and are accessible via a **single port (`localhost:8080`)** using NGINX as a reverse proxy.

---

## 📁 Project Structure

```
nginx-docker-setup/
├── docker-compose.yml
├── README.md
├── nginx/
│   ├── Dockerfile
│   └── nginx.conf
├── service_1/
│   ├── Dockerfile
│   └── main.go
├── service_2/
│   ├── Dockerfile
│   ├── app.py
│   └── requirements.txt
```

---

## ⚙️ Setup Instructions

Follow these steps to run the project:

### 1. ✅ Clone the Repository

```bash
git clone https://github.com/Shreejit123-M/nginx-docker-setup.git
cd nginx-docker-setup
```

### 2. ✅ Build and Start All Services

Use Docker Compose to build and run all containers:

```bash
docker compose up --build
```


## 🔁 How Routing Works

NGINX is configured to route incoming requests based on URL path prefixes:

```nginx
# nginx/nginx.conf

server {
    listen 80;

    location /service1/ {
        proxy_pass http://service1:8001/;
    }

    location /service2/ {
        proxy_pass http://service2:8002/;
    }
}
```

So:

- Visiting `/service1/` routes to Go app running on `service1:8001`
- Visiting `/service2/` routes to Python app running on `service2:8002`

All traffic is routed through NGINX which listens on port 8080 (mapped in docker-compose).

---

## 📋 Service Details

### 🔸 Go Web App (`service_1`)
- Uses `golang:1.20-alpine`
- Serves response at `/service1/`
- Simple handler that prints: `Hello from Go service!`

### 🔸 Python Flask App (`service_2`)
- Uses `python:3.9-slim`
- Serves response at `/service2/`
- Simple Flask app returning: `Hello from Python Flask service!`

---

## 📝 Docker Compose Setup

`docker-compose.yml` defines 3 services:
- `nginx` (reverse proxy)
- `service1` (Go app)
- `service2` (Flask app)

All services are on the same **bridge network** and accessible internally by name.

---

## ✅ Logs Example

```bash
docker compose logs nginx
```

Example log:

```
"GET /service1/ HTTP/1.1" 200
"GET /service2/ HTTP/1.1" 200
```

Logs include:
- Timestamp
- Requested URL
- Status Code

---

## 🏆 Bonus Features Implemented

- ✅ Reverse proxy routing based on path (`/service1`, `/service2`)
- ✅ NGINX access logging with timestamp and path
- ✅ All services run on **a single port (8080)** using NGINX
- ✅ `docker-compose up --build` to start entire system
- 🛠️ Optional: Healthcheck can be added in `docker-compose.yml` if required

---

## 🧼 To Stop the System

```bash
docker compose down
```

---

## 👨‍💻 Author

**Shreejit Magadum**  
DevOps Intern Assignment - DPDzero  
Date: June 2025

---