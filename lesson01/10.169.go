package main

import "fmt"

// LVX的函数
func average(x, y int) float64 {
	z := float64(x) / float64(y)
	return z
}

func main() {
	//LV0
	fmt.Println("王淑兰")
	//LV1
	const pi = 3.14
	r := 5
	area := pi * float64(r) * float64(r)
	fmt.Println(area)
	//LV2
	var sum int
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
	//LVX

	var j = 1
	var add, count = 0, 0
	for {
		fmt.Print("请输入一个整数（输入0结束）")
		_, _ = fmt.Scanf("%d", &j)
		if j == 0 {
			break
		}
		add += j
		count++
	}
	result := average(add, count)
	if result >= 60 {
		fmt.Printf("平均成绩为%.2f，成绩合格\n", result)
	} else {
		fmt.Printf("平均成绩为%.2f，成绩不合格\n", result)
	}
}
