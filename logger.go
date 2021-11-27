package logger

import (
	"log"
	"os"
)

var err = log.New(os.Stdout, "error", log.LstdFlags|log.Lshortfile)
var info = log.New(os.Stdout, "info", log.LstdFlags|log.Lshortfile)
var debug = log.New(os.Stdout, "debug", log.LstdFlags|log.Lshortfile)

func Error(v ...interface{}) {
	err.Println(v...)
}

func Info(v ...interface{}) {
	info.Println(v...)
}

func DebugFatal(v ...interface{}) {
	debug.Fatal(v...)
}

func ErrorFatal(v ...interface{}) {
	err.Fatal(v...)
}

func InfoFatal(v ...interface{}) {
	info.Fatal(v...)
}

func Debug(v ...interface{}) {
	debug.Println(v...)
}
