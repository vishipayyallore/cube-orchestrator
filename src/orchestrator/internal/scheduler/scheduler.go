// File: src/orchestrator/internal/scheduler/scheduler.go

package scheduler

type Scheduler interface {
	SelectCandidateNodes()
	Score()
	Pick()
}
