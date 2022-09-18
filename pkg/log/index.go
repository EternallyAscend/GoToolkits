package log

import "fmt"

func ConsoleLog(level uint, contents ...interface{}) {
	fmt.Println(contents)
}

func FileLog(level uint, contents ...interface{}) {
}
