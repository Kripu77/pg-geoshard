#!/bin/bash

set -e

# Check if project ID is provided
if [ -z "$1" ]; then
    echo "Error: Please provide your GCP project ID as an argument"
    echo "Usage: ./apply.sh PROJECT_ID"
    exit 1
fi

PROJECT_ID=$1

# Ensure gcloud is authenticated
echo "Checking gcloud authentication..."
if ! gcloud auth application-default print-access-token &>/dev/null; then
    echo "Authenticating with Google Cloud..."
    gcloud auth application-default login
fi

# Set the project
gcloud config set project $PROJECT_ID

# Navigate to Terraform directory
cd "$(dirname "$0")/../terraform/kubernetes"

# Initialize Terraform
echo "Initializing Terraform..."
terraform init

# Plan the changes
echo "Planning Terraform changes..."
terraform plan -var="project_id=$PROJECT_ID"

# Ask for confirmation
read -p "Do you want to apply these changes? (y/n) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    # Apply the changes
    echo "Applying Terraform changes..."
    terraform apply -auto-approve -var="project_id=$PROJECT_ID"

    # Configure kubectl
    echo "Configuring kubectl..."
    gcloud container clusters get-credentials pg-geoshard-cluster \
        --region us-central1 \
        --project $PROJECT_ID

    echo "Setup completed successfully!"
else
    echo "Operation cancelled"
    exit 0
fi