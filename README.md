# Cube Orchestrator

🛠️ Learning to build an orchestrator in Go by following the book "Build an Orchestrator in Go (From Scratch)" from Manning Publications. Exploring concepts like process management, containers, and scheduling from the ground up.

## Project Structure

```text
cube-orchestrator/
├── .copilot/           # GitHub Copilot configuration
│   └── settings.json   # Copilot settings for the project
├── .github/            # GitHub configuration
│   └── copilot-instructions.md # Copilot context and guidelines
├── docs/               # Documentation and images
│   └── images/
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

To manage third-party dependencies, use the Go module system. Run the following command in the project root:

```bash
go mod tidy

# Core dependencies for the orchestrator
$ go get github.com/golang-collections/collections/queue
$ go get github.com/google/uuid
$ go get github.com/docker/go-connections/nat

# Additional useful dependencies for container orchestration
$ go get github.com/docker/docker/api/types
$ go get github.com/docker/docker/client
$ go get github.com/gorilla/mux
$ go get github.com/shirou/gopsutil/v3/cpu
$ go get github.com/shirou/gopsutil/v3/mem
$ go get github.com/sirupsen/logrus
```

## Getting Started

### Running the Application

To run the cube orchestrator demo:

```bash
cd src
go run main.go
```

### Docker Setup

For Docker commands and container management instructions, see [Docker Commands](docs/docker-commands.md).
