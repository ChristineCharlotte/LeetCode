package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(decodeString("abc3[cd]xyz"))
}

func decodeString(s string) string {
	num, numStack := 0, make([]int, 0)
	str, strStack := "", make([]string, 0)

	for i := 0; i < len(s); i++ {
		target := s[i]
		if isNumber(target) {
			num = num*10 + int(target-'0')
		} else if isLetter(target) {
			str += string(target)
		} else if target == '[' {
			numStack = append(numStack, num)
			strStack = append(strStack, str)
			num, str = 0, ""
		} else if target == ']' {
			repeatTimes, item := numStack[len(numStack)-1], strStack[len(strStack)-1]
			numStack, strStack = numStack[:len(numStack)-1], strStack[:len(strStack)-1]
			str = item + strings.Repeat(str, repeatTimes)
		}
	}
	return str
}

func isLetter(u uint8) bool {
	return u >= 'A' && u <= 'Z' || u >= 'a' && u <= 'z'
}

func isNumber(u uint8) bool {
	return u >= '0' && u <= '9'
}
