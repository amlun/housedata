package main

import (
	"github.com/amlun/housedata/configs"
	"github.com/amlun/housedata/internal/job"
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	s := gocron.NewScheduler(time.Local)
	ljJob := job.NewJob(configs.LJ)

	ijJob := job.NewJob(configs.IJ)

	if _, err := s.Every("1h").Do(ljJob.Do); err != nil {
		return
	}
	if _, err := s.Every("1h").Do(ijJob.Do); err != nil {
		return
	}

	log.Info("start jobs ...")
	s.StartAsync()

	sig := <-c
	log.WithField("signal", sig).Info("receive signal")
	log.Info("stop jobs ...")
	s.Stop()
}
