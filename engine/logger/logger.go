package logger

import "fmt"

func Println(obj ...interface{}) {
	fmt.Println(obj...)
}

func Print(obj ...interface{}) {
	fmt.Print(obj...)
}

func Printf(format string, obj ...interface{}) {
	fmt.Printf(format, obj...)
}