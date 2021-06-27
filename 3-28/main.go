package main

import "fmt"

func total(n int, c chan int) {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	c <- t
}

func main() {
	c := make(chan int)
	go total(100, c)
	// チャンネルから値を取得する場合、その値が送られるまで処理を待つ
	fmt.Println("total:", <-c)
}

