package LV2

import (
	"fmt"
	"sync"
)

// 关于自增
func Task(number *int, account int) int {
	var mu sync.Mutex
	w := sync.WaitGroup{}
	for range account {
		w.Add(1)
		go func() {
			mu.Lock()
			defer w.Done()
			*number++
			mu.Unlock()
		}()
	}
	w.Wait()
	return *number
}

type Workpool struct {
	Task func(*int, int) int
}

// 关于分发任务
func Delivery(account int, accept int, number int, times int) {
	ch := make(chan Workpool, account)
	w := sync.WaitGroup{}
	for x := 0; x < account; x++ {
		w.Add(1)
		go func(id int) {
			workspace := number
			defer w.Done()
			fmt.Printf("唤醒第%dwoker\n", id)
			for t := range ch {
				result := t.Task(&workspace, times)
				fmt.Printf("第%dwoker，结果%d\n", id, result)
			}
		}(x)
	}
	for x := 0; x < accept; x++ {
		fmt.Printf("第%d次分发\n", x)
		task := Workpool{
			Task: Task,
		}
		ch <- task

	}

	fmt.Println("关闭通道")
	w.Wait()
	close(ch)
	fmt.Println("完成")
}
