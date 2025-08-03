# Cube Orchestrator

🛠️ Learning to build an orchestrator in Go by following the book `Build an Orchestrator in Go (From Scratch)` from Manning Publications. Exploring concepts like process management, containers, and scheduling from the ground up.

## Project Structure

```text
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
│   ├── main.go         # Main application entry point
│   ├── manager/        # Orchestrator manager component
│   │   └── manager.go  # Manager implementation
│   ├── node/           # Node management component
│   │   └── node.go     # Node implementation
│   ├── scheduler/      # Task scheduling component
│   │   └── scheduler.go # Scheduler implementation
│   ├── task/           # Task definition and management
│   │   └── task.go     # Task implementation
│   └── worker/         # Worker node component
│       └── worker.go   # Worker implementation
├── go.mod              # Go module definition
├── LICENSE             # Project license
└── README.md           # Project documentation
```

## Getting third-party dependencies

To manage third-party dependencies, use the Go module system. Run the following commands in the project root:

```bash
# Clean module cache and tidy dependencies
go clean -modcache
go mod tidy

# Core dependencies for the orchestrator (currently installed)
go get github.com/golang-collections/collections/queue
go get github.com/google/uuid
go get github.com/docker/go-connections/nat
```

### Additional Dependencies (for future chapters)

These dependencies will be needed as you progress through the book chapters:

```bash
# HTTP routing and API development
go get github.com/gorilla/mux

# System monitoring and resource management
go get github.com/shirou/gopsutil/v3/cpu
go get github.com/shirou/gopsutil/v3/mem

# Structured logging
go get github.com/sirupsen/logrus
```

### Docker Client Dependencies (Troubleshooting)

**Note**: The Docker client imports may cause module path conflicts. If you encounter errors like:

```text
module declares its path as: github.com/moby/moby/client but was required as: github.com/docker/docker/client
```

**Solutions**:

1. **Temporary approach**: Remove Docker client imports until needed in later chapters
2. **Alternative approach**: Use the Moby client directly:

   ```bash
   go get github.com/moby/moby/client
   ```

3. **Wait for book updates**: The book may provide updated import instructions

## Getting Started

### Running the Application

To run the cube orchestrator demo:

```bash
cd src
go run main.go
```

![Cube Orchestrator Demo Output](docs/images/After_Ch_2.PNG)

### Docker Container Integration (Chapter 3)

The application now includes Docker container management capabilities:

![Chapter 3 - Docker Integration](docs/images/After_Ch_3.PNG)

### Docker Setup

For Docker commands and container management instructions, see [Docker Commands](docs/docker-commands.md).

### Troubleshooting

If you encounter any issues with dependencies, imports, or compilation, see the [Troubleshooting Guide](docs/troubleshooting.md).

## Development
