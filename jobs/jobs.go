package jobs

import (
	"datahandler_go/jobs/samples"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func Jobs() {
	c := cron.New()

	// Define job schedules with their corresponding functions
	jobsList := []struct {
		name     string
		timeout  time.Duration
		interval string
		jobFunc  func() // Job function to be executed
	}{
		{
			name:     "samples/postgres_sample",
			timeout:  0, // run on start
			interval: "",
			jobFunc:  samples.Postgres_Sample_Job, // Reference to the actual function
		},
		{
			name:     "samples/mongo_sample",
			timeout:  0, // run on start
			interval: "",
			jobFunc:  samples.Mongo_Sample_Job, // Reference to the actual function
		},
	}

	for _, job := range jobsList {
		if job.interval == "" {
			// Schedule the job to run immediately as a goroutine
			_, err := c.AddFunc("@every 0s", func() {
				go job.jobFunc() // Run job function as a goroutine
			})
			if err != nil {
				log.Fatal(err)
			}
		} else {
			// Schedule the job with the given interval (cron syntax) as a goroutine
			_, err := c.AddFunc(job.interval, func() {
				go job.jobFunc() // Run job function as a goroutine
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// Start the scheduler
	c.Start()

	// Keep the program running
	select {}
}
