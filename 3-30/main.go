package main

import (
	"fmt"
	"time"
)

func total(c chan int) {
	n := <-c
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total:", t)
}

func main() {
	c := make(chan int)
	go total(c)
	// チャンネルを受け取れる準備ができてから代入（値がまだ用意されていない場合は、送られてくるまで処理を待つ）
	c <- 100
	time.Sleep(100 * time.Millisecond)
}
