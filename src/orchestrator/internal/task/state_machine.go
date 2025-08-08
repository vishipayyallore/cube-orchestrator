// File: src/orchestrator/internal/task/state_machine.go

package task

import "fmt"

// State transition mapping defines which states can transition to which other states
// Updated to match Chapter 4's expected behavior
var stateTransitionMap = map[State][]State{
	Pending:   {Scheduled},
	Scheduled: {Scheduled, Running, Failed}, // Chapter 4: Allow staying in Scheduled
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

// ValidStateTransition is an alias for ValidateStateTransition
// This provides Chapter 4 compatibility (the book uses this exact function name)
func ValidStateTransition(from, to State) bool {
	return ValidateStateTransition(from, to)
}

// GetAllowedStates returns all states that can be transitioned to from the current state
func GetAllowedStates(currentState State) []State {
	return stateTransitionMap[currentState]
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
