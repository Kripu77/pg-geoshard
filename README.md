# PG-GEOSHARD

## Planned Repo Structure:

```
├── README.md                  # Overview of the project
├── docs/                      # Documentation for the project
├── infra/                     # Infrastructure as Code (IaC) setup
├── services/                  # Backend microservices and modules
│   ├── api/                   # API gateway and client routing logic
│   ├── shard-manager/         # Shard management and routing service
│   ├── replication-engine/    # Handles async replication across regions
├── tests/                     # Integration and performance tests
│   ├── e2e/                   # End-to-end tests
│   ├── load-tests/            # Load testing scripts
│   └── unit-tests/            # Unit tests for individual components
├── shared/                    # Shared libraries and utilities
│   ├── common/                # Common utilities (logging, error handling, etc.)
│   ├── db/                    # Database abstraction and connection pooling
│   └── models/                # Shared data models
├── helm/                      # Helm charts for Kubernetes deployment
│   ├── api/                   # Chart for API service
│   ├── shard-manager/         # Chart for shard manager
│   ├── replication-engine/    # Chart for replication engine
│   └── global/                # Parent chart for all services
├── .github/                   # GitHub workflows and templates
│   ├── workflows/             # CI/CD pipelines
│   └── ISSUE_TEMPLATE/        # GitHub issue templates
├── scripts/                   # Helper scripts
│   ├── deploy.sh              # Script for deployment
│   └── test.sh                # Script for running tests
├── kustomize/                 # Kustomize configs for Kubernetes deployment
├── Makefile                   # Task automation commands
```
