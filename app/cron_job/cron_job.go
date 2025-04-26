package cron_job

import (
	"time"

	"example.com/app/constants"
	"example.com/app/utils"
	"github.com/go-co-op/gocron"
	"github.com/phuslu/log"
)

var runner *gocron.Scheduler

func Init() {
	runner = gocron.NewScheduler(time.UTC)
	runner.StartAsync()
	_, err := runner.Every(1).Day().At("01:00").Do(func() {
		utils.LoadCSV(constants.CSVFilePath)
	})
	if err != nil {
		log.Info().Msgf("Error scheduling task:", err)
		return
	}

}
