# Axilock Backend

A backend service for Axilock push protection application, built with Go and modern microservices architecture.

## Prerequisites

- Go 1.24.0 or higher
- Docker and Docker Compose
- PostgreSQL 16
- Redis 7.4.1
- Git

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/axilock/axi-backend.git
cd axi-backend
```

### 2. Configure Environment

Copy the example environment file and modify it according to your needs:

```bash
cp env.example .env
```

Edit the `.env` file with your specific configuration:
- Database connection details
- Redis configuration
- JWT secrets
- Other service-specific configurations

### 2.1 Configure Environment Variables

Fill the .env file with your specific configuration:

```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://admin:<Enter your database password>@<Enter your database address>:5432/axilockdb?sslmode=disable
HTTP_SERVER_ADDRESS=0.0.0.0:8080
GRPC_SERVER_ADDRESS=0.0.0.0:8090
RUNNING_ENV=<Enter your running environment>
REDIS_ADDR=<Enter your redis address>
REDIS_PASS=<Enter your redis password>
GITHUB_APP_ID=<Enter your github app id>
GITHUB_CLIENT_SECRET=<Enter your github client secret>
GITHUB_CLIENT_ID=<Enter your github client id>
DISCORD_WEBHOOK=<Enter your discord webhook>
```

You will also need the private key file for the github app to be stored in the root directory of the project and name the file as `axilock.pem`.

### 3. Start Services with Docker Compose

The application uses Docker Compose to manage services. Start all services with:

```bash
docker-compose up -d
```

This will start:
- PostgreSQL database (port 5432)
- Redis cache (port 6379)
- Backend service (ports 8080 and 8090)

### 4. Build and Run Locally

Alternatively, you can build and run the application locally:

```bash
# Build the project
go build -o axilock .

# Run the application
./axilock
```

### 5. Run Tests

To run the test suite:

```bash
make test
```

## Development

### Code Style and Linting

The project uses golangci-lint for code style enforcement. Run:

```bash
make lint
```

### API Documentation

The backend service provides RESTful APIs with endpoints documented in the codebase. The main entry points are:

- HTTP API: ``http://localhost:8080``
- gRPC Service: ``http://localhost:8090``

### Database Migrations

Database migrations are managed in the `migrations/` directory. To apply migrations:

```bash
docker-compose exec db psql -U admin -d axilockdb -f /path/to/migration.sql
```

## Security

The application uses:

- PostgreSQL to store commit metadata and store users details to map them to github users
- Redis to store cli token for session data.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the Apache2 License - see the LICENSE file for details.

## Contact

For support or questions, please open an issue in the [GitHub repository](https://github.com/axilock/axilock-backend).
