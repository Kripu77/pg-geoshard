variable "project_id" {
  description = "GCP Project ID"
  type        = string
}

variable "region" {
  description = "GCP region"
  type        = string
  default     = "us-central1"
}

variable "network_name" {
  description = "Name of the VPC network"
  type        = string
  default     = "pg-geoshard-network"
}

variable "subnet_name" {
  description = "Name of the subnet"
  type        = string
  default     = "pg-geoshard-subnet"
}

variable "ip_range_pods" {
  description = "CIDR range for pods"
  type        = string
  default     = "10.20.0.0/16"
}

variable "ip_range_services" {
  description = "CIDR range for services"
  type        = string
  default     = "10.30.0.0/16"
}