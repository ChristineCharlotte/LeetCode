package main

import "fmt"

func climbStairs(n int) int {
	s := []int{0, 1, 2}
	for i := 3; i <= n; i++ {
		s = append(s, s[i-1]+s[i-2])
	}
	return s[n]
}

func main() {
	re := climbStairs(3)
	fmt.Println(re)
}
