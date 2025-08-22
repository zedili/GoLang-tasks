package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// 加10
func add10(num *int) {
	*num += 10
}

/**
defer ：延迟执行，在函数退出时执行
加 defer 和不加 defer 的区别
场景			使用 defer wg.Done()		不使用 defer (直接 wg.Done())
执行时机		函数退出时自动执行				立即执行
异常安全		即使发生 panic 也会执行			panic 时可能不会执行
多返回点		自动覆盖所有返回路径			需在每个返回点手动调用
代码简洁性		只需写一次，位置灵活			需在多个地方重复调用
资源泄露风险	基本不会泄露					可能因忘记调用导致泄露
适用场景		推荐在 goroutine 中使用			简单函数中可能直接调用


什么是 panic ？
panic （程序终止信号）：
1、表示程序遇到了无法恢复的致命错误
2、类似于其他语言的异常（如 Java 的 RuntimeException），但设计更简单

defer 的交互：
1、panic 后，当前函数的 defer 会照常执行
2、panic 向调用栈上层传播
**/
// Go 提供 recover 内置函数来捕获 panic
func safeCall() {
	// 显式触发 panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常：", err)
		}
	}()

	panic("出错了")
}

/**
defer wg.Done() 是 Go 并发编程中的最佳实践：
1、它确保 WaitGroup 计数器一定会被减少
2、它使代码更健壮，能处理 panic 和复杂控制流
3、它使代码更简洁，减少错误
**/
// goroutine 启动协程
func printNums() {
	// sync 包中的一个结构，用于在并发编程中等待一组 goroutine 完成执行
	var wg sync.WaitGroup
	// 增加或减少 WaitGroup 中的计数器。如果 delta 为正，它会增加计数器；如果为负，它会减少计数器。
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i += 2 {
			fmt.Println("奇数：", i)
			time.Sleep(100 * time.Millisecond)
		}
		// defer wg.Done()
	}()

	go func() {
		defer wg.Done()
		for j := 0; j < 10; j += 2 {
			fmt.Println("偶数：", j)
			time.Sleep(100 * time.Millisecond)
		}
		// defer wg.Done()
	}()

	wg.Wait()
	// time.Sleep(2 * time.Second)
}

/**

题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。,

**/

/*
*

Go 的接口机制采用独特的 “鸭子类型” 设计，核心特点如下：
1、隐式实现，不需要声明 implements
2、只要类型实现了接口的所有方法，就自动满足该接口

*
*/
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

/**

题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
**/

type Persion struct {
	Name string
	Age  int
}

type Employee struct {
	Persion
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println(e.Persion)
	fmt.Println(e.EmployeeID)
}

/*
*

题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。

*
*/
func channelCommuinication() {

	/*
	   *
	   make：内置函数，用于初始化通道、切片、map 等引用类型

	   	chan：关键字，表示创建通道
	   	int：指定通道传输的数据类型（可以是任何类型）

	   *
	*/
	var wg2 sync.WaitGroup
	wg2.Add(2)
	ch := make(chan int)  // 创建一个无缓冲 channel
	ch2 := make(chan int) // 创建一个无缓冲 channel
	ch3 := make(chan int) // 创建一个无缓冲 channel

	/**
	// 无缓冲 channel ：
	1、需要发送和接收同时准备好才会执行
	2、主协程退出会强制结束所有子协程
	3、waitGroup 是最规范的同步方式
	优点：
	1、确保发送和接收的同步
	2、避免复杂的同步问题
	**/

	// 产生数据的协程
	go func() {
		defer wg2.Done()
		for i := 1; i < 10; i++ {
			// 每个发送操作都会阻塞，直到有接收操作准备好
			ch <- i // 发送数据到 channel
			ch2 <- i + 10
			ch3 <- i + 20
		}
		close(ch) // 关闭 channel ，channel 关闭后，只能接收channel 消息，不能再发送消息到 channel
	}()

	// 消费数据的协程
	go func() {
		defer wg2.Done()
		// range 读取消息只适合单个 channel 的场景
		// for num := range ch3 {
		// 	fmt.Println("从 ch3 接收到的数据：:", num)
		// }
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					// channel 已关闭
					return
				}
				fmt.Println("从 ch 接收到的数据：", v)
			case v, ok := <-ch2:
				if !ok {
					// channel 已关闭
					return
				}
				fmt.Println("从 ch2 接收到的数据：", v)
			case v, ok := <-ch3:
				if !ok {
					// channel 已关闭
					return
				}
				fmt.Println("从 ch3 接收到的数据：", v)
			}
		}

	}()

	wg2.Wait()
}

