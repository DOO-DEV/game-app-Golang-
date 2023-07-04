package scheduler

import (
	"fmt"
	"game-app/param"
	"game-app/service/matchingservice"
	"github.com/go-co-op/gocron"
	"sync"
	"time"
)

type Config struct {
	MatchWaitedUsersIntervalInSeconds int `koanf:"match_waited_users_interval_in_seconds"`
}

type Scheduler struct {
	sch      *gocron.Scheduler
	matchSvc matchingservice.Service
	config   Config
}

func New(matchSvc matchingservice.Service, config Config) Scheduler {
	return Scheduler{
		sch:      gocron.NewScheduler(time.UTC),
		matchSvc: matchSvc,
		config:   config,
	}
}

// Start Long running process
func (s Scheduler) Start(done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	s.sch.Every(s.config.MatchWaitedUsersIntervalInSeconds).Second().Do(s.MatchWaitedUser)

	s.sch.StartAsync()
	<-done
	fmt.Println("exiting scheduler...")
	s.sch.Stop()
}

func (s Scheduler) MatchWaitedUser() {
	fmt.Println("scheduler match user")
	s.matchSvc.MatchUsers(param.MatchedWaitedUsersRequest{})
}
