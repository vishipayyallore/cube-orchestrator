// File: src/orchestrator/cmd/main.go

package main

import (
	"cubeorchestrator/internal/docker"
	"cubeorchestrator/internal/manager"
	"cubeorchestrator/internal/node"
	"cubeorchestrator/internal/task"
	"cubeorchestrator/internal/worker"

	"fmt"
	"time"

	"github.com/docker/docker/client"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	t := task.Task{
		ID:     uuid.New(),
		Name:   "Task-1",
		State:  task.Pending,
		Image:  "Image-1",
		Cpu:    1.0,
		Memory: 1024,
		Disk:   1,
	}

	te := task.TaskEvent{
		ID:        uuid.New(),
		State:     task.Pending,
		Timestamp: time.Now(),
		Task:      t,
	}

	fmt.Printf("task: %v\n", t)
	fmt.Printf("task event: %v\n", te)

	// Demonstrate State Machine functionality
	fmt.Println("\n=== State Machine Demo ===")
	fmt.Printf("Initial task state: %s\n", task.StateToString(t.State))

	// Valid transition: Pending -> Scheduled
	if err := t.TransitionState(task.Scheduled); err != nil {
		fmt.Printf("Error transitioning to Scheduled: %v\n", err)
	} else {
		fmt.Printf("✅ Successfully transitioned to: %s\n", task.StateToString(t.State))
	}

	// Valid transition: Scheduled -> Running
	if err := t.TransitionState(task.Running); err != nil {
		fmt.Printf("Error transitioning to Running: %v\n", err)
	} else {
		fmt.Printf("✅ Successfully transitioned to: %s\n", task.StateToString(t.State))
	}

	// Invalid transition: Running -> Pending (should fail)
	if err := t.TransitionState(task.Pending); err != nil {
		fmt.Printf("❌ Expected error for invalid transition: %v\n", err)
	} else {
		fmt.Printf("Unexpected: transitioned to %s\n", task.StateToString(t.State))
	}

	// Valid transition: Running -> Completed
	if err := t.TransitionState(task.Completed); err != nil {
		fmt.Printf("Error transitioning to Completed: %v\n", err)
	} else {
		fmt.Printf("✅ Successfully transitioned to: %s\n", task.StateToString(t.State))
	}
	fmt.Println("=========================")

	// Chapter 4 Worker Demo - Real task processing
	fmt.Println("\n=== Chapter 4 Worker Task Processing Demo ===")
	w := worker.Worker{
		Name:  "worker-1",
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}
	fmt.Printf("worker: %v\n", w)
	w.CollectStats()

	// Create a test task for Chapter 4 workflow
	testTask := task.Task{
		ID:    uuid.New(),
		Name:  fmt.Sprintf("test-container-%s", uuid.New().String()[:8]), // Unique name
		State: task.Scheduled,
		Image: "strm/helloworld-http",
	}

	fmt.Printf("Adding task %s to worker queue\n", testTask.ID)
	w.AddTask(testTask)

	fmt.Println("Processing task via RunTask()...")
	result := w.RunTask()
	if result.Error != nil {
		fmt.Printf("Error processing task: %v\n", result.Error)
	} else {
		fmt.Printf("Task processed successfully, container: %s\n", result.ContainerId)
	}
	fmt.Println("=== Chapter 4 Worker Demo Complete ===\n")

	m := manager.Manager{
		Pending: *queue.New(),
		TaskDb:  make(map[string][]*task.Task),
		EventDb: make(map[string][]*task.TaskEvent),
		Workers: []string{w.Name},
	}

	fmt.Printf("manager: %v\n", m)
	m.SelectWorker()
	m.UpdateTasks()
	m.SendWork()

	n := node.Node{
		Name:   "Node-1",
		Ip:     "192.168.1.1",
		Cores:  4,
		Memory: 1024,
		Disk:   25,
		Role:   "worker",
	}

	fmt.Printf("node: %v\n", n)

	fmt.Printf("create a test container\n")
	dockerTask, createResult := createContainer()

	if dockerTask != nil && createResult != nil {
		time.Sleep(time.Second * 5)
		fmt.Printf("stopping container %s\n", createResult.ContainerId)
		_ = stopContainer(dockerTask, createResult.ContainerId)
	} else {
		fmt.Println("Skipping container stop due to creation failure")
	}
}

func createContainer() (*docker.Docker, *docker.DockerResult) {
	c := docker.Config{
		Name:  fmt.Sprintf("postgres-container-%s", uuid.New().String()[:8]), // Unique name
		Image: "postgres:13",
		Env: []string{
			"POSTGRES_USER=cube",
			"POSTGRES_PASSWORD=secret",
		},
	}

	dc, _ := client.NewClientWithOpts(client.FromEnv)
	d := docker.Docker{
		Client: dc,
		Config: c,
	}

	result := d.Run()
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		return nil, nil
	}

	fmt.Printf("Container %s is running with config %v\n", result.ContainerId, c)
	return &d, &result
}

func stopContainer(d *docker.Docker, id string) *docker.DockerResult {
	result := d.Stop(id)
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		return nil
	}

	fmt.Printf("Container %s has been stopped and removed\n", result.ContainerId)
	return &result
}
