# Redira API Documentation

## Overview

Redira provides a REST API for creating, managing, and analyzing shortened URLs.

The API supports:

- User authentication
- URL shortening
- Link management
- Redirect handling
- Analytics tracking
- Dashboard metrics

---

# Base URL

Local development:

```text
http://localhost:8080
```

API version:

```text
/api/v1
```

---

# Authentication

Redira uses JWT authentication.

Protected endpoints require:

```http
Authorization: Bearer <JWT_TOKEN>
```

JWT tokens are returned after successful login.

---

# API Endpoints

## Authentication

| Method | Endpoint | Description |
|---|---|---|
| POST | `/api/v1/auth/register` | Create a new user account |
| POST | `/api/v1/auth/login` | Authenticate user and receive JWT token |

---

## Links

| Method | Endpoint | Description |
|---|---|---|
| POST | `/api/v1/links` | Create a shortened URL |
| GET | `/api/v1/links` | Get user's links |
| GET | `/api/v1/links/{id}` | Get link details |
| PATCH | `/api/v1/links/{id}` | Update a link |
| DELETE | `/api/v1/links/{id}` | Delete a link |
| POST | `/api/v1/links/{id}/enable` | Enable a link |
| POST | `/api/v1/links/{id}/disable` | Disable a link |

---

## Redirects

| Method | Endpoint | Description |
|---|---|---|
| GET | `/r/{shortCode}` | Redirect to original URL |

The redirect engine:

1. Receives the short code
2. Looks up the link
3. Records analytics
4. Redirects the user

---

## Analytics

| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/v1/analytics/{id}` | Get link analytics |
| GET | `/api/v1/analytics/links/{id}` | Get detailed link analytics |
| GET | `/api/v1/analytics/links/{id}/timeline` | Get analytics timeline |

---

## Dashboard

| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/v1/dashboard` | Get analytics dashboard summary |

Authentication required.

---

## System

| Method | Endpoint | Description |
|---|---|---|
| GET | `/health` | Health check |
| GET | `/metrics` | Prometheus metrics |

---

# Example Request

Create a shortened URL:

```http
POST /api/v1/links
```

Headers:

```http
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json
```

Request body:

```json
{
  "url": "https://example.com"
}
```

---

# Example Response

```json
{
  "short_code": "abc123",
  "original_url": "https://example.com"
}
```

---

# Error Responses

Redira uses standard HTTP status codes.

| Status Code | Meaning |
|---|---|
| 200 | Successful request |
| 201 | Resource created |
| 400 | Invalid request |
| 401 | Authentication required |
| 403 | Permission denied |
| 404 | Resource not found |
| 500 | Internal server error |

---

# API Design Principles

Redira follows:

- REST API conventions
- JSON request and responses
- JWT-based authentication
- Versioned endpoints
- Standard HTTP methods
- Consistent error handling

---

# OpenAPI Specification

The complete machine-readable API specification is available at:

```text
docs/api/openapi.yaml
```

The OpenAPI specification can be used with:

- Swagger UI
- Postman
- API client generators
- Developer tooling