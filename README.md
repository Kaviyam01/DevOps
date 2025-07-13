# SimpleTimeService – Minimal Go Microservice

A lightweight HTTP service built in **Golang** using the **Gin framework**, designed for cloud-native environments. It returns the current UTC timestamp and the IP address of the client in a clean JSON format.

---

## Purpose

- Clean software development using Go
- Docker containerization with security best practices
- Minimal image footprint with multi-stage Docker builds
- Public deployment via DockerHub
- Clear documentation for team deployment and testing

---

## Example Output

Accessing the root path (`/`) will return:

```json
{
  "timestamp": "2025-07-13T12:45:00Z",
  "ip": "127.0.0.1"
}

---

#### 1. Getting Started – Clone Repo

```markdown
## Getting Started

### Clone the Repository

```bash
git clone https://github.com/Kaviyam01/DevOps.git
cd DevOps/app

---

#### 2. Docker Build Instructions

```markdown
## Docker Instructions

### Build the Image

```bash
docker build -t simpletimeservice .

### Run the Container

```bash
docker run -p 8080:8080 simpletimeservice

### Visit http://localhost:8080 on browser or use:

```bash
curl http://localhost:8080

---

#### 3. Pull from DockerHub (optional)

You can run the SimpleTimeService directly from DockerHub without building the image locally:

```bash
docker pull kaviyavk/simpletimeservice:latest
docker run -p 8080:8080 kaviyavk/simpletimeservice

---

#### 4. Project Structure

```markdown
## Project Structure

DevOps/
├── app/
│ ├── main.go # Gin web service
│ ├── Dockerfile # Secure, multi-stage build
│ ├── go.mod # Module definition
│ ├── go.sum # Dependency checksums
└── README.md # Project documentation

