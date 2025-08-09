# API Architecture Strategy for Cube Orchestrator

This document outlines the strategic approach for introducing APIs in the cube-orchestrator project, including the proper use of `pkg/` and `internal/` directories according to Go best practices.

## 🎯 Overview

As we evolve from the current demo application to a production-ready orchestrator, we need to introduce HTTP APIs for communication between managers, workers, and external clients. This document defines the architecture strategy for implementing these APIs while maintaining clean separation of concerns.

## 📁 Proposed Project Structure with APIs

### Current Structure

```text
src/orchestrator/
├── cmd/main.go                     # Demo application
├── internal/                       # Private implementation
│   ├── runtime/docker_wrapper.go   # Docker runtime abstraction
│   ├── manager/manager.go          # Manager component
│   ├── worker/worker.go            # Worker component
│   ├── task/                       # Task management
│   │   ├── task.go                 # Task definitions (cleaned up)
│   │   └── state_machine.go        # State transition logic
│   ├── node/node.go                # Node abstraction
│   └── scheduler/scheduler.go      # Scheduler component
├── pkg/                            # Future public APIs (planned)
│   ├── client/worker/              # Worker API client (planned)
│   └── types/                      # Public type definitions (planned)
├── go.mod
└── go.sum
```

### Proposed API-Enhanced Structure

```text
src/orchestrator/
├── cmd/                            # Application entry points
│   ├── orchestrator/               # Main orchestrator binary
│   │   └── main.go                 # Combined manager + worker
│   ├── worker/                     # Standalone worker binary (future)
│   │   └── main.go                 # Worker-only mode
│   └── manager/                    # Standalone manager binary (future)
│       └── main.go                 # Manager-only mode
├── internal/                       # Private implementation
│   ├── api/                        # HTTP API handlers (private)
│   │   ├── worker/
│   │   ├── manager/
│   │   ├── middleware/
│   │   └── server/
│   ├── service/                    # Business logic layer
│   │   ├── worker/
│   │   ├── manager/
│   │   └── task/
│   ├── repository/                 # Data access layer
│   ├── worker/                     # Current worker implementation
│   ├── manager/                    # Current manager implementation
│   ├── task/                       # Current task implementation
│   ├── node/                       # Current node implementation
│   └── runtime/                    # Docker runtime abstraction
├── pkg/                            # Public API packages
│   ├── client/                     # Client libraries for external use
│   │   ├── worker/
│   │   ├── manager/
│   │   └── orchestrator/
│   ├── types/                      # Public API types and models
│   ├── errors/                     # Public error types
│   └── version/                    # Version information
├── api/                            # API specifications and documentation
│   ├── openapi/                    # OpenAPI/Swagger specifications
│   └── docs/                       # Generated API documentation
├── web/                            # Static web assets (future)
├── go.mod
└── go.sum
```

## 🔄 API Communication Architecture

### Worker API Endpoints

Based on the worker architecture diagram, the following endpoints will be implemented:

```go
// Worker receives tasks from manager via API
POST   /api/v1/worker/tasks              # Submit task to worker queue
GET    /api/v1/worker/tasks              # List worker tasks
GET    /api/v1/worker/tasks/{id}         # Get specific task status
DELETE /api/v1/worker/tasks/{id}         # Cancel/stop task

// Worker provides metrics to manager via API
GET    /api/v1/worker/metrics            # Get worker host and task metrics
GET    /api/v1/worker/status             # Worker health check
GET    /api/v1/worker/info               # Worker node information

// Worker task lifecycle management
POST   /api/v1/worker/tasks/{id}/start   # Start specific task
POST   /api/v1/worker/tasks/{id}/stop    # Stop specific task
GET    /api/v1/worker/tasks/{id}/logs    # Get task logs from Docker runtime
```

### Manager API Endpoints

```go
// Cluster management
POST   /api/v1/manager/workers           # Register new worker with manager
GET    /api/v1/manager/workers           # List all workers
GET    /api/v1/manager/workers/{id}      # Get specific worker info
DELETE /api/v1/manager/workers/{id}      # Unregister worker

// Task scheduling and management
POST   /api/v1/manager/tasks             # Submit task for scheduling
GET    /api/v1/manager/tasks             # List all tasks in system
GET    /api/v1/manager/tasks/{id}        # Get task details
DELETE /api/v1/manager/tasks/{id}        # Cancel task

// Cluster status and metrics
GET    /api/v1/manager/cluster           # Overall cluster status
GET    /api/v1/manager/metrics           # Aggregated cluster metrics
GET    /api/v1/manager/events            # System events and task history
```

## 📦 Package Usage Strategy

### When to Use `internal/` vs `pkg/`

#### Use `internal/` for

- ✅ **HTTP Handlers**: Route implementations and request processing
- ✅ **Business Logic**: Core orchestrator domain logic
- ✅ **Database Access**: Repository patterns and data persistence
- ✅ **Authentication**: Security middleware and auth logic
- ✅ **Implementation Details**: Docker runtime, metrics collection, scheduling algorithms

#### Use `pkg/` for

- ✅ **Client Libraries**: SDKs for external projects to interact with orchestrator
- ✅ **Public Types**: API request/response structures that external clients need
- ✅ **Interface Definitions**: Contracts that external projects can implement
- ✅ **Utility Functions**: Reusable helpers that external projects might need
- ✅ **Error Types**: Standard error formats for API responses

### Example Package Contents

#### `pkg/types/worker.go` - Public API Types

