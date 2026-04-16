package helper

import "log"

type ErrorLogger func(err error, context string)

var currentErrorLogger ErrorLogger = defaultErrorLogger

func defaultErrorLogger(err error, context string) {
	log.Printf("%s: %v", context, err)
}

func SetErrorLogger(logger ErrorLogger) {
	if logger == nil {
		currentErrorLogger = defaultErrorLogger
		return
	}
	currentErrorLogger = logger
}

func LogError(err error, context string) {
	if err == nil {
		return
	}
	currentErrorLogger(err, context)
}
