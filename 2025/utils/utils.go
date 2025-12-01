package utils

import "fmt"

func Debug(debug bool, args ...any) {
	if debug {
		fmt.Println(args...)
	}
}
