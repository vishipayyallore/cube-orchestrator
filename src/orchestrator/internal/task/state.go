// File: src/orchestrator/internal/task/state.go

package task

// State represents the current state of a task in the orchestration lifecycle
type State int

// Task states defining the lifecycle of a task
const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

// StateToString converts a State to its string representation for logging/debugging
func StateToString(state State) string {
	switch state {
	case Pending:
		return "Pending"
	case Scheduled:
		return "Scheduled"
	case Running:
		return "Running"
	case Completed:
		return "Completed"
	case Failed:
		return "Failed"
	default:
		return "Unknown"
	}
}
