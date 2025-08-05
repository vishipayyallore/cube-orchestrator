# Project Structure

This project follows the Standard Go Project Layout with support for future web UI integration.

## Directory Structure

```text
/workspaces/cube-orchestrator/
├── .copilot/                     # GitHub Copilot configuration
│   └── settings.json             # Copilot project-specific settings
├── .git/                         # Git repository metadata
├── .github/                      # GitHub-specific configuration
│   └── copilot-instructions.md   # Copilot context and guidelines
├── .gitignore                    # Git ignore patterns
├── .vscode/                      # VS Code workspace configuration
│   ├── tasks.json                # Build and run tasks
│   └── launch.json               # Debug configurations
├── LICENSE                       # Project license (MIT)
├── README.md                     # Project overview and setup
├── builds/                       # Build outputs and executables
│   ├── cube-orchestrator_*       # Timestamped build artifacts
│   ├── cube-orchestrator_latest  # Symlink to latest build
│   └── cube-orchestrator-debug   # Debug build
├── chx/                          # Chapter exercises and references
├── docs/                         # Comprehensive documentation
│   ├── images/                   # Documentation images and diagrams
│   ├── build-system.md           # Build system documentation
│   ├── configuration-verification.md # Config verification report
│   ├── docker-commands.md        # Docker reference commands
│   ├── go-project-layout.md      # Go best practices guide
│   ├── postgresql-primer.md      # PostgreSQL setup guide
│   ├── project-structure.md      # This file - project organization
│   └── troubleshooting.md        # Common issues and solutions
├── scripts/                      # Build and utility scripts
│   ├── build.sh                  # Timestamped build script
│   └── cleanup-builds.sh         # Build artifact management
└── src/                          # Source code
    ├── orchestrator/             # Go backend application
    │   ├── cmd/                  # Application entry points
    │   │   └── main.go           # Main orchestrator application
    │   ├── internal/             # Private application code
    │   │   ├── manager/          # Task management and coordination
    │   │   ├── worker/           # Worker node implementation
    │   │   ├── node/             # Node abstraction and resources
    │   │   ├── scheduler/        # Task scheduling algorithms
    │   │   └── task/             # Task definitions and state machine
    │   ├── pkg/                  # Public library code (future)
    │   ├── go.mod                # Go module definition
    │   └── go.sum                # Dependency checksums and verification
    └── frontend/                 # React.js frontend application (future)
        ├── public/               # Static assets and index.html
        ├── src/                  # React source components
        └── package.json          # Node.js dependencies and scripts
```

## Why This Structure?

### Configuration Directories

#### `.copilot/` - GitHub Copilot Configuration

- **Purpose**: Project-specific AI assistant settings
- **Contains**: Custom instructions and context for GitHub Copilot
- **Benefit**: Provides tailored AI assistance for this orchestrator project

#### `.github/` - GitHub Integration

- **Purpose**: GitHub-specific project configuration
- **Contains**: Copilot instructions, workflows, templates
- **Future**: CI/CD workflows, issue templates, pull request templates

#### `.vscode/` - VS Code Workspace

- **Purpose**: Development environment configuration
- **Contains**: Tasks for building/running, debug configurations
- **Benefit**: Consistent development experience across team members

### Core Directories

### `/src/orchestrator/cmd/` Pattern

- **Industry Standard**: Follows the widely-adopted Go project layout
- **Multiple Binaries**: Easy to add more commands (worker, manager, scheduler as separate binaries)
- **Clean Separation**: Application entry points are clearly defined
- **Tooling Friendly**: Works well with Go modules and build tools

### `/src/orchestrator/internal/` Pattern

- **Encapsulation**: Code is private to this project
- **No External Dependencies**: Other projects can't import these packages
- **Clean Architecture**: Enforces proper dependency management
- **Refactoring Safety**: Internal restructuring doesn't break external users

### `/src/orchestrator/pkg/` Pattern (Future)

- **Public Libraries**: Code that could be reused by other projects
- **API Clients**: Orchestrator client libraries
- **Shared Utilities**: Common functionality for external consumption

### `/src/frontend/` Pattern

- **UI Integration**: React.js application alongside Go backend
- **Full-Stack Project**: Single repository for complete solution
- **Development Workflow**: Shared configuration and scripts
- **Deployment**: Unified build and deployment process

### Support Directories

#### `builds/` - Build Artifacts

- **Purpose**: Contains compiled executables and build outputs
- **Management**: Automated cleanup with configurable retention
- **Organization**: Timestamped builds with symlink to latest

#### `docs/` - Documentation Hub

- **Purpose**: Comprehensive project documentation
- **Structure**: Organized by topic (setup, troubleshooting, reference)
- **Maintenance**: Version-controlled alongside code

#### `scripts/` - Automation

- **Purpose**: Build scripts, utilities, and automation tools
- **Benefits**: Consistent build process, artifact management
- **Portability**: Works across different development environments

## Future React.js Integration

When adding the React.js UI, the structure will support:

```bash
# Development servers
npm run dev          # React development server (port 3000)
cd src/orchestrator/cmd && go run main.go  # Go API server (port 8080)

# Production build
npm run build        # Build React app to src/frontend/build/
go build -o builds/orchestrator src/orchestrator/cmd/main.go  # Build Go binary
```

## Build System

### Development

```bash
# Quick run
cd src/orchestrator/cmd
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

## Project Benefits

### Professional Organization

- **Industry Standards**: Follows established Go and modern development practices
- **Tool Integration**: Seamless VS Code, GitHub, and Copilot integration
- **Documentation-First**: Comprehensive docs for maintainability
- **Automation**: Build scripts and cleanup tools for efficient development

### Development Experience

- **IDE Ready**: Pre-configured VS Code tasks and debugging
- **AI Assisted**: GitHub Copilot with project-specific context
- **Quick Setup**: One-command build and run process
- **Clean Workspace**: Automated build artifact management

### Scalability & Maintenance

- **Modular Architecture**: Clear separation between components
- **Future-Proof**: Ready for frontend integration and microservices
- **Team Collaboration**: Consistent structure and workflows
- **Version Control**: Proper `.gitignore` and organized file structure

### Learning & Development

- **Manning Book Alignment**: Follows "Build an Orchestrator in Go" structure
- **Reference Material**: Docker commands, PostgreSQL setup, troubleshooting
- **Progressive Enhancement**: Supports iterative development approach
- **Best Practices**: Demonstrates professional Go project organization

## Development Workflow

The new structure supports modern development practices:

- Hot reloading for both Go (with tools) and React
- Separate development and production builds
- Integrated debugging across backend and frontend
- Professional CI/CD pipeline preparation
