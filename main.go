package main

import (
	"fmt"
	"sync"
)

type student struct {
	Name string
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(student)
	},
}

func main() {
	stu := studentPool.Get().(*student)
	fmt.Println("第一次Get操作：", stu)

	stu.Name = "zhang3"
	fmt.Println("设置 stu.Name =", stu.Name)

	studentPool.Put(stu) // 放回，实际使用时，记得进行reset

	stu = studentPool.Get().(*student)
	fmt.Println("放回后Get数据:", stu)

	stu1 := studentPool.Get().(*student)
	fmt.Println("不放回Get数据:", stu1)
}
