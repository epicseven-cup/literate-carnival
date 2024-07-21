package logger

import (
	"log"
	"os"
)

const (
	FILENAME   = "literate-carnival.log"
	FLAGS      = os.O_RDWR | os.O_CREATE | os.O_APPEND
	PERMISSION = 0666
)

var DefaultLogger *log.Logger

func init() {
	DefaultLogger = log.Default()
	file, err := os.OpenFile(FILENAME, FLAGS, PERMISSION)
	if err != nil {
		DefaultLogger.Fatalln("Error when trying to open the file for logging: ", err)
		return
	}
	DefaultLogger.SetOutput(file)
	DefaultLogger.SetFlags(FLAGS)
}
