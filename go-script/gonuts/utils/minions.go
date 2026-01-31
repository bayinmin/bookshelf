package utils

import (
	"context"
	"gonuts/config"
	"sync"

	"golang.org/x/time/rate"
)

// this is minions worker pool implementation

var limiter = rate.NewLimiter(10, 1) // 5 requests per second with a burst of 10

func Minion(jobs <-chan config.Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		limiter.Wait(context.Background()) // wait for rate limiter
		job()
	}
}

func RaiseMinions(numOfMinions int) (chan config.Job, *sync.WaitGroup) {
	jobs := make(chan config.Job, 100)
	var wg sync.WaitGroup
	for i := 0; i < numOfMinions; i++ {
		wg.Add(1)
		go Minion(jobs, &wg)
	}
	return jobs, &wg
}
