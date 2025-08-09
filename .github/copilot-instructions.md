# GitHub Copilot Instructions

This document provides context and coding guidelines for GitHub Copilot when working on the cube-orchestrator project.

## Project Overview

This is a learning project focused on building an orchestrator in Go from scratch, following the book "Build an Orchestrator in Go (From Scratch)" from Manning Publications. The project explores fundamental concepts like:

- Process management
- Container orchestration
- Task scheduling
- Resource allocation
- Distributed systems

## Project Structure

```
cube-orchestrator/
├── .copilot/           # GitHub Copilot configuration
│   └── settings.json   # Copilot settings for the project
├── .github/            # GitHub configuration
│   ├── copilot-instructions.md # Copilot context and guidelines
│   └── workflows/      # GitHub Actions workflows (CI)
│       └── docs-quality.yml    # Docs lint + link check
├── .vscode/            # VS Code workspace configuration
│   ├── launch.json     # Debug configurations
│   ├── settings.json   # Workspace settings
│   └── tasks.json      # Build and run tasks
├── .markdownlint.json  # Markdown lint config (root)
├── lychee.toml         # Link checker config (root)
├── builds/             # Build artifacts directory
│   ├── cube-orchestrator-debug    # Debug build
│   ├── cube-orchestrator_latest   # Latest timestamped build
│   └── cube-orchestrator_YYYYMMDD_HHMMSS # Timestamped builds
├── chx/                # Chapter exercises and examples
├── docs/               # Comprehensive documentation suite
│   ├── images/         # Documentation images and diagrams
│   ├── 00_README.md                  # Docs index and reading order
│   ├── 01_project-overview.md        # High-level project overview
│   ├── 02_project-structure.md       # Detailed structure documentation
│   ├── 03_configuration-verification.md # Environment setup verification
│   ├── 04_go-project-layout.md       # Go project structure guidelines
│   ├── 05_build-system.md            # Build system documentation
│   ├── 06_api-architecture.md        # API design patterns and structure
│   ├── 07_pkg-directory-plan.md      # Future API package planning
│   ├── 08_docker-images-reference.md # Docker images used
│   ├── 09_docker-commands.md         # Docker commands reference
│   ├── 10_postgresql-primer.md       # PostgreSQL guide
│   └── 11_troubleshooting.md         # Common issues and solutions
├── scripts/            # Build and utility scripts
│   ├── build.sh        # Professional build script with timestamping
│   └── cleanup-builds.sh # Build artifact cleanup utility
├── src/                # Source code directory
│   ├── frontend/       # Future frontend components
│   └── orchestrator/   # Main orchestrator application
│       ├── cmd/        # Application entry points
│       │   └── main.go # Main application with orchestrator demo
│       ├── internal/   # Private application packages
│       │   ├── runtime/    # Docker runtime abstraction (DockerWrapper)
│       │   ├── manager/    # Orchestrator manager component
│       │   ├── node/       # Node management component
│       │   ├── scheduler/  # Task scheduling component
│       │   ├── task/       # Task definition and management
│       │   └── worker/     # Worker node component
│       ├── pkg/        # Public API packages (planned)
│       ├── go.mod      # Go module definition with dependencies
│       └── go.sum      # Go module checksums
├── .gitignore          # Git ignore patterns
├── LICENSE             # Project license
└── README.md           # Project documentation
```

## Current Implementation Status

The project currently includes:

- **Task Management**: Task and TaskEvent structures with states (Pending, Scheduled, Running, Completed, Failed)
- **Docker Integration**: Dedicated Docker package with client abstraction for container management
- **Worker Implementation**: Worker nodes with task queues, databases, and lifecycle management
- **Manager Coordination**: Manager for task distribution, worker selection, and system coordination
- **Node Resources**: Node definitions with resource specifications (CPU, memory, disk)
- **Scheduler Logic**: Task scheduling algorithms and resource allocation
- **Demo Application**: Functional main.go demonstrating all components including Docker containers
- **Build System**: Professional build system with timestamped artifacts and cleanup automation
- **Documentation Suite**: Comprehensive documentation covering architecture, APIs, and troubleshooting