/*
*

有缓存 channel ：
1、解耦生产者和消费者，允许两者以不同速度工作
2、提高吞吐里，减少同步等待时间
3、平滑流量波动，缓冲区吸收突发流量

使用原则：
1、默认优先使用无缓冲通道
2、当需要解耦生产消费速度时使用有缓冲通道
3、合理设置缓冲区大小（通过性能测试确定）
4、始终确保有接收方处理数据，避免内存泄漏

明确通道所有权和关闭责任
*
*/
func bufferedChannel() {
	// 创建缓冲区大小为 10 的channel
	bufChannel := make(chan int, 10)

	var bfWg sync.WaitGroup
	bfWg.Add(2)

	go func() {
		defer bfWg.Done()
		for i := 0; i < 10; i++ {
			bufChannel <- i
		}
		close(bufChannel)
	}()

	go func() {
		defer bfWg.Done()
		for num := range bufChannel {
			fmt.Println("bufChan:", num)
		}
	}()

	bfWg.Wait()
}

/**

题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。,

sync.Mutx 互斥锁：
特点：
1、通过加锁机制保证临界区安全
2、适合用于复杂操作（如多变量修改）
3、可读性好，逻辑清晰、

注意事项：
1、必须成对使用 Lock() 、 Unlock()
2、避免嵌套导致死锁
3、性能开销较大（上下文切换）


atomic 原子操作：
特点：
1、CPU 指令级原子性操作
2、性能极高（无锁设计）
3、仅支持基本数据类型（int32、64等）

注意事项：
1、只能用于简单单一操作
2、需要指针操作
3、无读锁机制（需原子加载）


使用过原则
简单单一操作 => 原子操作
复杂结构体修改 => 互斥锁
高频操作 => 优选原子操作

**/

var (
	counter         int
	mutex           sync.Mutex
	wgIncreMentMutx sync.WaitGroup
)

func incrementMutx() {
	defer wgIncreMentMutx.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		counter++
		fmt.Println("incrementMutx:", counter)
		mutex.Unlock()
	}
}

func startIncrementMutx() {
	wgIncreMentMutx.Add(10)
	for i := 0; i < 10; i++ {
		go incrementMutx()
	}
	wgIncreMentMutx.Wait()
}

/**
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
**/

var (
	atomicCounter int32
	atmcCtWg      sync.WaitGroup
)

func incrementAtomic() {
	atmcCtWg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(&atomicCounter, 1)
		fmt.Println("incrementAtomic:", atomicCounter)
	}
}

func startIncrementAtomic() {
	atmcCtWg.Add(10)
	for i := 0; i < 10; i++ {
		go incrementAtomic()
	}

	atmcCtWg.Wait()
}

func main() {
	// i := 10
	// add10(&i)
	// fmt.Println(i)

	// printNums()

	// safeCall()

	// rect := Rectangle{Width: 1, Height: 2}
	// fmt.Println("rect area:", rect.Area())
	// fmt.Println("rect perimeter:", rect.Perimeter())

	// cic := Circle{Radius: 3}
	// fmt.Println("cic area:", cic.Area())
	// fmt.Println("cic perimeter:", cic.Perimeter())

	// emp := Employee{
	// 	Persion:    Persion{Name: "my-name", Age: 18},
	// 	EmployeeID: "123",
	// }

	// emp.PrintInfo()

	// channelCommuinication()

	// bufferedChannel()

	startIncrementMutx()
	startIncrementAtomic()

}
