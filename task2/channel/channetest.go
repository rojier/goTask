package channeltest

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

/*
管道是安全的 写入比读取慢 也不会有问题
读取会自动等待
*/
func fn1(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i + 1
		fmt.Println("writeData写入数据:", i+1)
		time.Sleep(time.Millisecond * 100)
	}
	close(intChan)
	wg.Done()
}
func fn2(intChan chan int) {
	for v := range intChan {
		fmt.Println("readData 读到数据=", v)
	}
	wg.Done()
}
func TestChannel() {
	ch1 := make(chan int, 100)
	wg.Add(1)
	go fn1(ch1)
	wg.Add(1)
	go fn2(ch1)
	wg.Wait()
	fmt.Println("读取完毕")
}

var count = 0

// 读写锁
var mutex sync.RWMutex

func TestLock() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println("读写执行完毕")
}
func write() {
	mutex.Lock()
	count++
	fmt.Println("执行写操作:", count)
	mutex.Unlock()
	wg.Done()
}

func read() {
	//读的时候并行去操作
	mutex.RLock()
	fmt.Println("-执行读操作: ", count)
	time.Sleep(time.Second * 3)
	mutex.RUnlock()
	wg.Done()
}
