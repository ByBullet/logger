package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	ErrOut      = "ErrOut"
	WarnOut     = "WarnOut"
	InfoOut     = "InfoOut"
	DebugOut    = "DebugOut"
	ErrOutFile  = "ErrOutFile"
	WarnOutFile = "WarnOutFile"
	InfoOutFile = "InfoOutFile"
	DebugOutFie = "DebugOutFie"
)

var err = log.New(os.Stdout, "error: ", log.LstdFlags|log.Lshortfile)
var warn = log.New(os.Stdout, "warn: ", log.LstdFlags|log.Lshortfile)
var info = log.New(os.Stdout, "info: ", log.LstdFlags|log.Lshortfile)
var debug = log.New(os.Stdout, "debug: ", log.LstdFlags|log.Lshortfile)

func init() {
	file, err := os.Open("logger.json")
	c := make(map[string]string)
	if err != nil {
		warn.Println(err)

	} else {
		defer file.Close()
		json.NewDecoder(file).Decode(c)
	}

}

func Error(v ...interface{}) {
	err.Output(2, fmt.Sprintln(v...))
}

func Warn(v ...interface{}) {
	warn.Output(2, fmt.Sprintln(v...))
}

func Info(v ...interface{}) {
	info.Output(2, fmt.Sprintln(v...))
}

func DebugFatal(v ...interface{}) {
	debug.Output(2, fmt.Sprintln(v...))

}

func ErrorFatal(v ...interface{}) {
	err.Fatal(v...)
}

func WarnFatal(v ...interface{}) {
	warn.Fatalln(v...)
}

func InfoFatal(v ...interface{}) {
	info.Fatal(v...)
}

func Debug(v ...interface{}) {
	debug.Println(v...)
}
