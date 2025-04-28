package logger

import "fmt"

var isPrint bool
var isDebug bool

func Init(print bool, debug bool) {
	isPrint = print
	isDebug = debug
}

func Println(v ...interface{}) {
	if isPrint {
		fmt.Println(v...)
	}
}

func Info(format string, v ...interface{}) {
	if isPrint {
		fmt.Printf(fmt.Sprintf("%s\n", format), v...)
	}
}

func Debug(format string, v ...interface{}) {
	if isDebug {
		fmt.Printf(fmt.Sprintf("%s\n", format), v...)
	}
}
