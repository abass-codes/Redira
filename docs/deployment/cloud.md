# Redira Cloud Deployment
## Architecture

Frontend

↓

Go API

↓

PostgreSQL + Redis

# Production Services
## API

- Go Gin server
- Docker container
- REST API
- JWT authentication
- Production logging


## Database

Managed PostgreSQL

Used for:

- User accounts
- Short links
- Click events
- Analytics data


## Cache

Managed Redis

Used for:

- URL redirect caching
- Rate limiting
- Fast lookups

# Environment Variables

Redira requires the following environment variables:


| Variable | Description |
|---|---|
| APP_NAME | Application name |
| APP_ENV | Runtime environment (development or production) |
| SERVER_PORT | API server port |
| DATABASE_URL | PostgreSQL connection string |
| REDIS_URL | Redis connection string |
| JWT_SECRET | JWT authentication secret |


Example production configuration:


```env
APP_NAME=Redira

APP_ENV=production

SERVER_PORT=8080

DATABASE_URL=postgres://USER:PASSWORD@HOST:5432/redira

REDIS_URL=redis://HOST:6379

JWT_SECRET=production_secret
```

# Production Deployment
## Build Docker Image

Build the production API image:


```bash
docker build -f deployments/docker/Dockerfile .
```

## Start Production Environment

Run the production stack:


```bash
docker compose -f docker-compose.prod.yml up
```

## Verify API Health

Check API availability:

```bash
curl http://localhost:8080/health
```

Expected response:


```json
{
  "status": "ok"
}
```

## Verify Service Readiness

Check database and Redis connections:

```bash
curl http://localhost:8080/health/ready
```

Expected response:

```json
{
  "database": "ok",
  "redis": "ok",
  "status": "ready"
}
```

# Cloud Deployment Architecture

Production architecture:

```
Users

↓

Frontend Hosting

(Vercel)

↓

HTTPS

↓

Go API

(Docker Container)

↓

-------------------------

PostgreSQL        Redis

(AWS RDS)     (ElastiCache)

-------------------------
```

# Recommended Cloud Services
## Frontend

Recommended:

- Vercel
- HTTPS enabled
- Production environment variables

## Backend API

Recommended:

- AWS ECS
- Docker deployment
- Load balancing
- Environment secrets


## Database

Recommended:

- AWS RDS PostgreSQL
- Automated backups
- Monitoring

## Cache

Recommended:

- AWS ElastiCache Redis
- Managed Redis instance


# Deployment Workflow
## Step 1 — Build Application

Build the Docker image:

```bash
docker build -f deployments/docker/Dockerfile .
```

## Step 2 — Configure Environment

Set production variables:

```env
APP_NAME

APP_ENV

SERVER_PORT

DATABASE_URL

REDIS_URL

JWT_SECRET
```

## Step 3 — Deploy Services

Deploy:

- Frontend
- API container
- PostgreSQL
- Redis


## Step 4 — Verify Deployment

Confirm:

- API health endpoint works
- PostgreSQL connection works
- Redis connection works
- Frontend communicates with API

# Production Checklist

- Docker image builds successfully

- API runs correctly in production mode

- PostgreSQL connection is working

- Redis connection is working

- Health endpoints are available

- Required environment variables are configured

- JWT authentication secrets are configured

- Frontend is connected to the API

- Deployment documentation is complete


# Future Improvements

- Custom domain
- HTTPS certificates
- Monitoring dashboards
- Error tracking
- Automated deployments
- Database backups
- Horizontal scaling
- Load balancing