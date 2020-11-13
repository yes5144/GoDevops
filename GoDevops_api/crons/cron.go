package crons

import "github.com/robfig/cron"

var crontab = cron.New()

// Start xxx
func Start() {
	crontab.Start()
}

// Stop xxx
func Stop() {
	crontab.Stop()
}
