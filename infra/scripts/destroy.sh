#!/bin/bash

set -e

# Check if project ID is provided
if [ -z "$1" ]; then
    echo "Error: Please provide your GCP project ID as an argument"
    echo "Usage: ./destroy.sh PROJECT_ID"
    exit 1
fi

PROJECT_ID=$1

# Navigate to Terraform directory
cd "$(dirname "$0")/../terraform/kubernetes"

# Show what will be destroyed
echo "Planning destruction..."
terraform plan -destroy -var="project_id=$PROJECT_ID"

# Ask for confirmation
echo "WARNING: This will destroy all resources in the GKE cluster!"
read -p "Are you sure you want to proceed with destruction? (y/n) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    # Double-check
    read -p "This action is irreversible. Type 'destroy' to confirm: " confirmation
    if [ "$confirmation" = "destroy" ]; then
        echo "Destroying infrastructure..."
        terraform destroy -auto-approve -var="project_id=$PROJECT_ID"
        echo "Infrastructure destroyed successfully!"
    else
        echo "Destruction cancelled"
        exit 0
    fi
else
    echo "Operation cancelled"
    exit 0
fi