```go
package types

// Public API types that external clients can import and use
type WorkerStatus struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Status      WorkerState       `json:"status"`
    Node        NodeInfo          `json:"node"`
    Metrics     WorkerMetrics     `json:"metrics"`
    Tasks       []TaskSummary     `json:"tasks"`
    LastSeen    time.Time         `json:"last_seen"`
}

type WorkerMetrics struct {
    CPUUsage      float64           `json:"cpu_usage"`
    MemoryUsage   int64             `json:"memory_usage"`
    MemoryTotal   int64             `json:"memory_total"`
    DiskUsage     int64             `json:"disk_usage"`
    DiskTotal     int64             `json:"disk_total"`
    TaskCount     int               `json:"task_count"`
    TasksRunning  int               `json:"tasks_running"`
    Uptime        time.Duration     `json:"uptime"`
}

type WorkerState string
const (
    WorkerStateIdle       WorkerState = "idle"
    WorkerStateBusy       WorkerState = "busy"
    WorkerStateOffline    WorkerState = "offline"
    WorkerStateMaintenance WorkerState = "maintenance"
)
```

#### `pkg/client/worker/client.go` - Public Client Library

```go
package worker

import (
    "context"
    "cubeorchestrator/pkg/types"
    "cubeorchestrator/pkg/errors"
)

// Client provides a programmatic interface to the Worker API
type Client struct {
    baseURL    string
    httpClient *http.Client
    apiKey     string
}

// NewClient creates a new worker API client
func NewClient(baseURL string, opts ...Option) *Client {
    client := &Client{
        baseURL:    baseURL,
        httpClient: &http.Client{Timeout: 30 * time.Second},
    }
    
    for _, opt := range opts {
        opt(client)
    }
    
    return client
}

// SubmitTask submits a task to the worker for execution
func (c *Client) SubmitTask(ctx context.Context, task *types.TaskRequest) (*types.TaskResponse, error) {
    // Implementation for external clients to submit tasks to workers
    // This would be used by managers or external orchestration tools
}

// GetWorkerStatus retrieves current worker status and metrics
func (c *Client) GetWorkerStatus(ctx context.Context) (*types.WorkerStatus, error) {
    // Implementation for external monitoring tools to get worker status
}

// GetTaskStatus retrieves status of a specific task
func (c *Client) GetTaskStatus(ctx context.Context, taskID string) (*types.TaskStatus, error) {
    // Implementation for tracking task progress
}
```

#### `internal/api/worker/handlers.go` - Private HTTP Handlers

```go
package worker

import (
    "cubeorchestrator/internal/service/worker"
    "cubeorchestrator/pkg/types"
    "cubeorchestrator/pkg/errors"
)

// Handler implements the HTTP handlers for worker API endpoints
type Handler struct {
    workerService *worker.Service
    logger        Logger
}

// SubmitTask handles POST /api/v1/worker/tasks
func (h *Handler) SubmitTask(w http.ResponseWriter, r *http.Request) {
    // Parse request
    var req types.TaskRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        h.writeError(w, errors.NewValidationError("invalid request body"))
        return
    }
    
    // Submit to worker service
    result, err := h.workerService.SubmitTask(r.Context(), &req)
    if err != nil {
        h.writeError(w, err)
        return
    }
    
    // Return response
    h.writeJSON(w, http.StatusCreated, result)
}

// GetMetrics handles GET /api/v1/worker/metrics
func (h *Handler) GetMetrics(w http.ResponseWriter, r *http.Request) {
    // Implementation matches the "Metrics" component from your architecture diagram
    metrics, err := h.workerService.GetMetrics(r.Context())
    if err != nil {
        h.writeError(w, err)
        return
    }
    
    h.writeJSON(w, http.StatusOK, metrics)
}
```

## 🚀 Implementation Phases

### Phase 1: Basic API Structure

1. Create `pkg/types/` with basic API types
2. Create `internal/api/` with simple HTTP handlers
3. Add HTTP server setup in `internal/api/server/`
4. Modify current `cmd/main.go` to start HTTP server alongside demo

### Phase 2: Worker API Implementation

1. Implement worker API endpoints based on architecture diagram
2. Create `pkg/client/worker/` client library
3. Add metrics collection API (CPU, memory, disk, task count)
4. Add task lifecycle management endpoints

### Phase 3: Manager API Implementation

1. Implement manager API endpoints for cluster management
2. Create `pkg/client/manager/` client library
3. Add worker registration and discovery
4. Add task scheduling API

### Phase 4: Full Client SDK

1. Create `pkg/client/orchestrator/` unified client
2. Add comprehensive error handling in `pkg/errors/`
3. Generate OpenAPI specifications
4. Add API documentation and examples

## 🔗 Integration with Current Demo

The current `main.go` can be enhanced to start HTTP servers:

```go
// Enhanced main.go structure
func main() {
    // Current demo code...
    
    // Start API servers
    go startWorkerAPI(w)    // Worker API on :8080
    go startManagerAPI(m)   // Manager API on :8081
    
    // Keep demo running
    select {}
}
```

## 📈 Benefits of This Architecture

### **Clean Separation**

- Business logic stays in `internal/service/`
- HTTP concerns isolated in `internal/api/`
- Public contracts defined in `pkg/`

### **External Integration**

- Other projects can import `pkg/client/` for programmatic access
- Clear API boundaries enable microservices evolution
- Standard Go project layout for enterprise adoption

### **Maintainability**

- APIs can evolve independently
- Clear dependency boundaries
- Testable components with mocked interfaces

### **Scalability**

- Multiple binaries (worker-only, manager-only) from same codebase
- Load balancer-friendly stateless API design
- Horizontal scaling of workers and managers

This architecture transforms your learning project into a production-ready orchestrator while maintaining the educational value and clean code structure! 🚀
