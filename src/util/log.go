package util

import "fmt"

func Log(message string, v bool) {
	if v {
		fmt.Println(message)
	}
}
