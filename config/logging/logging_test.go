package logging

import (
	"fmt"
	"github.com/joho/godotenv"
	"testing"
)

func TestLog(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println(err)
		return
	}

	logger := InitializeLog()
	logger.Info("initializing log")
}
