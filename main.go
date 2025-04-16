package main

import (
	"fmt"
	"time"
)

func main() {
	chLetter := make(chan struct{})
	chNum := make(chan struct{})

	go func() {
		for i := 1; i <= 26; i++ {
			<-chLetter
			fmt.Print(string(rune(i-1+'A')), " ")
			chNum <- struct{}{}
		}
	}()

	go func() {
		for i := 1; i <= 26; i++ {
			<-chNum
			fmt.Print(i, " ")
			if i != 26 {
				time.Sleep(time.Second)
				chLetter <- struct{}{} // 上面打印完Z后结束了，继续传递信号会阻塞
			}
		}
	}()

	chLetter <- struct{}{}
	time.Sleep(26 * time.Second)

}
