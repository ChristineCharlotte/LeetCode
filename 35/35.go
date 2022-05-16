package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "  a  bbb cc "
	arr := strings.Split(str, " ")
	for _, val := range arr {
		fmt.Println("1"+val+"1")
	}
}