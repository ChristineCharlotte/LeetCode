package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "123"
	aSlice := strings.Split(a, "")
	fmt.Println(aSlice)
}