package main

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[string]
	val  T
}

// Number 限制只允许数值类型
// type Number interface {
// 	~int | ~int64 | ~float32 | ~float64
// }

// Sum 计算数值切片之和
func Sum[T Number](numbers []T) T {
	var total T
	for _, n := range numbers {
		total += n
	}
	return total
}

// 自定义底层为 int 的类型
type MyInt int

// 即使使用了自定义类型 MyInt，因为使用了 ~int，也可以正常运行
// var customNums []MyInt = []MyInt{1, 2, 3}
// fmt.Println(Sum(customNums)) // 6
