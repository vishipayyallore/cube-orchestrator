# Cube Orchestrator

🛠️ Learning to build an orchestrator in Go by following the book `Build an Orchestrator in Go (From Scratch)` from Manning Publications. Exploring concepts like process management, containers, and scheduling from the ground up.

## Project Structure

```text
cube-orchestrator/
├── .copilot/           # GitHub Copilot configuration
│   └── settings.json   # Copilot settings for the project
├── .github/            # GitHub configuration
│   └── copilot-instructions.md # Copilot context and guidelines
├── docs/               # Comprehensive documentation suite
│   ├── images/         # Documentation images and diagrams
│   ├── 00_README.md                    # Docs index and reading order
│   ├── 01_project-overview.md          # High-level project overview
│   ├── 02_project-structure.md         # Detailed structure documentation
│   ├── 03_configuration-verification.md# Environment setup verification
│   ├── 04_go-project-layout.md         # Go project structure guidelines
│   ├── 05_build-system.md              # Build system documentation
│   ├── 06_api-architecture.md          # API design patterns and structure
│   ├── 07_pkg-directory-plan.md        # Future API package planning
│   ├── 08_docker-images-reference.md   # Docker images used
│   ├── 09_docker-commands.md           # Docker commands reference
│   ├── 10_postgresql-primer.md         # PostgreSQL guide
│   └── 11_troubleshooting.md           # Common issues and solutions
├── scripts/            # Build and utility scripts
│   ├── build.sh        # Professional build script with timestamping
│   └── cleanup-builds.sh # Build artifact cleanup utility
├── src/                # Source code directory
│   └── orchestrator/   # Main orchestrator application
│       ├── cmd/main.go # Main application with orchestrator demo
│       ├── internal/   # Private application packages
│       │   ├── runtime/    # Docker runtime abstraction (DockerWrapper)
│       │   ├── manager/    # Orchestrator manager component
│       │   ├── worker/     # Worker node implementation
│       │   ├── node/       # Node abstraction and resources
│       │   ├── scheduler/  # Task scheduling algorithms
│       │   └── task/       # Task definitions and state machine
│       ├── pkg/        # Public API packages (planned)
│       ├── go.mod      # Go module definition
│       └── go.sum      # Dependency checksums
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

**Security Update**: This project now uses Docker v28.3.3+incompatible to address security vulnerabilities:

- ✅ **Fixed GO-2023-1699**: Docker Swarm encrypted overlay network authentication issue
- ✅ **Fixed GO-2023-1700**: Docker Swarm encrypted overlay network encryption issue  
- ✅ **Fixed GO-2023-1701**: Docker Swarm single endpoint authentication issue

**Note**: If you encounter API compatibility issues with newer Docker versions, the codebase has been updated to use the v28 API structure.

**Solutions for older versions**:

1. **Recommended**: Use Docker v28.3.3+incompatible (current implementation)
2. **Alternative approach**: Use the Moby client directly:

   ```bash
   go get github.com/moby/moby/client
   ```

3. **Wait for book updates**: The book may provide updated import instructions

## Getting Started

### Running the Application

To run the cube orchestrator demo:

```bash
cd src/orchestrator
go run ./cmd/main.go
```

![Cube Orchestrator Demo Output](docs/images/After_Ch_2.PNG)

### Docker Container Integration (Chapter 3)

The application now includes Docker container management capabilities:

![Chapter 3 - Docker Integration](docs/images/After_Ch_3.PNG)

### Docker Setup

For Docker commands and container management instructions, see [Docker Commands](docs/09_docker-commands.md).

### Troubleshooting

If you encounter any issues with dependencies, imports, or compilation, see the [Troubleshooting Guide](docs/11_troubleshooting.md).

## Development

### Docs quality checks (local)

Run Markdown lint against README and all docs before opening a PR:

```powershell
# From repo root
npx --yes markdownlint-cli2 "README.md" "docs/**/*.md"
```

This uses the repository's .markdownlint.json automatically.

### Link check (Lychee)

Run a quick local link check using Lychee (via Docker):

```powershell
# Extract links only (does not validate)
docker run --rm -w /input -v "${PWD}:/input" lycheeverse/lychee:latest --config lychee.toml --no-progress --dump README.md docs/**/*.md

# Validate links (recommended; matches CI behavior)
docker run --rm -w /input -v "${PWD}:/input" lycheeverse/lychee:latest --config lychee.toml --no-progress README.md docs/**/*.md
```

### Optional: hooks for docs checks

Enable Git hooks in this repo (applies to both pre-commit and pre-push hooks):

```powershell
# From repo root (one-time setup)
git config core.hooksPath .githooks
```

Pre-commit hook: runs on any commit that includes README/docs changes.

Pre-push hook (lightweight, recommended): runs only when pushing to main and only if README/docs changed in the push range.

Skip temporarily if needed:

```powershell
# Skip markdown lint once
$env:SKIP_DOCS_LINT = '1'; git commit -m "skip docs lint once"; Remove-Item Env:SKIP_DOCS_LINT

# Skip link check once
$env:SKIP_LINK_CHECK = '1'; git commit -m "skip link check once"; Remove-Item Env:SKIP_LINK_CHECK
```

Skip on push instead (for the pre-push hook):

```powershell
# Skip markdown lint for the next push only
$env:SKIP_DOCS_LINT = '1'; git push; Remove-Item Env:SKIP_DOCS_LINT

# Skip link check for the next push only
$env:SKIP_LINK_CHECK = '1'; git push; Remove-Item Env:SKIP_LINK_CHECK
```
