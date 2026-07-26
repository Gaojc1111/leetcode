package main

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	var l List[int] = List[int]{
		next: &List[string]{
			next: nil,
			val:  "zhang3",
		},
		val: 18,
	}
	fmt.Printf("%v", l)
}
