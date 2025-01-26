# Setup:
infra/
├── terraform/
│   ├── kubernetes/            # Kubernetes cluster setup
│   └── networking/            # Load balancer and VPC setup
└── scripts/
    ├── apply.sh               # Apply Terraform configs
    └── destroy.sh             # Destroy resources
