# go-example
example for go


### Goroutines
Go使用关键字go开启并发,当代码执行到go hello()是会创建一个Goruntines去执行hello()方法，随后立即执行下一句，main方法执行完毕后，无论goruntines是否执行完毕都会被终止执行,
因此运行下面的代码你会发现结果只输出了world

```go
package main

import "fmt"

func hello(){
    fmt.Println("hello")
}

func main(){
    go hello()
    fmt.Println("world")
}
```

对以上代码进行修改，让它可以顺利输出hello world
```go
package main

import "fmt"

func hello(flag *bool){
	fmt.Println("Hello")
	*flag = true
}

func main() {
	flag := false
	go hello(&flag)
	for !flag {

	}
	fmt.Println("World")
}
```
这里我通过向hello方法传递一个bool类型的标签flag，随后使用了一个for循环来判断flag的值，如果hello方法执行完毕，那么flag的值将被设置为true，跳出for循环，之后执行world的输出
运行代码你会发现这次输出了我们期待的结果。
在以上代码中我们实现了线程之间的通信，但是过于粗暴，go提供了chan来实现线程之间的通信
chan是一种带有类型的管道，它可以接受发送或者接送该类型的数据，我们可以吧数据的接受与发送类比为管道的流入与流出，例如:一个类型为int的chan，我们可以将数字1，2从管道的一段流入管道，
从管道的另一端流出刚刚我们流入的1，2，对应到go中的语法就是，流入： chan <- 1, 流出： <- chan
默认管道的放入与取出都会阻塞程序，当从一个空地管道中取出数据的时候，程序会阻塞，直到有数据放入管道中。反之当管道中存在数据却没有取出的时候程序就会发生死锁。
根据管道的特性将代码修改如下
```go
package main

import (
	"fmt"
)

func hello(flag chan bool){
	fmt.Println("Hello")
	flag <- true
	
}

func main() {
	flag := make(chan bool)
	go hello(flag)
	<- flag
	fmt.Println("World")
}
```
代码-3中，我通过make方法初始化了一个bool类型的管道flag，执行hello方法，随后从管道中取出数据，此时如果管道中没有数据程序会阻塞在这里，直到hello方法执行完毕，将true放入管道中
，运行代码程序输出hello world，你可能会注意到hello方法中放入了一个true，但是后续中并没有用到这个值，这里的true只是一个标识，表示管道中放入了一个数据，即使放入false程序也是可以正常执行的
，go允许取出管道中的值而不作处理。可以通过下面的方式进行处理
```
func main() {
	c := make(chan bool)
	go hello(c)
	success := <- c
	if success {
		fmt.Println("处理业务逻辑")
	}
	fmt.Println("World")
}
```

#### sync.Mutex

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func add(i *int) {
	mu.Lock()
	*i ++
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
```



