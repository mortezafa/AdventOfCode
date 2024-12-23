package utils

import "log"

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func AbsDiffInt(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
