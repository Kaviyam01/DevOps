# SimpleTimeService – Minimal Go Microservice

A lightweight HTTP service built in **Golang** using the **Gin framework**, designed for cloud-native environments. It returns the current UTC timestamp and the IP address of the client in a clean JSON format.

---

## Purpose

- Clean software development using Go
- Docker containerization with security best practices
- Minimal image footprint with multi-stage Docker builds
- Public deployment via DockerHub

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Kaviyam01/DevOps.git
cd DevOps/app
```

---

## Docker Instructions

### 2. Build the Image

```bash
docker build -t simpletimeservice .
```

### 3. Run the Container

```bash
docker run -p 8080:8080 simpletimeservice
```

### 4. Test the Service

Visit `http://localhost:8080` in your browser or open New Terminal and test by using the below command:

```bash
curl http://localhost:8080
```

---

## Pull from DockerHub (Optional)

You can run the SimpleTimeService directly from DockerHub without building the image locally:

```bash
docker pull kaviyavk/simpletimeservice:latest
docker run -p 8080:8080 kaviyavk/simpletimeservice
```

---

## Project Structure

```
DevOps/
├── app/
│   ├── Dockerfile       # Secure, multi-stage build
│   ├── go.mod          # Module definition
│   ├── go.sum          # Dependency checksums
│   └── main.go         # Gin web service
└── README.md           # Project documentation
```

---

### Testing

```bash
# Test the endpoint
curl http://localhost:8080

# Output

Accessing the root path (`/`) will return:

```json
{
  "timestamp": "2025-07-13T12:45:00Z",
  "ip": "127.0.0.1"
}
```
