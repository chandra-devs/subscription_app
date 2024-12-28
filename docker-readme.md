## 🐳 Docker Setup

### Prerequisites
- Docker
- Docker Compose

### Quick Start with Docker

1. **Copy environment files**
```bash
cp .env.example .env
cp .env.docker .env.docker
```

2. **Update environment variables**
Edit `.env.docker` with your secure credentials:
```env
DB_PASSWORD=your_secure_password
JWT_SECRET=your_jwt_secret_key
```

3. **Build and run with Docker Compose**
```bash
docker-compose up --build
```

The application will be available at `http://localhost:3000`

### Docker Commands

- **Start services**
```bash
docker-compose up -d
```

- **Stop services**
```bash
docker-compose down
```

- **View logs**
```bash
docker-compose logs -f
```

- **Rebuild containers**
```bash
docker-compose up --build
```

### Container Structure
- **app**: Go application container
  - Exposes port 3000
  - Depends on postgres container
  - Auto-restarts unless stopped manually

- **postgres**: PostgreSQL database container
  - Exposes port 5432
  - Persists data using Docker volumes
  - Auto-restarts unless stopped manually

### Development with Docker
For development, you can use the hot-reload feature by installing [Air](https://github.com/cosmtrek/air):

1. Install Air globally:
```bash
go install github.com/cosmtrek/air@latest
```

2. Run the development server:
```bash
air
```

### Production Deployment
For production deployment, make sure to:
1. Use strong passwords in `.env.docker`
2. Change default PostgreSQL credentials
3. Configure proper SSL/TLS certificates
4. Set up proper logging
5. Configure backup strategy for the database volume
