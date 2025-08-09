# pkg/ Directory Structure - Future Implementation Guide

This document outlines what will be implemented in the `pkg/` directory when we add API functionality to the cube-orchestrator project in future development phases.

## 📁 Overview

The `pkg/` directory will contain **public API packages** that external projects can import and use. This follows Go best practices where `pkg/` contains library code that's safe for external consumption, while `internal/` contains private implementation details.

## 🚀 Future pkg/ Structure

```text
src/orchestrator/pkg/
├── client/                     # Client libraries for external use
│   ├── worker/                 # Worker API client
│   ├── manager/                # Manager API client
│   └── orchestrator/           # Unified orchestrator client
├── types/                      # Public API types and models
├── errors/                     # Public error types and codes
└── version/                    # Version information and compatibility
```

## 📦 Detailed Directory Contents

### `pkg/client/` - API Client Libraries

#### `pkg/client/worker/`

**Purpose**: Allow external projects to interact with worker nodes programmatically

**Files that will be created**:

- `client.go` - Main worker API client implementation
- `options.go` - Client configuration options and builders
- `types.go` - Worker-specific request/response types

**Example Usage** (Future):

```go
import "cubeorchestrator/pkg/client/worker"

client := worker.NewClient("http://worker:8080")
status, err := client.GetStatus(ctx)
result, err := client.SubmitTask(ctx, task)
```

#### `pkg/client/manager/`

**Purpose**: Allow external projects to interact with the orchestrator manager

**Files that will be created**:

- `client.go` - Main manager API client implementation
- `options.go` - Client configuration options
- `cluster.go` - Cluster management operations

**Example Usage** (Future):

```go
import "cubeorchestrator/pkg/client/manager"

client := manager.NewClient("http://manager:8081")
workers, err := client.ListWorkers(ctx)
err := client.ScheduleTask(ctx, taskDef)
```

#### `pkg/client/orchestrator/`

**Purpose**: Unified client that can interact with the entire orchestrator system

**Files that will be created**:

- `client.go` - Unified orchestrator client
- `config.go` - Configuration management
- `discovery.go` - Service discovery helpers

**Example Usage** (Future):

```go
import "cubeorchestrator/pkg/client/orchestrator"

client := orchestrator.NewClient(config)
cluster, err := client.GetClusterStatus(ctx)
```

### `pkg/types/` - Public API Types

**Purpose**: Shared data structures that external clients need to interact with the API

**Files that will be created**:

- `worker.go` - Worker-related types (WorkerStatus, WorkerMetrics, etc.)
- `manager.go` - Manager-related types (ClusterStatus, SchedulingPolicy, etc.)
- `task.go` - Task-related types (TaskDefinition, TaskResult, etc.)
- `node.go` - Node-related types (NodeInfo, ResourceSpec, etc.)
- `common.go` - Common types and enums used across the API

**Key Types to Include**:

```go
// Example types that will be implemented
type WorkerStatus struct {
    ID       string        `json:"id"`
    Name     string        `json:"name"`
    Status   WorkerState   `json:"status"`
    Metrics  WorkerMetrics `json:"metrics"`
    Tasks    []TaskSummary `json:"tasks"`
    LastSeen time.Time     `json:"last_seen"`
}

type TaskDefinition struct {
    ID      uuid.UUID         `json:"id"`
    Name    string            `json:"name"`
    Image   string            `json:"image"`
    Command []string          `json:"command,omitempty"`
    Env     map[string]string `json:"env,omitempty"`
    CPU     float64           `json:"cpu"`
    Memory  int64             `json:"memory"`
    Disk    int64             `json:"disk"`
}
```

### `pkg/errors/` - Public Error Types

**Purpose**: Standardized error types that external clients can handle appropriately

**Files that will be created**:

- `errors.go` - Core error types and constructors
- `codes.go` - Error codes and status mappings
- `http.go` - HTTP-specific error handling

**Key Error Types**:

```go
// Example error types that will be implemented
type APIError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Details string `json:"details,omitempty"`
}

type ValidationError struct {
    Field   string `json:"field"`
    Value   string `json:"value"`
    Message string `json:"message"`
}
```

### `pkg/version/` - Version Information

**Purpose**: Version compatibility and API versioning information

**Files that will be created**:

- `version.go` - Version constants and compatibility checking
- `compatibility.go` - API version compatibility matrix

## 🎯 When to Implement

### Phase 1: Basic Structure

- Create basic `pkg/types/` with core data structures
- Implement simple error types in `pkg/errors/`

### Phase 2: Worker Client

- Implement `pkg/client/worker/` for worker API interaction
- Add worker-specific types to `pkg/types/worker.go`

### Phase 3: Manager Client

- Implement `pkg/client/manager/` for manager API interaction
- Add cluster management types

### Phase 4: Unified Client

- Create `pkg/client/orchestrator/` as a unified interface
- Add version management and compatibility checking

## 🔗 Integration with Current Code

### Current Internal Structure (Keep As-Is)

```text
src/orchestrator/internal/
├── manager/manager.go          # Manager implementation
├── worker/worker.go            # Worker implementation  
├── task/task.go                # Task definitions
├── node/node.go                # Node abstraction
└── scheduler/scheduler.go      # Scheduler component
```

### Future API Layer (To Be Added)

```text
src/orchestrator/internal/
├── api/                        # HTTP API handlers (private)
│   ├── worker/handlers.go      # Worker endpoints
│   └── manager/handlers.go     # Manager endpoints
├── service/                    # Business logic layer
│   ├── worker/service.go       # Worker business logic
│   └── manager/service.go      # Manager business logic
└── [existing internal packages]
```

## 📋 Implementation Notes

### Import Path Strategy

When implemented, external projects will import like:

```go
import (
    "cubeorchestrator/pkg/client/worker"
    "cubeorchestrator/pkg/types"
    "cubeorchestrator/pkg/errors"
)
```

### Backward Compatibility

- All public APIs in `pkg/` should maintain backward compatibility
- Use semantic versioning for breaking changes
- Deprecate rather than remove functionality when possible

### Testing Strategy

- Each `pkg/` package should have comprehensive tests
- Include integration tests that verify client-server communication
- Mock external dependencies for unit testing

## 🎓 Learning Progression

This structure aligns with your current learning path:

1. **Current**: Focus on core orchestrator concepts (tasks, workers, managers)
2. **Next**: Add HTTP APIs for communication between components
3. **Then**: Create public client libraries for external integration
4. **Finally**: Build unified orchestrator SDK for enterprise use

## 📚 References

- [Go Project Layout](04_go-project-layout.md) - Current project structure documentation
- [API Architecture](06_api-architecture.md) - Detailed API design strategy
- [Troubleshooting](11_troubleshooting.md) - Common issues and solutions

---

**Note**: This is a planning document. The `pkg/` directory will be implemented when you're ready to add API functionality to the cube-orchestrator project. For now, focus on completing the core orchestrator implementation in the `internal/` packages!
