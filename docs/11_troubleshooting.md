# Troubleshooting Guide

This document contains solutions to common issues encountered while building the cube-orchestrator.

## Docker Client Import Issues

### Security Update - Docker v28.3.3

**Current Implementation**: This project uses Docker v28.3.3+incompatible to fix critical security vulnerabilities:

- GO-2023-1699: Docker Swarm encrypted overlay network authentication
- GO-2023-1700: Docker Swarm encrypted overlay network encryption  
- GO-2023-1701: Docker Swarm single endpoint authentication

### API Changes in Docker v28

Docker v28 introduced API changes. The code has been updated accordingly:

```go
// Old API (v20.x)
types.ImagePullOptions{}
types.ContainerStartOptions{}
types.ContainerRemoveOptions{}

// New API (v28.x) 
image.PullOptions{}
container.StartOptions{}
container.RemoveOptions{}
```

### Problem

When importing Docker client libraries, you may encounter module path conflicts:

```text
Command 'gopls.go_get_package' failed: Error: err: exit status 1: stderr: 
go: finding module for package github.com/docker/docker/client 
go: github.com/docker/docker/client@v0.0.0-20250801223428-0f9c087c91c0: 
parsing go.mod: module declares its path as: github.com/moby/moby/client 
but was required as: github.com/docker/docker/client
```

### Root Cause

Historically, the Docker client library organization caused path confusion between `github.com/docker/docker` and `github.com/moby/moby`. This project uses the official `github.com/docker/docker` SDK at v28.x and does not depend on `github.com/moby/moby` directly.

### Solutions

#### Preferred: Use official Docker SDK v28.x (current project setup)

The repo is already configured to use the official Docker SDK:

```bash
go get github.com/docker/docker@v28.3.3+incompatible
```

Notes:

- Some package names changed in v28 (e.g., `image.PullOptions`, `container.StartOptions`).
- The `stdcopy` import should be `github.com/docker/docker/pkg/stdcopy`.

#### Optional (early chapters without Docker): Remove Docker client usage

If you are experimenting with non-Docker chapters, you can remove Docker client usage temporarily and run `go mod tidy` to simplify dependencies.

### Verification

After applying any solution, verify the fix:

```bash
# Clean and rebuild
go clean -modcache
go mod tidy

# Test compilation (from workspace root)
cd src/orchestrator/cmd
go run main.go
```

## Module Cache Issues

### Module Cache Problem

Sometimes the Go module cache can become corrupted or contain conflicting versions.

### Module Cache Solution

```bash
# Clear module cache
go clean -modcache

# Reinitialize modules
go mod tidy

# If issues persist, delete go.sum and regenerate
rm go.sum
go mod tidy
```

## Import Path Issues

### Import Path Problem

Import paths don't match the actual module structure.

### Import Path Solution

1. **Check go.mod file** to see the module name:

   ```go
   module cubeorchestrator
   ```

2. **Update imports accordingly** (no `src/` in import paths; use module path + package):

   ```go
   import (
      "cubeorchestrator/internal/task"
      "cubeorchestrator/internal/worker"
      "cubeorchestrator/internal/runtime"
   )
   ```

## Dependency Version Conflicts

### Dependency Conflict Problem

Different packages require conflicting versions of the same dependency.

### Dependency Conflict Solution

1. **Check for conflicts**:

   ```bash
   go mod why -m github.com/conflicting/package
   ```

2. **Force specific version**:

   ```bash
   go get github.com/package@v1.2.3
   ```

3. **Use replace directive in go.mod**:

   ```go
   replace github.com/old/package => github.com/new/package v1.0.0
   ```

## Building and Running Issues

### Runtime Problem

Code compiles but fails at runtime.

### Common Solutions

1. **Check working directory**:

   ```bash
   # From workspace root
   cd src/orchestrator/cmd
   go run main.go
   ```

2. **Verify all files are present**:

   ```bash
   # Check project structure
   find . -name "*.go" -type f
   ```

3. **Check for missing dependencies**:

   ```bash
   go mod verify
   go mod download
   ```

   ## Docker Daemon Availability

   ### Symptom

   - Docker-based demos are skipped with a message like: "Docker unavailable: REASON".

   ### Root Causes

   - Docker Desktop/daemon not running
   - Insufficient permissions to access the Docker socket
   - Windows named pipe not reachable: `\\.\pipe\docker_engine`

   ### Fixes

   - Start Docker Desktop (Windows/macOS) or the Docker service (Linux)
   - On Windows, verify the named pipe exists and your user has access
   - If using WSL2, ensure integration is enabled for your distribution
   - Re-run from an elevated shell if required (or adjust group membership on Linux)

   The code uses a pre-flight `DockerAvailable()` check to avoid noisy failures and will continue the rest of the demo when Docker isn't accessible.

## VS Code and gopls Issues

### VS Code Problem

VS Code Go extension shows errors or fails to work properly.

### VS Code Solution

1. **Reload VS Code window**: Ctrl+Shift+P → "Developer: Reload Window"

2. **Restart Go language server**: Ctrl+Shift+P → "Go: Restart Language Server"

3. **Clear gopls cache**:

   ```bash
   rm -rf ~/.cache/go-build
   go clean -cache
   ```

4. **Check Go extension settings** in VS Code

## Getting Help

If you encounter issues not covered here:

1. **Check the book's errata** or updated examples
2. **Review the project's GitHub issues** (if available)
3. **Consult Go module documentation**: `go help modules`
4. **Use Go community resources** like golang.org/help

## Useful Commands

```bash
# Module management
go mod init <module-name>
go mod tidy
go mod verify
go mod download
go mod graph

# Dependency analysis
go list -m all
go mod why -m <module>
go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all

# Cleaning
go clean -cache
go clean -modcache
go clean -testcache
```
