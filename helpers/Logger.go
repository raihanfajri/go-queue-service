package helpers

import (
	"log"
	"os"
	"time"
)

func LogError(msg string, errMsg error) {
	pathLog := "../logs"

	_ = os.Mkdir(pathLog, 0700)

	filename := pathLog + "/debug-" + getTimeNow() + ".log"

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Println(msg)
	log.Println(errMsg)
}

func getTimeNow() string {
	t := time.Now().Local()
	return t.Format("2006-01-02")
}
