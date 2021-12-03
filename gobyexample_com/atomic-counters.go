package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"sync"
	"sync/atomic"
)

// The primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools. There are a few other options for managing state though. Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
// Go 中最主要的状态管理机制是依靠通道间的通信来完成的。 我们已经在工作池的例子中看到过，并且，还有一些其他的方法来管理状态。 这里，我们来看看如何使用 sync/atomic 包在多个协程中进行 _原子计数_。

// atomicCountersCmd represents the atomic counters command
var atomicCountersCmd = &cobra.Command{
	Use:   "go-by-example:atomic-counters",
	Short: "https://gobyexample.com/atomic-counters",
	Run: func(cmd *cobra.Command, args []string) {

		// We’ll use an unsigned integer to represent our (always-positive) counter.
		// 我们将使用一个无符号整型（永远是正整数）变量来表示这个计数器。
		var ops uint64

		// A WaitGroup will help us wait for all goroutines to finish their work.
		// WaitGroup 帮助我们等待所有协程完成它们的工作。
		var wg sync.WaitGroup

		// We’ll start 50 goroutines that each increment the counter exactly 1000 times.
		// 我们会启动 50 个协程，并且每个协程会将计数器递增 1000 次。
		for i := 0; i < 50; i++ {
			wg.Add(1)

			// To atomically increment the counter we use AddUint64, giving it the memory address of our ops counter with the & syntax.
			// 使用 AddUint64 来让计数器自动增加， 使用 & 语法给定 ops 的内存地址。
			go func() {
				for c := 0; c < 1000; c++ {

					atomic.AddUint64(&ops, 1)
				}
				wg.Done()
			}()
		}

		// Wait until all the goroutines are done.
		// 等待，直到所有协程完成。
		wg.Wait()

		fmt.Println("ops:", ops)

		// It’s safe to access ops now because we know no other goroutine is writing to it. Reading atomics safely while they are being updated is also possible, using functions like atomic.LoadUint64.
		// 现在可以安全的访问 ops，因为我们知道，此时没有协程写入 ops， 此外，还可以使用 atomic.LoadUint64 之类的函数，在原子更新的同时安全地读取它们。

		// We expect to get exactly 50,000 operations. Had we used the non-atomic ops++ to increment the counter, we’d likely get a different number, changing between runs, because the goroutines would interfere with each other. Moreover, we’d get data race failures when running with the -race flag.
		// 预计会进行 50,000 次操作。如果我们使用非原子的 ops++ 来增加计数器， 由于多个协程会互相干扰，运行时值会改变，可能会导致我们得到一个不同的数字。 此外，运行程序时带上 -race 标志，我们可以获取数据竞争失败的详情。

		// Next we’ll look at mutexes, another tool for managing state.
		// 接下来，我们看一下管理状态的另一个工具——互斥锁。
	},
}

func init() {
	cmd.RootCmd.AddCommand(atomicCountersCmd)
}
