package gobyexample_com

import (
	"fmt"
	"go-labs/cmd"
	"time"

	"github.com/spf13/cobra"
)

// A goroutine is a lightweight thread of execution.
// 协程(goroutine) 是轻量级的执行线程。

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// goroutinesCmd represents the goroutines command
var goroutinesCmd = &cobra.Command{
	Use:   "go-by-example:goroutines",
	Short: "https://gobyexample.com/goroutines",
	Run: func(cmd *cobra.Command, args []string) {

		// Suppose we have a function call f(s). Here’s how we’d call that in the usual way, running it synchronously.
		// 假设我们有一个函数叫做 f(s)。 我们一般会这样 同步地 调用它
		f("direct")

		// To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
		// 使用 go f(s) 在一个协程中调用这个函数。 这个新的 Go 协程将会 并发地 执行这个函数。
		go f("goroutine")

		// You can also start a goroutine for an anonymous function call.
		// 你也可以为匿名函数启动一个协程。
		go func(msg string) {
			fmt.Println(msg)
		}("going")

		// Our two function calls are running asynchronously in separate goroutines now. Wait for them to finish (for a more robust approach, use a WaitGroup).
		// 现在两个协程在独立的协程中 异步地 运行， 然后等待两个协程完成（更好的方法是使用 WaitGroup）。
		time.Sleep(time.Second)
		fmt.Println("done")

		// When we run this program, we see the output of the blocking call first, then the output of the two goroutines. The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.
		// 当我们运行这个程序时，首先会看到阻塞式调用的输出，然后是两个协程的交替输出。 这种交替的情况表示 Go runtime 是以并发的方式运行协程的。

		// Next we’ll look at a complement to goroutines in concurrent Go programs: channels.
		// 接下来我们将学习协程的辅助特性：通道（channels）。
	},
}

func init() {
	cmd.RootCmd.AddCommand(goroutinesCmd)
}
