# Project Structure

This project follows the Standard Go Project Layout with support for future web UI integration.

## Directory Structure

```text
/workspaces/cube-orchestrator/
├── src/
│   ├── cmd/
│   │   └── orchestrator/          # Main application entry point
│   │       └── main.go
│   ├── internal/                  # Private application code
│   │   ├── manager/               # Task management logic
│   │   ├── worker/                # Worker node implementation
│   │   ├── node/                  # Node abstraction
│   │   ├── scheduler/             # Task scheduling algorithms
│   │   └── task/                  # Task definitions and state machine
│   ├── pkg/                       # Public library code (when needed)
│   └── web/                       # Future: React.js UI components
│       ├── public/                # Static assets
│       ├── src/                   # React source code
│       └── package.json           # Node.js dependencies
├── scripts/                       # Build and utility scripts
│   ├── build.sh                   # Timestamped build script
│   └── cleanup-builds.sh          # Build artifact cleanup
├── builds/                        # Build outputs
├── docs/                          # Documentation
└── .vscode/                       # VS Code configuration
    ├── tasks.json                 # Build and run tasks
    └── launch.json                # Debug configurations
```

## Why This Structure?

### `/src/cmd/` Pattern

- **Industry Standard**: Follows the widely-adopted Go project layout
- **Multiple Binaries**: Easy to add more commands (worker, manager, scheduler as separate binaries)
- **Clean Separation**: Application entry points are clearly defined
- **Tooling Friendly**: Works well with Go modules and build tools

### `/src/internal/` Pattern

- **Encapsulation**: Code is private to this project
- **No External Dependencies**: Other projects can't import these packages
- **Clean Architecture**: Enforces proper dependency management
- **Refactoring Safety**: Internal restructuring doesn't break external users

### `/src/pkg/` Pattern (Future)

- **Public Libraries**: Code that could be reused by other projects
- **API Clients**: Orchestrator client libraries
- **Shared Utilities**: Common functionality for external consumption

### `/src/web/` Pattern (Future)

- **UI Integration**: React.js application alongside Go backend
- **Full-Stack Project**: Single repository for complete solution
- **Development Workflow**: Shared configuration and scripts
- **Deployment**: Unified build and deployment process

## Future React.js Integration

When adding the React.js UI, the structure will support:

```bash
# Development servers
npm run dev          # React development server (port 3000)
go run ./src/cmd/orchestrator  # Go API server (port 8080)

# Production build
npm run build        # Build React app to src/web/build/
go build ./src/cmd/orchestrator  # Build Go binary with embedded UI
```

## Build System

### Development

```bash
# Quick run
cd src/cmd/orchestrator
go run main.go

# Debug build
./scripts/build.sh

# Run latest build
./builds/cube-orchestrator_latest
```

### VS Code Integration

- **F5**: Start debugging
- **Ctrl+Shift+P** → "Tasks: Run Test Task": Quick run
- **Ctrl+Shift+P** → "Tasks: Run Build Task": Timestamped build

## Migration Benefits

1. **Professional Structure**: Industry-standard Go project layout
2. **Future-Proof**: Ready for React.js UI integration
3. **Tooling Support**: Better IDE and build tool integration
4. **Scalability**: Easy to add new commands and services
5. **Team Development**: Clear separation of concerns
6. **Deployment**: Simplified containerization and deployment

## Development Workflow

The new structure supports modern development practices:

- Hot reloading for both Go (with tools) and React
- Separate development and production builds
- Integrated debugging across backend and frontend
- Professional CI/CD pipeline preparation
