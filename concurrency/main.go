package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	time.Sleep(time.Second * 1)
	var input string
	fmt.Scanln(&input)
}
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n,": ", i)
	}
}
