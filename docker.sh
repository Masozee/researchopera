#!/bin/bash
# Docker environment management script for Research Opera

# Set the version from VERSION file
VERSION=$(cat VERSION)
export VERSION

case "$1" in
  dev)
    echo "Starting Research Opera in DEVELOPMENT mode..."
    docker-compose -f docker-compose.dev.yml up "$2"
    ;;
  prod)
    echo "Starting Research Opera in PRODUCTION mode..."
    docker-compose up "$2"
    ;;
  build)
    echo "Building Research Opera images..."
    docker-compose build
    docker-compose -f docker-compose.dev.yml build
    ;;
  down)
    echo "Stopping Research Opera containers..."
    docker-compose down
    docker-compose -f docker-compose.dev.yml down
    ;;
  clean)
    echo "Cleaning up Research Opera containers and volumes..."
    docker-compose down -v
    docker-compose -f docker-compose.dev.yml down -v
    ;;
  *)
    echo "Research Opera Docker Management"
    echo "Usage: $0 [command] [options]"
    echo ""
    echo "Commands:"
    echo "  dev   - Start in development mode with hot-reload"
    echo "  prod  - Start in production mode"
    echo "  build - Build all container images"
    echo "  down  - Stop all containers" 
    echo "  clean - Stop all containers and remove volumes"
    echo ""
    echo "Options:"
    echo "  For 'dev' and 'prod' commands, you can specify additional docker-compose options."
    echo "  For example: $0 dev --build"
    ;;
esac
