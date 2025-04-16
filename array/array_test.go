package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	doappend := func(s []int) {
		s = append(s, 1)
		printLenAndCap(s)
	}

	s := make([]int, 8, 8)
	doappend(s[:4]) // 长度取决于截取长度，容量取决于起始位置到原切片末尾
	printLenAndCap(s)

	doappend(s)
	printLenAndCap(s)
}

func printLenAndCap(s []int) {
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
}
