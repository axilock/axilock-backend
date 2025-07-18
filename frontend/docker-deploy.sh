#!/bin/bash

# Default values
IMAGE_NAME="axilock-frontend"
IMAGE_TAG="latest"
REGISTRY=""  # Empty for Docker Hub, or e.g., "gcr.io/your-project" for GCR
API_URL="https://api.axilock.ai/v1/"

# Parse command line arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    --api-url)
      API_URL="$2"
      shift 2
      ;;
    --tag)
      IMAGE_TAG="$2"
      shift 2
      ;;
    --registry)
      REGISTRY="$2"
      shift 2
      ;;
    --help)
      echo "Usage: $0 [options]"
      echo "Options:"
      echo "  --api-url URL     Set the API base URL (default: https://api.axilock.ai/v1/)"
      echo "  --tag TAG         Set the image tag (default: latest)"
      echo "  --registry REG    Set the container registry (default: Docker Hub)"
      echo "  --help            Show this help message"
      exit 0
      ;;
    *)
      echo "Unknown option: $1"
      exit 1
      ;;
  esac
done

# Set full image name with registry if provided
if [ -n "$REGISTRY" ]; then
  FULL_IMAGE_NAME="$REGISTRY/$IMAGE_NAME:$IMAGE_TAG"
else
  FULL_IMAGE_NAME="$IMAGE_NAME:$IMAGE_TAG"
fi

echo "Building Docker image with API URL: $API_URL"
echo "Image will be tagged as: $FULL_IMAGE_NAME"

# Build the Docker image
docker build \
  --build-arg REACT_APP_API_BASE_URL="$API_URL" \
  -t "$FULL_IMAGE_NAME" \
  .

# Check if build was successful
if [ $? -ne 0 ]; then
  echo "Error: Docker build failed"
  exit 1
fi

echo "Docker image built successfully: $FULL_IMAGE_NAME"

# Ask if the user wants to push the image
read -p "Do you want to push the image to the registry? (y/n): " PUSH_IMAGE

if [[ "$PUSH_IMAGE" =~ ^[Yy]$ ]]; then
  echo "Pushing image to registry..."
  docker push "$FULL_IMAGE_NAME"
  
  if [ $? -ne 0 ]; then
    echo "Error: Failed to push image to registry"
    exit 1
  fi
  
  echo "Image pushed successfully to: $FULL_IMAGE_NAME"
else
  echo "Skipping image push"
fi

echo "Deployment script completed"

# Provide instructions for running the container
echo ""
echo "To run the container locally:"
echo "docker run -d -p 8080:80 $FULL_IMAGE_NAME"
echo ""
echo "The application will be available at: http://localhost:8080"
