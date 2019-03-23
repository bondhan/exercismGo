package secret

import (
	"math"
)

var consts = []string{"", "wink", "double blink", "close your eyes", "jump"}

func reverseSlice(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseSlice(input[1:]), input[0])
}

// func reverseSlice(input []string) []string {
// 	var reversedSlice []string

// 	len := len(input)
// 	for i := len - 1; i >= 0; i-- {
// 		reversedSlice = append(reversedSlice, input[i])
// 	}

// 	return reversedSlice
// }

func Handshake(code uint) []string {
	var compile []string

	for index := 0; index <= 4; index++ {
		switch value := code & uint(math.Pow(2, float64(index))); value {
		case uint(1):
			compile = append(compile, consts[1])
		case uint(2):
			compile = append(compile, consts[2])
		case uint(4):
			compile = append(compile, consts[3])
		case uint(8):
			compile = append(compile, consts[4])
		case uint(16):
			compile = reverseSlice(compile)
		}
	}

	return compile
}
