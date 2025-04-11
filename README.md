# Real-Time Video Ranking System

This project is a backend service that tracks and ranks videos in real time based on user interactions like views, likes, comments, shares, and watch time. It uses Redis for fast scoring and exposes HTTP APIs for clients to send interaction data and retrieve rankings.

---

## Features

- Real-time scoring of videos based on weighted interactions
- Global and per-user video rankings stored in Redis
- Swagger-generated API documentation
- Unit tested and easy to extend

---

## Project Structure
```
├── cmd
│   ├── api # Main application entrypoint 
├── docs # Swagger documentation (generated)
├── internal
│   ├── model # Data models and types 
│   ├── redisdb # Database
│   ├── score # Scoring weight config 
│   ├── server # HTTP handlers and server logic 
```

---
## Prerequisites

- Go 1.20+
- Redis
- [Swag](https://github.com/swaggo/swag) (`go install github.com/swaggo/swag/cmd/swag@latest`)

---

## Makefile Commands

| Command        | Description                                              |
|----------------|----------------------------------------------------------|
| `make build`   | Build the Go binary (`main`)                             |
| `make run`     | Run the app directly with `go run`                       |
| `make test`    | Run all Go tests                                         |
| `make swagger` | Generate Swagger docs (`docs/`)                          |
| `make all`     | Build binary, run tests, and generate Swagger docs       |

---

## API Endpoints
Swagger UI is available at:
```
http://localhost:8080/swagger/index.html
```
To update docs after changes:

```
make swagger
```