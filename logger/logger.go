package logger

import (
	"github.com/sirupsen/logrus"

	"github.com/HackGT/SponsorshipPortal/config"
)

func New(config *config.Config) *logrus.Logger {
	log := logrus.New()
	if config.Prod {
		log.SetLevel(logrus.InfoLevel)
		log.Formatter = &logrus.JSONFormatter{}
	} else {
		log.SetLevel(logrus.DebugLevel)
		log.Formatter = &logrus.TextFormatter{}
	}

	return log
}
