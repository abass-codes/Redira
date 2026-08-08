# Redira

Production-grade URL shortener built with Go, PostgreSQL, Redis, and modern engineering practices.

Redira provides fast URL redirection, link management, analytics, authentication, and production-ready infrastructure.

---

# Features

## Core Platform

- URL shortening
- Fast redirects
- User authentication
- Link management
- Link expiration and controls
- Analytics tracking
- Dashboard metrics


## Engineering Features

- Redis caching
- Rate limiting
- JWT authentication
- Security middleware
- Request logging
- Request tracing
- Prometheus metrics
- Automated testing
- Database migrations
- CI/CD workflow


---

# Tech Stack

## Backend

- Go
- Gin
- PostgreSQL
- Redis


## Frontend

- Next.js
- React
- TypeScript


## Infrastructure

- Docker
- Docker Compose
- GitHub Actions
- Prometheus


---

# Architecture

```
                Users

                  |

                HTTPS

                  |

             Redira API

                  |

        -------------------

        |                 |

   PostgreSQL          Redis

   Database            Cache
```

---

# Project Structure

```
Redira

├── cmd/
│   └── server/

├── internal/

│   ├── auth/

│   ├── links/

│   ├── redirect/

│   ├── analytics/

│   ├── database/

│   └── middleware/


├── frontend/

├── migrations/

├── deployments/

├── docs/

└── docker-compose.yml
```

---

# Running Locally

## Start Infrastructure

```bash
docker compose up -d
```


## Run Backend

```bash
go run cmd/server/main.go
```


## Run Frontend

```bash
cd frontend

npm install

npm run dev
```

---

# Testing

Run:

```bash
go test ./...
```

Coverage:

```bash
go test ./... -cover
```

---

# Build

Backend:

```bash
go build ./...
```

Frontend:

```bash
cd frontend

npm run build
```

---

# Documentation

API:

```
docs/api/
```

Database:

```
docs/database/
```

Infrastructure:

```
docs/infrastructure/
```

Reliability:

```
docs/reliability/
```

Developer setup:

```
docs/development/setup.md
```

---

# Engineering Highlights

Redira demonstrates:

- Backend API design
- Database engineering
- Distributed caching
- Authentication systems
- Observability
- Testing strategy
- Production deployment practices
- Reliability engineering


---

# License

MIT License