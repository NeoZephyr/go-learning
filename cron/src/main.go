package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr *cronexpr.Expression
	next time.Time
}

func main() {
	var (
		expr *cronexpr.Expression
		err error
	)

	now := time.Now()
	var scheduleTable = make(map[string]*CronJob)

	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	scheduleTable["job1"] = &CronJob{
		expr: expr,
		next: expr.Next(now),
	}

	if expr, err = cronexpr.Parse("*/1 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	scheduleTable["job2"] = &CronJob{
		expr: expr,
		next: expr.Next(now),
	}

	go func() {
		for {
			now = time.Now()

			for jobName, cronJob := range scheduleTable {
				if now.After(cronJob.next) || now.Equal(cronJob.next) {
					go func(jobName string) {
						fmt.Printf("exec job: %v, now: %v, next: %v\n", jobName, now, cronJob.next)
					}(jobName)

					cronJob.next = cronJob.expr.Next(now)
				}
			}

			select {
			case <- time.NewTimer(100 * time.Millisecond).C:
			}
		}
	}()

	time.Sleep(100 * time.Second)
}
