# Redira Database Engineering

## Overview

Redira uses PostgreSQL as the primary database for storing application data.

The database layer is designed around:

- Data consistency
- Reliable schema changes
- Efficient queries
- Scalable storage patterns

Database engineering practices include:

- Version-controlled migrations
- Database indexing
- Foreign key relationships
- Query optimization

---

## Database Technology

Primary database:

- PostgreSQL 17

Used for storing:

- User accounts
- Authentication data
- Short links
- Redirect information
- Analytics events

---

## Database Schema

### Users

Stores registered users.

Main fields:

- id
- email
- password_hash
- created_at
- updated_at

Purpose:

- User authentication
- Account management
- Link ownership

---

### Links

Stores shortened URLs created by users.

Main fields:

- id
- user_id
- short_code
- original_url
- active
- created_at
- updated_at

Relationship:

```
User → Many Links
```

Purpose:

- Store original URLs
- Generate short links
- Manage link ownership
- Support redirects

---

### Analytics Events

Stores redirect activity.

Main fields:

- id
- link_id
- ip_address
- user_agent
- created_at

Relationship:

```
Link → Many Analytics Events
```

Purpose:

- Track redirect activity
- Support analytics dashboards
- Store usage information

---

## Database Migrations

Database changes are managed using SQL migration files.

Migration structure:

```
migrations/

001_initial_schema.sql

002_indexes.sql
```

Benefits:

- Repeatable database setup
- Version-controlled schema changes
- Safer deployments
- Consistent environments

---

## Initial Schema Migration

The initial migration creates:

- Users table
- Links table
- Analytics events table

Responsibilities:

- Create database tables
- Define relationships
- Add constraints
- Establish initial structure

---

## Indexing Strategy

Indexes are added for frequently accessed queries.

Current indexes:

- Short code lookup
- User link retrieval
- Link creation sorting
- Analytics queries

Important indexed fields:

```
links(short_code)

links(user_id)

analytics_events(link_id)

analytics_events(created_at)
```

Benefits:

- Faster redirects
- Improved dashboard performance
- More efficient analytics queries

---

## Database Reliability

Redira uses database reliability practices:

### Constraints

Database constraints protect data integrity.

Examples:

- Unique email addresses
- Unique short codes
- Foreign key relationships


### Foreign Keys

Relationships between tables are enforced using foreign keys.

Examples:

- Users own links
- Links contain analytics events


### Controlled Migrations

Schema changes are introduced through migration files instead of manual database changes.

Benefits:

- Safer updates
- Easier deployment
- Consistent development environments

---

## Development Database

Local development uses PostgreSQL through Docker.

Services:

```
PostgreSQL Container

Redis Container
```

Database configuration is provided through:

```
DATABASE_URL
```

Example:

```env
DATABASE_URL=postgres://USER:PASSWORD@HOST:5432/redira
```

---

## Production Database Practices

Production databases should use:

- Managed PostgreSQL services
- Automated backups
- Secure credentials
- Restricted database access
- Database monitoring

---

## Future Improvements

Future database improvements:

- Automated migrations in CI/CD
- Database rollback support
- Query performance monitoring
- Connection pool optimization
- Read replicas
- Advanced backup strategies