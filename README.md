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
