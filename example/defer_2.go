package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	// defer语句中嵌套的函数签名还是按照代码堆栈顺序执行
	// defer语句按照逆序顺序执行，执行时间在return之前，ret指令之后
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
