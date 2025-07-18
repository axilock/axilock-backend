# Docker Deployment for Axilock Frontend

This repository contains Docker configuration files for deploying the Axilock frontend application.

## Files Overview

- `Dockerfile`: Multi-stage build file that builds the React application and serves it using Nginx
- `nginx.conf`: Nginx configuration for serving the Single Page Application (SPA)
- `.dockerignore`: Specifies files to exclude from the Docker build context
- `docker-deploy.sh`: Helper script for building and deploying the Docker image

## Configuration

The only required configuration is the API base URL, which can be set during the build process. By default, it points to `https://api.axilock.ai/v1/`.

## Building the Docker Image

You can build the Docker image using the provided script:

```bash
./docker-deploy.sh
```

### Custom API URL

To specify a custom API URL:

```bash
./docker-deploy.sh --api-url "https://your-custom-api.com/v1/"
```

### Custom Image Tag

To specify a custom image tag:

```bash
./docker-deploy.sh --tag "v1.0.0"
```

### Custom Registry

To push to a specific container registry:

```bash
./docker-deploy.sh --registry "gcr.io/your-project"
```

## Running the Container

After building the image, you can run it locally:

```bash
docker run -d -p 8080:80 axilock-frontend:latest
```

The application will be available at http://localhost:8080

## Environment Variables

The following environment variables can be set during the build process:

- `REACT_APP_API_BASE_URL`: The base URL for API requests (default: https://api.axilock.ai/v1/)

## Production Deployment

For production deployment, consider the following:

1. Use a specific version tag for your image
2. Push the image to a secure container registry
3. Deploy using Kubernetes, Docker Swarm, or a cloud service like AWS ECS, Google Cloud Run, or Azure Container Instances
4. Set up proper networking and security rules
5. Configure a domain name and TLS certificate

### Example Kubernetes Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: axilock-frontend
spec:
  replicas: 3
  selector:
    matchLabels:
      app: axilock-frontend
  template:
    metadata:
      labels:
        app: axilock-frontend
    spec:
      containers:
      - name: axilock-frontend
        image: your-registry/axilock-frontend:latest
        ports:
        - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: axilock-frontend-service
spec:
  selector:
    app: axilock-frontend
  ports:
  - port: 80
    targetPort: 80
  type: LoadBalancer
```

## Troubleshooting

If you encounter issues:

1. Check that the API URL is correctly set and accessible
2. Verify that the container is running: `docker ps`
3. Check container logs: `docker logs <container-id>`
4. Ensure port 8080 is not in use by another application when testing locally
