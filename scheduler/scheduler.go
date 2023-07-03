package scheduler

import (
	"fmt"
	"time"
)

type Scheduler struct {
}

func New() Scheduler {
	return Scheduler{}
}

// Start Long running process
func (s Scheduler) Start(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("exiting scheduler...")
			return
		default:
			fmt.Println("scheduler", time.Now())
			time.Sleep(time.Second * 1)
		}
	}
}
