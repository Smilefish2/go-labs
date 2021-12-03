package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"math/rand"
	"sync/atomic"
	"time"
)

// In the previous example we used explicit locking with mutexes to synchronize access to shared state across multiple goroutines. Another option is to use the built-in synchronization features of goroutines and channels to achieve the same result. This channel-based approach aligns with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.
// 在前面的例子中，我们用 互斥锁 进行了明确的锁定， 来让共享的 state 跨多个 Go 协程同步访问。 另一个选择是，使用内建协程和通道的同步特性来达到同样的效果。 Go 共享内存的思想是，通过通信使每个数据仅被单个协程所拥有，即通过通信实现共享内存。 基于通道的方法与该思想完全一致！

// In this example our state will be owned by a single goroutine. This will guarantee that the data is never corrupted with concurrent access. In order to read or write that state, other goroutines will send messages to the owning goroutine and receive corresponding replies. These readOp and writeOp structs encapsulate those requests and a way for the owning goroutine to respond.
// 在这个例子中，state 将被一个单独的协程拥有。 这能保证数据在并行读取时不会混乱。 为了对 state 进行读取或者写入， 其它的协程将发送一条数据到目前拥有数据的协程中， 然后等待接收对应的回复。 结构体 readOp 和 writeOp 封装了这些请求，并提供了响应协程的方法。

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

// statefulGoroutinesCmd represents the stateful goroutines command
var statefulGoroutinesCmd = &cobra.Command{
	Use:   "go-by-example:stateful-goroutines",
	Short: "https://gobyexample.com/stateful-goroutines",
	Run: func(cmd *cobra.Command, args []string) {

		// As before we’ll count how many operations we perform.
		// 和前面的例子一样，我们会计算操作执行的次数。
		var readOps uint64
		var writeOps uint64

		// The reads and writes channels will be used by other goroutines to issue read and write requests, respectively.
		// 其他协程将通过 reads 和 writes 通道来发布 读 和 写 请求。
		reads := make(chan readOp)
		writes := make(chan writeOp)

		// Here is the goroutine that owns the state, which is a map as in the previous example but now private to the stateful goroutine. This goroutine repeatedly selects on the reads and writes channels, responding to requests as they arrive. A response is executed by first performing the requested operation and then sending a value on the response channel resp to indicate success (and the desired value in the case of reads).
		// 这就是拥有 state 的那个协程， 和前面例子中的 map 一样，不过这里的 state 是被这个状态协程私有的。 这个协程不断地在 reads 和 writes 通道上进行选择，并在请求到达时做出响应。 首先，执行请求的操作；然后，执行响应，在响应通道 resp 上发送一个值，表明请求成功（reads 的值则为 state 对应的值）。
		go func() {
			var state = make(map[int]int)
			for {
				select {
				case read := <-reads:
					read.resp <- state[read.key]
				case write := <-writes:
					state[write.key] = write.val
					write.resp <- true
				}
			}
		}()

		// This starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel. Each read requires constructing a readOp, sending it over the reads channel, and the receiving the result over the provided resp channel.
		// 启动 100 个协程通过 reads 通道向拥有 state 的协程发起读取请求。 每个读取请求需要构造一个 readOp，发送它到 reads 通道中， 并通过给定的 resp 通道接收结果。

		for r := 0; r < 100; r++ {
			go func() {
				for {
					read := readOp{
						key:  rand.Intn(5),
						resp: make(chan int)}
					reads <- read
					<-read.resp
					atomic.AddUint64(&readOps, 1)
					time.Sleep(time.Millisecond)
				}
			}()
		}

		// We start 10 writes as well, using a similar approach.
		// 用相同的方法启动 10 个写操作。
		for w := 0; w < 10; w++ {
			go func() {
				for {
					write := writeOp{
						key:  rand.Intn(5),
						val:  rand.Intn(100),
						resp: make(chan bool)}
					writes <- write
					<-write.resp
					atomic.AddUint64(&writeOps, 1)
					time.Sleep(time.Millisecond)
				}
			}()
		}

		// Let the goroutines work for a second.
		// 让协程们跑 1s。
		time.Sleep(time.Second)

		// Finally, capture and report the op counts.
		// 最后，获取并报告 ops 值。
		readOpsFinal := atomic.LoadUint64(&readOps)
		fmt.Println("readOps:", readOpsFinal)
		writeOpsFinal := atomic.LoadUint64(&writeOps)
		fmt.Println("writeOps:", writeOpsFinal)

		// Running our program shows that the goroutine-based state management example completes about 80,000 total operations.
		// 运行这个程序显示这个基于协程的状态管理的例子 达到了每秒大约 80,000 次操作。

		// For this particular case the goroutine-based approach was a bit more involved than the mutex-based one. It might be useful in certain cases though, for example where you have other channels involved or when managing multiple such mutexes would be error-prone. You should use whichever approach feels most natural, especially with respect to understanding the correctness of your program.
		// 通过这个例子我们可以看到，基于协程的方法比基于互斥锁的方法要复杂得多。 但是，在某些情况下它可能很有用， 例如，当你涉及其他通道，或者管理多个同类互斥锁时，会很容易出错。 您应该使用最自然的方法，尤其是在理解程序正确性方面。
	},
}

func init() {
	cmd.RootCmd.AddCommand(statefulGoroutinesCmd)
}
