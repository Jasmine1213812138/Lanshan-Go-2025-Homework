package main

import (
	"errors"
	"fmt"
	"os"
)

type operate func(a, b int) int
type cal func(int, int) (int, error)

func calculate(op operate) cal {
	return func(a, b int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(a, b), nil
	}

}
func add(a, b int) int {
	return a + b
}
func subtract(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	if b == 0 {
		//panic("division by zero")
		fmt.Println("零不能做除数")
		os.Exit(1)
	}
	return a / b
}
func main() {
	var a, b int
	var cal string
	var sign byte
	for {
		fmt.Println("欢迎使用！输入-1退出")
		fmt.Println("请输入第一个数字")
		_, err0 := fmt.Scan(&a)
		if a == -1 || err0 != nil {
			os.Exit(0)
		}
		fmt.Println("请输入第二个数字")
		_, err1 := fmt.Scan(&b)
		if b == -1 || err1 != nil {
			os.Exit(0)
		}
		fmt.Println("请选择你想进行的运算（+，-，*，/）")
		fmt.Scan(&cal)
		if cal == "-1" {
			os.Exit(0)
		}
		sign = cal[0]
		addcal := calculate(add)
		subcal := calculate(subtract)
		mulcal := calculate(mul)
		divcal := calculate(div)
		switch sign {
		case '+':
			result, _ := addcal(a, b)
			fmt.Println(result)
		case '-':
			result, _ := subcal(a, b)
			fmt.Println(result)
		case '*':
			result, _ := mulcal(a, b)
			fmt.Println(result)
		case '/':
			result, _ := divcal(a, b)
			fmt.Println(result)
		default:
			fmt.Println("输入出问题啦")
		}
		fmt.Println("不想算了就按-1")
		var judge string
		fmt.Scan(&judge)
		if judge == "-1" {
			os.Exit(0)
		}
	}
}
