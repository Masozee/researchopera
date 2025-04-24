# Research Opera

**Version:** 0.0.0-averroes

Research Opera is a platform for managing research projects and collaborations across institutions. It provides a modern full-stack architecture with a Go backend, Next.js frontend (with Tailwind CSS, shadcn/ui, and Radix UI), and a MySQL database, all orchestrated with Docker Compose.

## Project Structure

- `/backend`: Go REST API service
- `/frontend`: Next.js app with modern UI libraries
- `/db`: MySQL database (managed via Docker Compose)

## Getting Started

### Prerequisites
- Docker & Docker Compose
- Node.js (for local frontend dev)
- Go (for local backend dev)

### Setup
1. **Clone the repository:**
   ```sh
   git clone <your-repo-url>
   cd researchopera
   ```
2. **Initialize Backend (if not done):**
   ```sh
   cd backend
   go mod init backend
   go get -u github.com/go-sql-driver/mysql
   cd ..
   ```
3. **Initialize Frontend (if not done):**
   ```sh
   npx create-next-app@latest frontend --ts --eslint --tailwind --app --src-dir --import-alias "@/*" --no-interactive
   # Setup shadcn/ui and Radix UI as needed
   ```
4. **Start all services:**
   ```sh
   docker-compose up --build
   ```

## Useful Commands
- `docker-compose up --build` — Build and start all services
- `docker-compose down` — Stop all services

## License
MIT
