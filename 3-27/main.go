package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	message := "start"
	parameterMessage := func(nm string, n int) {
		fmt.Println(nm, message)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
	hello := func(n int) {
		const nm string = "hello"
		for i := 0; i < 10; i++ {
			message += " h" + strconv.Itoa(i)
			parameterMessage(nm, n)
		}
	}
	main := func(n int) {
		const nm string = "*main"
		for i := 0; i < 5; i++ {
			message += " m" + strconv.Itoa(i)
			parameterMessage(nm, 100)
		}
	}
	go hello(60)
	main(100)
}
