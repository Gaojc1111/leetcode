package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// 三个goroutine，每秒钟打印 cat、dog、fish
func TestCatDogFish(t *testing.T) {
	chCat := make(chan struct{})
	chDog := make(chan struct{})
	chFish := make(chan struct{})

	go func() {
		for {
			<-chCat
			fmt.Println("cat")
			time.Sleep(time.Second)
			chDog <- struct{}{}
		}
	}()

	go func() {
		for {
			<-chDog
			fmt.Println("dog")
			time.Sleep(time.Second)
			chFish <- struct{}{}
		}
	}()

	go func() {
		for {
			<-chFish
			fmt.Println("fish")
			time.Sleep(time.Second)
			chCat <- struct{}{}
		}
	}()

	chCat <- struct{}{}
	time.Sleep(15 * time.Second)
}

// 两个goroutine轮流输出 A 1 B 2 C 3 ... Z 26
// 在main函数中可以使用
func TestA1Z26(t *testing.T) {
	chLetter := make(chan struct{})
	chNum := make(chan struct{})

	go func() {
		for i := 1; i <= 26; i++ {
			<-chLetter
			fmt.Println(string(rune(i-1+'A')), " ")
			chNum <- struct{}{}
		}
	}()

	go func() {
		for i := 1; i <= 26; i++ {
			<-chNum
			fmt.Println(i, " ")
			if i != 26 {
				time.Sleep(time.Second)
				chLetter <- struct{}{} // 上面打印完Z后结束了，继续传递信号会阻塞
			}
		}
	}()

	chLetter <- struct{}{}
	time.Sleep(26 * time.Second)
}
