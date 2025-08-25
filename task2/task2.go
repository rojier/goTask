package main

import (
	"fmt"
	"sync"
	"time"

	channeltest "github.com/learn/learn02_pk/GoWorkSpace/goTask/task2/channel"
	interface_test "github.com/learn/learn02_pk/GoWorkSpace/goTask/task2/interface"
)

/*
编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
*/
func pointAdd(num *int) {
	*num = *num + 1
}
func sliceX(a []int) {
	for k, _ := range a {
		a[k] = a[k] * 2
	}
}

var wg sync.WaitGroup //1、定义全局的 WaitGroup
func pirntJS() {
	for i := 0; i < 11; i++ {
		if i%2 == 1 {
			fmt.Println("奇数: ", i)
		}
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() // 4、goroutine 结束就登记-1
}
func pirntOS() {
	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			fmt.Println("偶数: ", i)
		}
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() // 4、goroutine 结束就登记-1
}

/*
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
*/
func testJO() {
	wg.Add(1) //、启动一个 goroutine 就登记+1
	go pirntJS()
	wg.Add(1)
	go pirntOS()
	wg.Wait()
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
func fn1(n int) {
	for num := (n-1)*300 + 1; num <= n*300; num++ {
		flag := true //假设是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(num)
		}
	}
	wg.Done()
}

func tj() {
	start := time.Now().Unix()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go fn1(i)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println("使用异步方法统计素数的耗时: ", end-start)

}
func main() {
	a := 1
	pointAdd(&a)
	fmt.Println(a)

	b := []int{11, 22, 33}
	sliceX(b)
	fmt.Println(b)
	fmt.Println("*****并发判断奇偶******")
	testJO()
	tj()
	fmt.Println("*****接口多态测试******")
	interface_test.TestInterface()
	fmt.Println("*****测试修改结构体字段******")
	interface_test.TestStruct()
	fmt.Println("*****测试管道协程写入和读取******")
	channeltest.TestChannel()
	fmt.Println("***********测试锁机制*****************")
	channeltest.TestLock()
}
