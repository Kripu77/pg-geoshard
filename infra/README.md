# Setup:
infra/
├── terraform/
│   ├── kubernetes/            # Kubernetes cluster setup
│   ├── databases/             # Postgres shards provisioning
│   └── networking/            # Load balancer and VPC setup
└── scripts/
    ├── apply.sh               # Apply Terraform configs
    └── destroy.sh             # Destroy resources
