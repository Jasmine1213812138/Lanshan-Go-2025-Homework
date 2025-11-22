package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type SearchTask struct {
	FilePath string
	Keyword  string
}
type SearchResult struct {
	FilePath string
}

// 此为主函数
func main() {
	var directory string
	var keyword string
	_, err := fmt.Scanf("%s,%s", &directory, &keyword)
	if err != nil {
		fmt.Println(err)
	}
	runSearch(directory, keyword)
}

// 把所有任务都用一个函数来解决
func runSearch(directory string, keyword string) {
	filePaths := getAllFile(directory)
	taskChan := make(chan SearchTask, 10)
	resultChan := make(chan SearchResult, 20)
	pool := createWorkerPool(5, resultChan, taskChan)
	dispatchTask(filePaths, keyword, taskChan)
	pool.Wait()
	close(resultChan)
	collectResult(resultChan)

}

// 获取所有路径
func getAllFile(dir string) []string {
	var filePaths []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		filePaths = append(filePaths, path)
		return nil
	})
	return filePaths
}

// 创建协程池并等待任务
func createWorkerPool(workerCount int, resultCh chan SearchResult, ch chan SearchTask) *sync.WaitGroup {
	wait := &sync.WaitGroup{}
	wait.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func(id int) {
			defer wait.Done()
			fmt.Printf("创造第%d个worker", id)
			for do := range ch {
				searchInFile(do, resultCh)
			}
		}(i)
	}
	return wait
}

// 分发任务
func dispatchTask(filePaths []string, keyword string, taskChan chan SearchTask) {
	for _, filePath := range filePaths {
		searchTask := SearchTask{filePath, keyword}
		taskChan <- searchTask
	}
	close(taskChan)
}

// 需要执行的任务
func searchInFile(task SearchTask, resultChan chan SearchResult) {
	if strings.Contains(task.FilePath, task.Keyword) {
		result := SearchResult{
			FilePath: task.FilePath,
		}
		resultChan <- result
	}
}

// 收集结果并格式化输出
func collectResult(resultChan chan SearchResult) {
	counter := 0
	for result := range resultChan {
		counter++
		fmt.Printf("文件路径：=%s", result.FilePath)
	}
	fmt.Printf("共找到%d个文件", counter)
}
