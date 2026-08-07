# Redira Continuous Integration

## Overview

Redira uses Continuous Integration (CI) to automatically verify code changes before they are merged.

The CI pipeline ensures that new changes maintain application reliability and do not introduce unexpected issues.

---

## CI Workflow

Redira uses GitHub Actions to automatically run checks whenever code is pushed or a pull request is created.

The workflow validates:

- Backend code
- Frontend code
- Automated tests
- Production builds

---

## Automated Checks

### Backend Tests

The CI pipeline runs:

```bash
go test ./...
```

This verifies:

- Unit tests
- Service tests
- API tests
- Integration tests

---

### Backend Build

The application is compiled using:

```bash
go build ./...
```

This ensures:

- The code compiles successfully
- Dependencies are correctly configured
- Packages build correctly

---

### Frontend Build

The frontend production build is verified using:

```bash
npm run build
```

This checks:

- TypeScript compilation
- React application building
- Production readiness

---

### Docker Build

The production Docker configuration is validated using:

```bash
docker compose -f docker-compose.prod.yml build
```

This confirms:

- Docker configuration is valid
- The API container builds successfully
- The production environment can be recreated

---

## Pull Request Process

Development workflow:

1. Create a feature branch

```bash
git checkout -b feature/new-feature
```

2. Implement changes

3. Run tests locally

```bash
go test ./...
```

4. Push changes

5. Open a pull request

6. CI automatically validates the changes

7. Merge after all checks pass

---

## Benefits

Continuous Integration provides:

- Faster feedback
- Safer code changes
- Consistent testing
- Reliable releases
- Improved development workflow

---

## Future Improvements

Future CI improvements include:

- Automated deployments
- Security scanning
- Dependency checks
- Coverage enforcement
- Docker security scanning
```