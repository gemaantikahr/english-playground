package playground

import (
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestError(t *testing.T) {

	/*
		LOGRUS INIT
	*/
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	log.Error("Something failed but I'm not quitting.")
}
