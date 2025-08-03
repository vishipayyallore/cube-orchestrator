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
│   └── copilot-instructions.md # Copilot context and guidelines
├── docs/               # Documentation and images
│   ├── images/
│   ├── docker-commands.md # Docker commands reference  
│   ├── postgresql-primer.md # PostgreSQL guide
│   └── troubleshooting.md # Common issues and solutions
├── src/                # Source code directory
│   ├── main.go         # Main application with orchestrator demo
│   ├── manager/        # Orchestrator manager component
│   │   └── manager.go  # Manager implementation for task coordination
│   ├── node/           # Node management component
│   │   └── node.go     # Node implementation for cluster resources
│   ├── scheduler/      # Task scheduling component
│   │   └── scheduler.go # Scheduler implementation for task distribution
│   ├── task/           # Task definition and management
│   │   └── task.go     # Task and TaskEvent implementations
│   └── worker/         # Worker node component
│       └── worker.go   # Worker implementation for task execution
├── go.mod              # Go module definition with dependencies
├── LICENSE             # Project license
└── README.md           # Project documentation
```

## Current Implementation Status

The project currently includes:
- **Task Management**: Task and TaskEvent structures with states (Pending, Running, Completed, Failed)
- **Worker Implementation**: Worker nodes with task queues, databases, and lifecycle management
- **Manager Coordination**: Manager for task distribution, worker selection, and system coordination
- **Node Resources**: Node definitions with resource specifications (CPU, memory, disk)
- **Docker Integration**: Complete Docker client functionality for container management
- **Demo Application**: Functional main.go demonstrating all components including Docker containers

## Key Dependencies

- `github.com/golang-collections/collections/queue` - Queue data structure for task management
- `github.com/google/uuid` - UUID generation for unique task and event IDs
- `github.com/docker/go-connections/nat` - Docker networking utilities
- `github.com/docker/docker/client` - Docker client for container management
- `github.com/docker/docker/api/types` - Docker API types and structures
- Additional Docker, monitoring, and HTTP routing libraries for full orchestrator functionality

## Coding Guidelines

### Go Standards
- Follow standard Go conventions and idioms
- Use `gofmt` for code formatting
- Write clear, descriptive variable and function names
- Include appropriate error handling
- Add meaningful comments for complex logic

### Architecture Principles
- Design for modularity and testability
- Implement clear separation of concerns
- Use interfaces to define contracts between components
- Keep functions focused and single-purpose

### Orchestrator-Specific Patterns
- Follow event-driven architecture where appropriate
- Implement proper state management for tasks and workers (Pending, Running, Completed, Failed states)
- Use goroutines and channels for concurrent operations
- Design with scalability and fault tolerance in mind
- Utilize queue-based task distribution patterns
- Implement resource-aware scheduling algorithms
- Use UUID-based identification for tasks and events
- Maintain separation between manager, worker, and node responsibilities

### Code Organization
- Place related functionality in appropriate packages
- Use descriptive package names that reflect their purpose
- Keep main business logic separate from infrastructure code
- Write comprehensive tests for critical components

### Documentation
- Include package-level documentation
- Document exported functions and types
- Provide examples for complex APIs
- Keep README and docs up to date

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
