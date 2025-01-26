output "network_name" {
  description = "The name of the VPC network"
  value       = google_compute_network.vpc.name
}

output "subnet_name" {
  description = "The name of the subnet"
  value       = google_compute_subnetwork.subnet.name
}

output "ingress_ip" {
  description = "The external IP address for the load balancer"
  value       = google_compute_address.ingress_ip.address
}

output "nat_ip" {
  description = "The NAT IP addresses"
  value       = google_compute_router_nat.nat.nat_ips
}