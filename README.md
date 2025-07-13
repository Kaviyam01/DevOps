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
