package util

import (
	"log"
)

func SetLogTime(showTime bool) {
	if showTime {
		log.SetFlags(log.Ldate | log.Ltime)
	} else {
		log.SetFlags(0) // 不显示时间
	}
}

func PrintLog(message string, showTime bool) {
	SetLogTime(showTime)
	log.Println(message)
}
