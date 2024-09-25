package jobs

import (
	"datahandler_go/jobs/samples"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type job struct {
	name     string
	timeout  time.Duration // Duration before the job starts (>= 0)
	interval uint          // Interval in seconds for cron to run job after initial execution (>= 0)
	jobFunc  func()        // Job function to be executed
}

func scheduleJobWithInterval(c *cron.Cron, newJob job) {
	// Convert the interval to a cron-compatible string like "@every 60s"
	interval := fmt.Sprintf("@every %ds", newJob.interval)

	// Schedule the job with cron
	_, err := c.AddFunc(interval, func() {
		go newJob.jobFunc()
	})

	if err != nil {
		log.Fatalf("Failed to schedule job %s: %v", newJob.name, err)
	}

	log.Printf("Scheduled job %s to run every %d seconds", newJob.name, newJob.interval)
}

func RunJobs() {
	c := cron.New()

	// Define job schedules with their corresponding functions
	jobsList := []job{
		{
			name:     "samples/postgres_sample",
			timeout:  0, // Run on start
			interval: 60,
			jobFunc:  samples.Postgres_Sample_Job,
		},
		{
			name:     "samples/mongo_sample",
			timeout:  10 * time.Second, // Run after 60 seconds
			interval: 60,
			jobFunc:  samples.Mongo_Sample_Job,
		},
	}

	for _, job := range jobsList {
		// Ensure valid timeout value
		if job.timeout < 0 {
			log.Fatalf("Job %s has an invalid timeout value: %v", job.name, job.timeout)
		}

		// Schedule the job with the initial delay (timeout) using time.AfterFunc
		if job.timeout > 0 {
			log.Printf("Scheduling job %s to start after %v", job.name, job.timeout)
			time.AfterFunc(job.timeout, func() {
				// Run the job immediately after the timeout
				go job.jobFunc()
				// Schedule the job to run at the specified interval after the timeout
				scheduleJobWithInterval(c, job)
			})
		} else {
			// If no timeout, schedule recurring runs immediately
			go job.jobFunc()
			scheduleJobWithInterval(c, job)
		}
	}

	// Start the scheduler
	c.Start()
}
