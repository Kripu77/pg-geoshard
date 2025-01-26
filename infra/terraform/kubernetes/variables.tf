variable "project_id" {
  description = "GCP Project ID"
  type        = string
}

variable "region" {
  description = "GCP region"
  type        = string
  default     = "us-central1"
}

variable "cluster_name" {
  description = "Name of the GKE cluster"
  type        = string
  default     = "pg-geoshard-cluster"
}

variable "network_name" {
  description = "Name for the VPC network"
  type        = string
  default     = "pg-geoshard-network"
}

variable "subnet_name" {
  description = "Name for the subnet"
  type        = string
  default     = "pg-geoshard-subnet"
}

variable "subnet_cidr" {
  description = "CIDR range for the subnet"
  type        = string
  default     = "10.0.0.0/16"
}