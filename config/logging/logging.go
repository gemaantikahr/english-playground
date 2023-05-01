package logging

import (
	config2 "english_playground/config"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitializeLog() *zap.Logger {

	t := time.Now()
	filename := fmt.Sprintf("%d-%02d-%02d",
		t.Year(), t.Month(), t.Day()) + ".log"

	path := config2.FilePath()
	//Create output path
	outPath := filepath.Join(path, "/storage/logs/others")

	fmt.Println("file", outPath)
	// Create dir output using above code
	if err := os.MkdirAll(outPath, os.ModeSticky|os.ModePerm); err != nil {
		panic("cannot create folder")
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile(outPath+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger
}
