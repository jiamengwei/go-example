package main

import (
	"fmt"
	"sync"
	"time"
)

/**
go 中的并发叫做gorouting，可以看作是轻量级的线程，
与线程相比成本小，只占用几kb的占空间，而且可以根据程序需要动态增加或缩小，而线程的栈大小是固定的
gorouting是多路复用少量线程，一个线程可以上千个Gorouting，当一个Gorouting被阻塞，其它Goroutinr立即启动
，移动到新的thread，

Goroutines使用通道进行通信，通信被设计成在访问共享内存时不发生程竞态条件，

使用make方法构造的Goroutines默认buffer容量为0，写入一个数据后如果没有并发Gorutines读出会造成死锁
cap表示总容量
len表示buffer已经使用的容量
*/

var mu sync.Mutex

func add(i *int) {
	mu.Lock()
	*i++
	mu.Unlock()
}

func main() {
	counter := 0
	for i := 0; i < 1000; i++ {
		go add(&counter)
	}
	time.Sleep(time.Second * 2)
	fmt.Println(counter)
}
