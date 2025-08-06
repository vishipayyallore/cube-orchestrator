// File: src/orchestrator/internal/worker/worker.go

package worker

import (
	"cubeorchestrator/internal/runtime"
	"cubeorchestrator/internal/task"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

// AddTask adds a task to the worker's queue for processing
// This is the entry point for Chapter 4's task workflow
func (w *Worker) AddTask(t task.Task) {
	w.Queue.Enqueue(t)
}

func (w *Worker) CollectStats() {
	fmt.Println("I will collect stats")
}

func (w *Worker) RunTask() runtime.RuntimeResult {
	// Dequeue the next task from the worker's queue
	t := w.Queue.Dequeue()
	if t == nil {
		log.Println("No tasks in the queue")
		return runtime.RuntimeResult{Error: nil}
	}

	// Get the task from the queue (type assertion)
	taskQueued := t.(task.Task)

	// Get the persisted task from database or create new entry
	taskPersisted := w.Db[taskQueued.ID]
	if taskPersisted == nil {
		taskPersisted = &taskQueued
		w.Db[taskQueued.ID] = &taskQueued
	}

	var result runtime.RuntimeResult // Validate state transition before processing
	if task.ValidateStateTransition(taskPersisted.State, taskQueued.State) {
		switch taskQueued.State {
		case task.Scheduled:
			// Start the task
			result = w.StartTask(taskQueued)
		case task.Completed:
			// Stop the task
			result = w.StopTask(taskQueued)
		default:
			result.Error = errors.New("invalid state for task processing")
		}
	} else {
		err := fmt.Errorf("invalid transition from %v to %v",
			task.StateToString(taskPersisted.State),
			task.StateToString(taskQueued.State))
		result.Error = err
		log.Printf("State transition error for task %v: %v", taskQueued.ID, err)
	}

	return result
}

func (w *Worker) StartTask(t task.Task) runtime.RuntimeResult {
	// Set start time
	t.StartTime = time.Now().UTC()

	// Create runtime configuration from task
	config := runtime.NewConfig(&t)
	r := runtime.NewRuntime(config)

	// Check if runtime client was created successfully
	if r.Client == nil {
		err := errors.New("failed to create runtime client")
		log.Printf("Error creating runtime client for task %v: %v", t.ID, err)
		t.State = task.Failed
		w.Db[t.ID] = &t
		return runtime.RuntimeResult{Error: err}
	}

	// Run the container
	result := r.Run()
	if result.Error != nil {
		log.Printf("Error running task %v: %v", t.ID, result.Error)
		t.State = task.Failed
		w.Db[t.ID] = &t
		return result
	}

	// Update task state and container ID
	t.ContainerID = result.ContainerId
	t.State = task.Running
	w.Db[t.ID] = &t

	log.Printf("Started container %v for task %v", result.ContainerId, t.ID)
	return result
}

func (w *Worker) StopTask(t task.Task) runtime.RuntimeResult {
	// Create runtime configuration for stopping
	config := runtime.NewConfig(&t)
	r := runtime.NewRuntime(config)

	// Check if runtime client was created successfully
	if r.Client == nil {
		err := errors.New("failed to create runtime client")
		log.Printf("Error creating runtime client for stopping task %v: %v", t.ID, err)
		return runtime.RuntimeResult{Error: err}
	}

	// Stop and remove the container
	result := r.Stop(t.ContainerID)
	if result.Error != nil {
		log.Printf("Error stopping container %v: %v", t.ContainerID, result.Error)
		return result
	}

	// Update task state and finish time
	t.FinishTime = time.Now().UTC()
	t.State = task.Completed
	w.Db[t.ID] = &t

	log.Printf("Stopped container %v for task %v", t.ContainerID, t.ID)
	return result
}
