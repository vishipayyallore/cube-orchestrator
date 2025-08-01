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
├── go.mod              # Go module definition
├── README.md           # Project documentation
├── docs/               # Documentation and images
│   └── images/
└── src/                # Source code directory
```

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
- Implement proper state management for tasks and workers
- Use goroutines and channels for concurrent operations
- Design with scalability and fault tolerance in mind

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
- Implementing scheduling algorithms
- Managing cluster state and health
- Handling failure scenarios gracefully
- Building robust APIs for orchestrator control
