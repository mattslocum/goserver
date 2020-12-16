package logger

import (
	"io/ioutil"
	"log"
	"os"
)


var Debug = log.New(os.Stdout, "(DEBUG) ", log.Ldate|log.Ltime|log.Lmsgprefix)
var Info  = log.New(os.Stdout, "(INFO) ",  log.Ldate|log.Ltime|log.Lmsgprefix)
var Warn  = log.New(os.Stdout, "(WARN) ",  log.Ldate|log.Ltime|log.Lmsgprefix|log.Lshortfile)
var Error = log.New(os.Stderr, "(ERROR) ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Lshortfile)


func SetupLogs() {
	switch os.Getenv("LOG_LEVEL") {
	case "INFO":
		Debug.SetOutput(ioutil.Discard)
	case "WARN":
		Debug.SetOutput(ioutil.Discard)
		Info.SetOutput(ioutil.Discard)
	case "ERROR":
		Debug.SetOutput(ioutil.Discard)
		Info.SetOutput(ioutil.Discard)
		Warn.SetOutput(ioutil.Discard)
	}
}
