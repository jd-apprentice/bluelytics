package utils

import (
	"fmt"
	"log"
)

const (
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func ErrorMessageFmt(errorMessage string, errorCause error) error {
	return fmt.Errorf(Red+errorMessage+": "+Yellow+"%v"+Reset, errorCause)
}

func LogFatalFmt(errorMessage string, errorCause error) {
	log.Fatalf(Red+errorMessage+": "+Yellow+"%v"+Reset, errorCause)
}
