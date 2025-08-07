# Go Project Layout Best Practices

This document outlines the standard Go project layout and explains the structure choices for the cube-orchestrator project.

## Standard Go Project Layout

The Go community has established conventions for project structure, popularized by the [golang-standards/project-layout](https://github.com/golang-standards/project-layout) repository.

### Recommended Structure for Our Project

```text
/workspaces/cube-orchestrator/
├── src/
│   ├── orchestrator/                 # Backend Go application
│   │   ├── cmd/                      # Main applications for this project
│   │   │   └── main.go              # Entry point (NOT in subdirectory)
│   │   ├── internal/                 # Private application and library code
│   │   │   ├── manager/             # Manager component
│   │   │   ├── worker/              # Worker component  
│   │   │   ├── node/                # Node abstraction
│   │   │   ├── scheduler/           # Scheduling logic
│   │   │   └── task/                # Task definitions
│   │   ├── pkg/                     # Library code ok to use by external apps
│   │   ├── api/                     # API definitions (future)
│   │   ├── web/                     # Web application assets (future)
│   │   ├── go.mod                   # Go module definition
│   │   └── go.sum                   # Dependency checksums
│   └── frontend/                     # React.js frontend application
│       ├── public/                   # Static assets
│       ├── src/                      # React source code
│       └── package.json              # Node.js dependencies
├── scripts/                          # Build and utility scripts
├── docs/                            # Documentation
├── builds/                          # Build outputs
└── .vscode/                         # VS Code configuration
```

## Key Principles

### 1. `/cmd` Directory

- **Purpose**: Contains main applications for the project
- **Best Practice**: Each subdirectory should be named after the executable you want to have
- **Current Issue**: We should have `cmd/main.go` directly, not `cmd/orchestrator/main.go`
- **Correct**: `src/orchestrator/cmd/main.go` (for single binary)
- **Alternative**: `src/orchestrator/cmd/orchestrator/main.go` (if multiple binaries planned)

### 2. `/internal` Directory

- **Purpose**: Private application and library code
- **Benefit**: Go compiler prevents other projects from importing these packages
- **Use Case**: All core orchestrator logic that shouldn't be reused elsewhere

### 3. `/pkg` Directory

- **Purpose**: Library code that's ok to use by external applications
- **Use Case**: Reusable components, client libraries, common utilities
- **Example**: `pkg/client/` for orchestrator API client

### 4. Module Structure

- **Single Module**: One `go.mod` per project is usually sufficient
- **Location**: Should be at the root of the Go project (`src/orchestrator/go.mod`)
- **Module Name**: Should reflect the import path structure

## Our Current Structure Analysis

### ✅ What's Good

- Separation of backend (`orchestrator`) and frontend
- Use of `internal/` for private code
- Single module approach

### ❌ What Needs Improvement

- `cmd/` structure could be simplified for single binary
- Path depth might be excessive for simple project

## Recommended Changes

### Option 1: Single Binary (Recommended)

```text
src/orchestrator/
├── cmd/
│   └── main.go                 # Single entry point
├── internal/                   # Private packages
├── go.mod
└── go.sum
```

### Option 2: Multiple Binaries (Future-ready)

```text
src/orchestrator/
├── cmd/
│   ├── orchestrator/           # Main orchestrator binary
│   │   └── main.go
│   ├── worker/                 # Standalone worker binary
│   │   └── main.go
│   └── manager/                # Standalone manager binary
│       └── main.go
├── internal/                   # Shared private packages
├── go.mod
└── go.sum
```

## Import Path Considerations

### Current Module Name

```go
module cubeorchestrator
```

### Import Paths

```go
import (
    "cubeorchestrator/internal/manager"
    "cubeorchestrator/internal/task"
    "cubeorchestrator/internal/worker"
)
```

### Alternative Module Name (More Professional)

```go
module github.com/vishipayyallore/cube-orchestrator
```

With imports:

```go
import (
    "github.com/vishipayyallore/cube-orchestrator/internal/manager"
    "github.com/vishipayyallore/cube-orchestrator/internal/task"
)
```

## Benefits of Proper Structure

1. **Industry Standard**: Follows Go community conventions
2. **Tooling Support**: Better IDE and build tool integration
3. **Scalability**: Easy to add new components and binaries
4. **Team Collaboration**: Clear boundaries and responsibilities
5. **Open Source Ready**: Structure supports public repositories

## Build and Development

### Development Commands

```bash
# From orchestrator root
cd src/orchestrator

# Run application
go run cmd/main.go

# Build binary
go build -o ../../builds/orchestrator cmd/main.go

# Run tests
go test ./internal/...
```

### VS Code Integration

- Tasks should reference correct paths
- Debugger should point to `cmd/main.go`
- Module resolution should work seamlessly

## Future Considerations

### Microservices Evolution

If the project grows to multiple services:

```text
src/
├── orchestrator/
│   ├── cmd/main.go
│   └── internal/
├── worker/
│   ├── cmd/main.go
│   └── internal/
└── shared/
    └── pkg/
```

### API Server

For REST API endpoints:

```text
src/orchestrator/
├── cmd/
│   ├── server/main.go          # API server
│   └── cli/main.go             # CLI tool
├── internal/
│   ├── api/                    # HTTP handlers
│   ├── service/                # Business logic
│   └── repository/             # Data access
└── pkg/
    └── client/                 # API client library
```

## Conclusion

The current structure is functional but can be simplified. For a learning project building an orchestrator from scratch, **Option 1 (Single Binary)** is recommended for simplicity while maintaining professional standards.

The key is consistency and following Go community conventions to ensure the project remains maintainable and familiar to other Go developers.
