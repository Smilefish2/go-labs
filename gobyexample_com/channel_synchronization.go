package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"
	"time"

	"github.com/spf13/cobra"
)

// We can use channels to synchronize execution across goroutines. Here’s an example of using a blocking receive to wait for a goroutine to finish. When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
// 我们可以使用通道来同步协程之间的执行状态。 这儿有一个例子，使用阻塞接收的方式，实现了等待另一个协程完成。 如果需要等待多个协程，WaitGroup 是一个更好的选择。

// This is the function we’ll run in a goroutine. The done channel will be used to notify another goroutine that this function’s work is done.
// 我们将要在协程中运行这个函数。 done 通道将被用于通知其他协程这个函数已经完成工作。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we’re done.
	// 发送一个值来通知我们已经完工啦。
	done <- true
}

// channelSynchronizationCmd represents the Channel Synchronization command
var channelSynchronizationCmd = &cobra.Command{
	Use:   "go-by-example:channel_synchronization",
	Short: "https://gobyexample.com/channel-synchronization",
	Run: func(cmd *cobra.Command, args []string) {

		// Start a worker goroutine, giving it the channel to notify on.
		// 运行一个 worker 协程，并给予用于通知的通道。
		done := make(chan bool, 1)

		// Block until we receive a notification from the worker on the channel.
		// 程序将一直阻塞，直至收到 worker 使用通道发送的通知。
		go worker(done)

		// If you removed the <- done line from this program, the program would exit before the worker even started.
		// 如果你把 <- done 这行代码从程序中移除， 程序甚至可能在 worker 开始运行前就结束了。
		<-done
	},
}

func init() {
	cmd.RootCmd.AddCommand(channelSynchronizationCmd)
}
