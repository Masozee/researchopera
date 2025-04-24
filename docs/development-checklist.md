# Research Opera Backend Development Checklist

## ✅ Completed
- Project structure scaffolded (DDD-inspired, modular Go backend)
- Domain models for all core entities (Tenant, User, Project, Publication, Dataset, Literature, Ethics, Grant, Sponsor, Compliance, Safety, Integrity, Audit)
- Many-to-many join tables for User↔Project and User↔Publication
- Database initialization utility (`internal/util/db.go`)
- Main entrypoint with automigration for all models
- Multi-tenancy middleware for Fiber (extract TenantID from JWT or subdomain)
- Middleware integrated into Fiber app initialization (`internal/api/server.go`)
- Sample tenant-scoped API handler scaffolded (`internal/api/project_handler.go`)

## ⏳ In Progress / Next Steps
- Wire up GORM DB instance to API layer (for handler access)
- Expand Project CRUD endpoints (Create, Update, Delete)
- Add GORM query scoping for tenant restriction
- Repository interfaces and implementations for each domain
- Application services/use cases for business logic
- API handlers/controllers for all modules
- Unit/integration tests for middleware and APIs
- Documentation for API endpoints and architecture
- Docker Compose integration for backend, DB, and frontend
- Database migration files and tooling
- Security hardening (JWT validation, roles, etc.)
- CI/CD pipeline setup

---

**Update this checklist as development progresses.**
