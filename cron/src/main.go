package main

import (
	"cron/src/handler"
	"cron/src/meta"
	"fmt"
	"github.com/gorhill/cronexpr"
	"net/http"
	"time"
)

type CronJob struct {
	expr *cronexpr.Expression
	next time.Time
}

func startCron() {
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

func main() {
	meta.Get()
	// meta.Store()

	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/ok", handler.UploadOkHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Start server failed! error: %s\n", err)
	}
}
