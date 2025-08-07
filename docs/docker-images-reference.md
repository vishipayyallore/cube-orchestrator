# Docker Images Reference

This document provides detailed information about the Docker images used in the cube-orchestrator project, their purposes, and how they're integrated into the orchestration workflow.

## Overview

The cube-orchestrator project uses various Docker images for different purposes:

- **Testing and demonstration**: Simple, lightweight images for validating orchestration functionality
- **Database services**: PostgreSQL for data persistence and state management
- **Application workloads**: Various containerized applications that demonstrate real-world orchestration scenarios

## Images Used

### 1. strm/helloworld-http

**Purpose**: Lightweight HTTP server for testing container orchestration

**Description**:
The `strm/helloworld-http` image is a minimal Docker image that runs a simple HTTP web server. It's specifically designed for testing containerized applications and orchestration systems.

**Key Characteristics**:

- **Lightweight**: Small image size (~10-20MB) for fast downloads and startup
- **Simple**: Serves a basic "Hello World" HTTP response
- **Reliable**: Predictable behavior ideal for testing scenarios
- **HTTP endpoint**: Typically serves on port 80 or 8080
- **No dependencies**: Self-contained with no external service requirements

**Usage in cube-orchestrator**:

1. **Worker Task Processing Demo** (`main.go:89`):

   ```go
   testTask := task.Task{
       ID:    uuid.New(),
       Name:  fmt.Sprintf("test-container-%s", uuid.New().String()[:8]),
       State: task.Scheduled,
       Image: "strm/helloworld-http",  // ← Used here for testing
   }
   ```

2. **Runtime Client Testing** (`internal/runtime/client.go:153`):
   - Used as a test workload for container lifecycle management
   - Demonstrates task scheduling, container creation, and state transitions

**Expected Behavior**:

- Container starts quickly (usually within 1-2 seconds)
- Exposes HTTP endpoint serving "Hello World" response
- Runs indefinitely until stopped by orchestrator
- Minimal resource consumption (CPU: ~0.1%, Memory: ~10MB)

**Testing Scenarios**:

- ✅ Container creation and startup
- ✅ Task state transitions (Pending → Scheduled → Running → Completed)
- ✅ Worker queue processing
- ✅ Container lifecycle management (start/stop/remove)
- ✅ Resource allocation and monitoring

### 2. postgres:13

**Purpose**: Database service for data persistence and advanced orchestration scenarios

**Description**:
PostgreSQL 13 official image used for demonstrating database container orchestration and service dependencies.

**Configuration in cube-orchestrator**:

```go
c := runtime.Config{
    Name:  fmt.Sprintf("postgres-container-%s", uuid.New().String()[:8]),
    Image: "postgres:13",
    Env: []string{
        "POSTGRES_USER=cube",
        "POSTGRES_PASSWORD=secret",
    },
}
```

**Key Features**:

- **Database service**: Full PostgreSQL database functionality
- **Persistent storage**: Can be configured with volume mounts
- **Environment configuration**: Username, password, and database settings
- **Service dependencies**: Demonstrates complex application orchestration

**Usage Patterns**:

- Container creation with environment variables
- Service lifecycle management
- Database initialization and configuration
- Network connectivity testing
- Resource monitoring and management

## Image Selection Criteria

When choosing Docker images for the cube-orchestrator project, consider:

### For Testing and Development

- **Size**: Prefer smaller images for faster testing cycles
- **Startup time**: Images that start quickly reduce development iteration time
- **Predictability**: Consistent behavior across different environments
- **HTTP endpoints**: For connectivity and health checking

### For Production Demonstrations

- **Real-world relevance**: Images that represent actual application workloads
- **Configuration options**: Environment variables and runtime parameters
- **Resource requirements**: Various CPU and memory profiles for scheduling tests
- **Service dependencies**: Multi-container applications for complex orchestration

## Integration with Orchestrator Components

### Task Definition

```go
type Task struct {
    ID     uuid.UUID
    Name   string
    State  State
    Image  string        // ← Docker image specification
    Cpu    float64
    Memory int64
    Disk   int64
}
```

### Worker Processing

1. **Task Reception**: Worker receives task with image specification
2. **Container Creation**: Runtime client pulls and creates container from image
3. **Lifecycle Management**: Start, monitor, and stop container as needed
4. **State Tracking**: Update task state based on container status

### Runtime Integration

The `runtime.Runtime` component handles:

- Image pulling and caching
- Container creation with specified image
- Port mapping and network configuration
- Environment variable injection
- Container monitoring and health checks

## Best Practices

### Image Naming

- Use specific tags rather than `latest` for reproducibility
- Include registry URL for private registries
- Consider image digest for immutable deployments

### Resource Planning

- **strm/helloworld-http**: Minimal resources (0.1 CPU, 64MB RAM)
- **postgres:13**: Moderate resources (0.5 CPU, 256MB RAM minimum)
- Plan container resource limits based on image requirements

### Testing Strategy

1. **Unit Tests**: Mock container behavior for fast testing
2. **Integration Tests**: Use lightweight images like `strm/helloworld-http`
3. **Performance Tests**: Use resource-intensive images to validate scheduling
4. **Failure Tests**: Test container failures and recovery scenarios

## Troubleshooting

### Common Issues

**Image Pull Failures**:

- Verify Docker daemon is running
- Check network connectivity to registry
- Validate image name and tag

**Container Startup Failures**:

- Review container logs via Docker CLI
- Check environment variable configuration
- Verify resource availability on worker nodes

**Port Conflicts**:

- Use dynamic port allocation
- Implement port conflict detection
- Configure proper network isolation

### Debugging Commands

```bash
# List running containers
docker ps

# Check container logs
docker logs <container-id>

# Inspect container configuration
docker inspect <container-id>

# Monitor resource usage
docker stats <container-id>
```

## Future Considerations

### Additional Test Images

- **nginx**: Web server with configurable content
- **redis**: In-memory data store for caching scenarios
- **busybox**: Minimal utilities for debugging and testing

### Production Images

- Application-specific images for real workloads
- Multi-stage builds for optimized production containers
- Security-hardened images with minimal attack surface

### Registry Integration

- Private registry support for enterprise scenarios
- Image vulnerability scanning integration
- Automated image updates and rollback capabilities

---

*This documentation is part of the cube-orchestrator learning project. For more information about the overall architecture, see [project-overview.md](project-overview.md).*
