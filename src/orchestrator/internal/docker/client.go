package docker

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

type Docker struct {
	Client *client.Client
	Config Config
}

type DockerResult struct {
	Error       error
	Action      string
	ContainerId string
	Result      string
}

func (d *Docker) Run() DockerResult {
	ctx := context.Background()
	reader, err := d.Client.ImagePull(
		ctx, d.Config.Image, image.PullOptions{})
	if err != nil {
		log.Printf("Error pulling image %s: %v\n", d.Config.Image, err)
		return DockerResult{Error: err}
	}
	io.Copy(os.Stdout, reader)

	rp := container.RestartPolicy{
		Name: container.RestartPolicyMode(d.Config.RestartPolicy),
	}

	r := container.Resources{
		Memory: d.Config.Memory,
	}

	cc := container.Config{
		Image:        d.Config.Image,
		Tty:          false,
		Env:          d.Config.Env,
		ExposedPorts: d.Config.ExposedPorts,
	}

	hc := container.HostConfig{
		RestartPolicy:   rp,
		Resources:       r,
		PublishAllPorts: true,
	}

	resp, err := d.Client.ContainerCreate(ctx, &cc, &hc, nil, nil, d.Config.Name)
	if err != nil {
		log.Printf("Error creating container using image %s: %v\n", d.Config.Image, err)
		return DockerResult{Error: err}
	}

	if err = d.Client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Printf("Error starting container %s: %v\n", resp.ID, err)
		return DockerResult{Error: err}
	}

	out, err := d.Client.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		log.Printf("Error getting logs for container %s: %v\n", resp.ID, err)
		return DockerResult{Error: err}
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return DockerResult{ContainerId: resp.ID, Action: "start", Result: "success"}
}

func (d *Docker) Stop(id string) DockerResult {
	log.Printf("Attempting to stop container %v", id)
	ctx := context.Background()
	err := d.Client.ContainerStop(ctx, id, container.StopOptions{})
	if err != nil {
		log.Printf("Error stopping container %s: %v\n", id, err)
		return DockerResult{Error: err}
	}

	err = d.Client.ContainerRemove(ctx, id, container.RemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   false,
		Force:         false,
	})
	if err != nil {
		log.Printf("Error removing container %s: %v\n", id, err)
		return DockerResult{Error: err}
	}

	return DockerResult{Action: "stop", Result: "success", Error: nil}
}

func (d *Docker) Remove(containerId string) DockerResult {
	ctx := context.Background()

	err := d.Client.ContainerRemove(ctx, containerId, container.RemoveOptions{})
	if err != nil {
		log.Printf("Error removing container %s: %v\n", containerId, err)
		return DockerResult{Error: err}
	}

	log.Printf("Container %s removed successfully\n", containerId)
	return DockerResult{
		Action:      "remove",
		ContainerId: containerId,
		Result:      "success",
	}
}

// NewConfig creates a Docker configuration from a task
// This bridges Chapter 4's pattern with our modern implementation
func NewConfig(t interface{}) *Config {
	// For now, we'll use a simple approach that can handle different task types
	// This will be properly integrated with task.Task in the next step

	// Create a default config that works for Chapter 4 examples
	config := &Config{
		Name:          "test-container",
		Image:         "strm/helloworld-http",
		RestartPolicy: "no",
		ExposedPorts:  make(nat.PortSet),
		AttachStdin:   false,
		AttachStdout:  true,
		AttachStderr:  true,
		Env:           []string{},
	}

	// TODO: In the next step, we'll add proper task.Task integration here
	// This placeholder allows Chapter 4 patterns to work immediately

	return config
}

// NewDocker creates a new Docker client with the given configuration
// This matches Chapter 4's pattern while using our modern Docker client
func NewDocker(c *Config) *Docker {
	dc, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Printf("Error creating Docker client: %v", err)
		// Return nil client - calling code should check for errors
		return &Docker{
			Client: nil,
			Config: *c,
		}
	}

	return &Docker{
		Client: dc,
		Config: *c,
	}
}
