// File: src/orchestrator/internal/runtime/docker_wrapper.go

package runtime

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/moby/moby/pkg/stdcopy"

	"cubeorchestrator/internal/task"
)

type Config struct {
	Name          string
	AttachStdin   bool
	AttachStdout  bool
	AttachStderr  bool
	ExposedPorts  nat.PortSet
	Cmd           []string
	Image         string
	Cpu           float64
	Memory        int64
	Disk          int64
	Env           []string
	RestartPolicy string
}

type DockerWrapper struct {
	Client *client.Client
	Config Config
}

type RuntimeResult struct {
	Error       error
	Action      string
	ContainerId string
	Result      string
}

func (r *DockerWrapper) Run() RuntimeResult {
	ctx := context.Background()
	reader, err := r.Client.ImagePull(
		ctx, r.Config.Image, image.PullOptions{})
	if err != nil {
		log.Printf("Error pulling image %s: %v\n", r.Config.Image, err)
		return RuntimeResult{Error: err}
	}
	io.Copy(os.Stdout, reader)

	rp := container.RestartPolicy{
		Name: container.RestartPolicyMode(r.Config.RestartPolicy),
	}

	resources := container.Resources{
		Memory: r.Config.Memory,
	}

	cc := container.Config{
		Image:        r.Config.Image,
		Tty:          false,
		Env:          r.Config.Env,
		ExposedPorts: r.Config.ExposedPorts,
	}

	hc := container.HostConfig{
		RestartPolicy:   rp,
		Resources:       resources,
		PublishAllPorts: true,
	}

	resp, err := r.Client.ContainerCreate(ctx, &cc, &hc, nil, nil, r.Config.Name)
	if err != nil {
		log.Printf("Error creating container using image %s: %v\n", r.Config.Image, err)
		return RuntimeResult{Error: err}
	}

	if err = r.Client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Printf("Error starting container %s: %v\n", resp.ID, err)
		return RuntimeResult{Error: err}
	}

	out, err := r.Client.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		log.Printf("Error getting logs for container %s: %v\n", resp.ID, err)
		return RuntimeResult{Error: err}
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return RuntimeResult{ContainerId: resp.ID, Action: "start", Result: "success"}
}

func (r *DockerWrapper) Stop(id string) RuntimeResult {
	log.Printf("Attempting to stop container %v", id)
	ctx := context.Background()
	err := r.Client.ContainerStop(ctx, id, container.StopOptions{})
	if err != nil {
		log.Printf("Error stopping container %s: %v\n", id, err)
		return RuntimeResult{Error: err}
	}

	err = r.Client.ContainerRemove(ctx, id, container.RemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   false,
		Force:         false,
	})
	if err != nil {
		log.Printf("Error removing container %s: %v\n", id, err)
		return RuntimeResult{Error: err}
	}

	return RuntimeResult{Action: "stop", Result: "success", Error: nil}
}

func (r *DockerWrapper) Remove(containerId string) RuntimeResult {
	ctx := context.Background()

	err := r.Client.ContainerRemove(ctx, containerId, container.RemoveOptions{})
	if err != nil {
		log.Printf("Error removing container %s: %v\n", containerId, err)
		return RuntimeResult{Error: err}
	}

	log.Printf("Container %s removed successfully\n", containerId)
	return RuntimeResult{
		Action:      "remove",
		ContainerId: containerId,
		Result:      "success",
	}
}

// NewConfig creates a runtime configuration from a task
// This bridges Chapter 4's pattern with our modern implementation
func NewConfig(t interface{}) *Config {
	// Handle both *task.Task and task.Task
	var taskObj *task.Task

	switch v := t.(type) {
	case *task.Task:
		taskObj = v
	case task.Task:
		taskObj = &v
	default:
		// Fallback for unexpected types
		log.Printf("Warning: NewConfig received unexpected type %T, using default config", t)
		return &Config{
			Name:          "default-container",
			Image:         "strm/helloworld-http",
			RestartPolicy: "no",
			ExposedPorts:  make(nat.PortSet),
			AttachStdin:   false,
			AttachStdout:  true,
			AttachStderr:  true,
			Env:           []string{},
		}
	}

	// Create config from task properties
	exposedPorts := taskObj.ExposedPorts
	if exposedPorts == nil {
		exposedPorts = make(nat.PortSet)
	}

	restartPolicy := taskObj.RestartPolicy
	if restartPolicy == "" {
		restartPolicy = "no"
	}

	return &Config{
		Name:          taskObj.Name,
		Image:         taskObj.Image,
		ExposedPorts:  exposedPorts,
		Cpu:           taskObj.Cpu,
		Memory:        taskObj.Memory,
		Disk:          taskObj.Disk,
		RestartPolicy: restartPolicy,
		AttachStdin:   false,
		AttachStdout:  true,
		AttachStderr:  true,
		Env:           []string{},
	}
}

// NewRuntime creates a new runtime client with the given configuration
// This matches Chapter 4's pattern while using our modern Docker client
func NewRuntime(c *Config) *DockerWrapper {
	dc, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Printf("Error creating Docker client: %v", err)
		// Return nil client - calling code should check for errors
		return &DockerWrapper{
			Client: nil,
			Config: *c,
		}
	}

	return &DockerWrapper{
		Client: dc,
		Config: *c,
	}
}
