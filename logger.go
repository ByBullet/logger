package logger

import (
	"log"
	"os"
)

var err = log.New(os.Stdout, "error", log.LstdFlags|log.Lshortfile)
var info = log.New(os.Stdout, "info", log.LstdFlags|log.Lshortfile)
var debug = log.New(os.Stdout, "debug", log.LstdFlags|log.Lshortfile)

func Error(s interface{}, v ...interface{}) {
	err.Println(s, v)
}

func Info(s interface{}, v ...interface{}) {
	info.Println(s, v)
}

func DebugFatal(s interface{}, v ...interface{}) {
	debug.Fatal(s, v)
}

func ErrorFatal(s interface{}, v ...interface{}) {
	err.Fatal(s, v)
}

func InfoFatal(s interface{}, v ...interface{}) {
	info.Fatal(s, v)
}

func Debug(s interface{}, v ...interface{}) {
	debug.Println(s, v)
}
