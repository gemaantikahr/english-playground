package config

import "os"

func FilePath() string {
	return os.Getenv("FILE_PATH")
}
