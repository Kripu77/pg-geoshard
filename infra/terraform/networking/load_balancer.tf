resource "google_compute_address" "ingress_ip" {
  name = "${var.network_name}-ingress-ip"
}

resource "kubernetes_namespace" "ingress_nginx" {
  metadata {
    name = "ingress-nginx"
  }
}

resource "kubernetes_config_map" "nginx_config" {
  metadata {
    name      = "nginx-configuration"
    namespace = kubernetes_namespace.ingress_nginx.metadata[0].name
  }

  data = {
    "use-proxy-protocol" = "true"
    "proxy-real-ip-cidr" = "0.0.0.0/0"
  }
}

resource "kubernetes_service" "ingress_nginx" {
  metadata {
    name      = "ingress-nginx"
    namespace = kubernetes_namespace.ingress_nginx.metadata[0].name
  }

  spec {
    type             = "LoadBalancer"
    load_balancer_ip = google_compute_address.ingress_ip.address

    port {
      name        = "http"
      port        = 80
      target_port = 80
    }

    port {
      name        = "https"
      port        = 443
      target_port = 443
    }

    selector = {
      "app.kubernetes.io/name" = "ingress-nginx"
    }
  }
}