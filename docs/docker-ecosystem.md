# Research Opera Docker Ecosystem Instructions

## Overview
This document describes how to use the advanced Docker ecosystem for Research Opera, designed for efficient resource usage, maximum performance, security, and automatic updates with code changes.

## Production Environment (`docker-compose.yml`)
- **Resource Management**: CPU and memory limits for each service
- **Security Enhancements**: Non-root users, read-only filesystems, Docker secrets for passwords
- **Network Isolation**: Separate networks for frontend, backend, and database communication
- **Health Monitoring**: Healthchecks for all services
- **Logging Controls**: JSON logging with rotation

## Development Environment (`docker-compose.dev.yml`)
- **Hot-Reloading**: Code changes instantly reflected
  - Backend: Uses Air tool for Go hot-reload
  - Frontend: Uses Next.js dev server
- **Volume Mounting**: Local directories mounted for live updates
- **Debugging**: Delve debugger enabled for backend

## Key Files
- `docker-compose.yml`: Production configuration
- `docker-compose.dev.yml`: Development configuration with hot-reload
- `backend/Dockerfile`, `frontend/Dockerfile`: Production Dockerfiles
- `backend/Dockerfile.dev`, `frontend/Dockerfile.dev`: Development Dockerfiles
- `backend/.air.toml`: Air config for Go hot-reloading
- `secrets/db_password.txt`: Secure database password
- `docker.sh`: Shell script for managing environments

## Usage

### Development Mode
```
sh docker.sh dev
```

### Production Mode
```
sh docker.sh prod
```

### Other Commands
- `sh docker.sh build` - Build all images
- `sh docker.sh down` - Stop all containers
- `sh docker.sh clean` - Remove containers and volumes

---

For more details, see the README or reach out to the maintainers.