## Architecture Evolution

The project follows Go standard project layout:

- `cmd/`: Application entry points and main packages
- `internal/`: Private packages not intended for external use
- `pkg/`: Public API packages (planned for future external consumption)
- `docs/`: Comprehensive documentation with architecture guides
- `builds/`: Professional build artifacts with timestamping
- `scripts/`: Build automation and utility scripts

## Key Dependencies

- `github.com/golang-collections/collections/queue` - Queue data structure for task management
- `github.com/google/uuid` - UUID generation for unique task and event IDs
- `github.com/docker/go-connections/nat` - Docker networking utilities
- `github.com/docker/docker/client` - Docker client for container management (v28.3.3+incompatible)
- `github.com/docker/docker/api/types/container` - Docker container API types
- `github.com/docker/docker/api/types/image` - Docker image API types
- `github.com/docker/docker/pkg/stdcopy` - Stream container logs (v28 path; no moby dependency)
- Additional Docker, monitoring, and HTTP routing libraries for full orchestrator functionality

**Security Note**: Uses Docker v28.3.3 to address GO-2023-1699, GO-2023-1700, and GO-2023-1701 vulnerabilities.

Note on Docker SDK v28: some types were renamed (e.g., `image.PullOptions`, `container.StartOptions`). Use `github.com/docker/docker/pkg/stdcopy` for log copying.

## Coding Guidelines

### Go Standards

- Follow standard Go conventions and idioms
- Use `gofmt` for code formatting
- Write clear, descriptive variable and function names
- Include appropriate error handling
- Add meaningful comments for complex logic
- **File Headers**: Always include a file path comment at the top of Go files using the format:

  ```go
  // File: src/orchestrator/internal/package/filename.go

  package packagename
  ```

### Architecture Principles

- Design for modularity and testability
- Implement clear separation of concerns
- Use interfaces to define contracts between components
- Keep functions focused and single-purpose

### Orchestrator-Specific Patterns

- Follow event-driven architecture where appropriate
- Implement proper state management for tasks and workers (Pending, Scheduled, Running, Completed, Failed states)
- Use goroutines and channels for concurrent operations
- Design with scalability and fault tolerance in mind
- Utilize queue-based task distribution patterns
- Implement resource-aware scheduling algorithms
- Use UUID-based identification for tasks and events
- Maintain separation between manager, worker, and node responsibilities
- Abstract Docker operations through dedicated client interface
- Follow internal/ vs pkg/ package organization for API development

### Code Organization

- Place private functionality in `internal/` packages
- Reserve `pkg/` for future public APIs
- Use descriptive package names that reflect their purpose
- Keep main business logic separate from infrastructure code
- Write comprehensive tests for critical components
- Follow established patterns for Docker abstraction and container lifecycle management

### Documentation

- Include package-level documentation
- Document exported functions and types
- Provide examples for complex APIs
- Keep README and docs up to date
- Follow numbered docs convention under `docs/`
- CI enforces docs quality: markdownlint and link checks via `.github/workflows/docs-quality.yml`
- Project-wide configs: `.markdownlint.json` and `lychee.toml` at repo root

## Development Environment

- Target Go version: Latest stable
- Development container: Ubuntu 24.04.2 LTS
- Available tools: docker, kubectl, git, gh, and standard Unix utilities

## Testing Strategy

- Write unit tests for individual components
- Include integration tests for orchestrator workflows
- Test error conditions and edge cases
- Aim for good test coverage of critical paths

## Learning Focus Areas

When suggesting code or improvements, consider these key learning objectives:

- Understanding container runtime interfaces
- Implementing scheduling algorithms based on resource availability
- Managing cluster state and health monitoring
- Handling failure scenarios gracefully (task failures, worker disconnections)
- Building robust APIs for orchestrator control
- Task lifecycle management (creation, scheduling, execution, completion)
- Worker node management and resource allocation
- Event-driven communication patterns between components
