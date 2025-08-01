// File: src/scheduler/scheduler.go

package scheduler

type Scheduler interface {
	SelectCandidateNodes()
	Score()
	Pick()
}
