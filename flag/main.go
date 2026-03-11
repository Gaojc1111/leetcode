package main

import (
	"flag"
	"fmt"
)

// 1. flag参数定义
var (
	intFlag    = flag.Int("age", 18, "set your age")
	stringFlag = flag.String("name", "zhang3", "set your name")
)

func main() {
	// 2. 解析命令行输入
	flag.Parse()
	fmt.Printf("age: %d, name: %s\n", *intFlag, *stringFlag)

	// 3. go run .\main.go 不指定参数，使用默认参数
	// 4. go run .\main.go -age 23 -name Hbzhtd 指定参数
	// 5. go run .\main.go -h or -help  查看参数用法
}
