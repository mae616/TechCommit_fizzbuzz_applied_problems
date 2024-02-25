package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FizzBuzz(num int) string {
	switch {
	case num%3 == 0 || strings.HasSuffix(strconv.Itoa(num), "3"):
		return fmt.Sprintf("%d!", num)
	default:
		return fmt.Sprintf("%d", num)
	}
}

func main() {
	fmt.Println("FizzBuzz start")

	for i := 1; i <= 150; i++ {
		fmt.Println(FizzBuzz(i))
	}

	fmt.Println("FizzBuzz end!")
}
