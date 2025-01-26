# Networking Infrastructure

This directory contains Terraform configurations for setting up the networking infrastructure, including VPC, subnets, and load balancers for the GKE cluster.

## Components

- VPC with custom subnets
- Cloud NAT for egress traffic
- External Load Balancer with NGINX ingress
- Firewall rules for internal and external traffic
- Secondary IP ranges for pods and services

## Prerequisites

- Google Cloud SDK installed and configured
- Terraform installed (version >= 1.0.0)
- kubectl configured with GKE cluster access
- Appropriate GCP permissions

## Configuration

Before applying, ensure you have the following variables ready:

- `project_id`: Your GCP project ID
- `region`: GCP region (defaults to us-central1)

## Usage

1. Initialize Terraform:
```bash
terraform init
```

