# Redira Developer Setup

## Overview

This guide explains how to run Redira locally for development.

---

## Requirements

Install:

- Go 1.26+
- Docker
- Docker Compose
- Node.js
- npm

Verify:

```bash
go version
docker --version
docker compose version
node --version
npm --version
```

---

## Clone Repository

```bash
git clone https://github.com/abass-codes/Redira.git

cd Redira
```

---

## Environment Setup

Create environment file:

```bash
cp .env.example .env
```

Required variables:

```env
APP_ENV=development
SERVER_PORT=8080
DATABASE_URL=postgres://USER:PASSWORD@localhost:5432/redira
REDIS_URL=redis://localhost:6379
JWT_SECRET=your_secret
```

---

## Start Services

Start PostgreSQL and Redis:

```bash
docker compose up -d
```

Check containers:

```bash
docker ps
```

---

## Run Backend

Install dependencies:

```bash
go mod download
```

Start API:

```bash
go run cmd/server/main.go
```

API:

```text
http://localhost:8080
```

Health check:

```bash
curl http://localhost:8080/health
```

---

## Run Frontend

```bash
cd frontend

npm install

npm run dev
```

Frontend:

```text
http://localhost:3000
```

---

## Testing

Run tests:

```bash
go test ./...
```

Run coverage:

```bash
go test ./... -cover
```

---

## Build

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

## Docker Production Build

Build:

```bash
docker compose -f docker-compose.prod.yml build
```

Run:

```bash
docker compose -f docker-compose.prod.yml up -d
```

---

## Project Structure

```text
Redira

├── cmd/
├── internal/
├── frontend/
├── migrations/
├── docs/
└── docker-compose.prod.yml
```

---

## API Documentation

API reference:

```text
docs/api/api.md
```

OpenAPI specification:

```text
docs/api/openapi.yaml
```