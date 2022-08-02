package main

import "fmt"

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
}

func decodeString(s string) string {
	var decode func(start int) (string, int)
	decode = func(start int) (str string, end int) {
		num := 0
		for i := start; i < len(s); i++ {
			if isNumber(s[i]) {
				num = num*10 + int(s[i]-'0')
			} else if isLetter(s[i]) {
				str += string(s[i])
			} else if s[i] == '[' {
				item, index := decode(i + 1)
				for num != 0 {
					str += item
					num--
				}
				i = index
			} else if s[i] == ']' {
				end = i
				break
			}
		}
		return str, end
	}

	res, _ := decode(0)
	return res
}

func isLetter(u uint8) bool {
	return u >= 'A' && u <= 'Z' || u >= 'a' && u <= 'z'
}

func isNumber(u uint8) bool {
	return u >= '0' && u <= '9'
}

// https://leetcode.cn/problems/decode-string/solution/shuang-zhan-di-gui-goban-ben-by-liansy-3rlz/
