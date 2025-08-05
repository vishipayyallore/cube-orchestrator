// File: src/task/state_machine.go

package task

import "fmt"

// State transition mapping defines which states can transition to which other states
var stateTransitionMap = map[State][]State{
	Pending:   {Scheduled},
	Scheduled: {Running, Failed},
	Running:   {Running, Completed, Failed},
	Completed: {},
	Failed:    {},
}

// ValidateStateTransition checks if a transition from one state to another is valid
func ValidateStateTransition(from, to State) bool {
	allowedStates := stateTransitionMap[from]
	for _, state := range allowedStates {
		if state == to {
			return true
		}
	}
	return false
}

// GetAllowedStates returns all states that can be transitioned to from the current state
func GetAllowedStates(currentState State) []State {
	return stateTransitionMap[currentState]
}

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

// TransitionTaskState safely transitions a task to a new state with validation
func (t *Task) TransitionState(newState State) error {
	if ValidateStateTransition(t.State, newState) {
		t.State = newState
		return nil
	}
	return fmt.Errorf("invalid state transition from %s to %s",
		StateToString(t.State), StateToString(newState))
}
