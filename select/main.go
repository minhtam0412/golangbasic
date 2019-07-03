package main

import (
	"fmt"
	"time"
)

func main() {
	//khai báo 2 channel
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			//gửi dữ liệu vào channel 1
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			//gủi dữ liêu vào channel 2
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			//tuỳ ý xử lý tuỳ theo channel
			select {
			//xử lý channel 1
			case msg1 := <-c1:
				fmt.Println(msg1)
				//xử lý channel 2
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
