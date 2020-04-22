package cwf

import (
	"log"
)

const (
	black   = "\033[1;30m"
	red     = "\033[1;31m"
	green   = "\033[1;32m"
	yellow  = "\033[1;33m"
	blue    = "\033[1;34m"
	magenta = "\033[1;35m"
	cyan    = "\033[1;36m"
	white   = "\033[1;37m"
	reset   = "\033[0m"
)

type clogInterface interface {
	Debug(...interface{})
	Info(...interface{})
	Warning(...interface{})
	Error(...interface{})
	Fatal(...interface{})
}

type clog struct{}

func (l *clog) Debug(v ...interface{}) {
	log.Printf("%sDEBUG%s: %v\n", blue, reset, v)
}

func (l *clog) Info(v ...interface{}) {
	log.Printf("%sINFO%s: %v\n", green, reset, v)
}

func (l *clog) Warning(v ...interface{}) {
	log.Printf("%sWARN%s: %v\n", yellow, reset, v)
}

func (l *clog) Error(v ...interface{}) {
	log.Printf("%sERROR%s: %v\n", red, reset, v)
}

func (l *clog) Fatal(v ...interface{}) {
	log.Fatalf("\\e[05;31mFATAL\\e[0m: %v\n", v)
}

var logger = &clog{}
