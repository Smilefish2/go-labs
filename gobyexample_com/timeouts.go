package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"time"
)

// Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time. Implementing timeouts in Go is easy and elegant thanks to channels and select.
// 超时 对于一个需要连接外部资源， 或者有耗时较长的操作的程序而言是很重要的。 得益于通道和 select，在 Go 中实现超时操作是简洁而优雅的。

// timeoutsCmd represents the timeouts command
var timeoutsCmd = &cobra.Command{
	Use:   "go-by-example:timeouts",
	Short: "https://gobyexample.com/timeouts",
	Run: func(cmd *cobra.Command, args []string) {

		// For our example, suppose we’re executing an external call that returns its result on a channel c1 after 2s. Note that the channel is buffered, so the send in the goroutine is nonblocking. This is a common pattern to prevent goroutine leaks in case the channel is never read.
		// 在这个例子中，假如我们执行一个外部调用， 并在 2 秒后使用通道 c1 返回它的执行结果。

		c1 := make(chan string, 1)
		go func() {
			time.Sleep(2 * time.Second)
			c1 <- "result 1"
		}()

		// Here’s the select implementing a timeout. res := <-c1 awaits the result and <-time.After awaits a value to be sent after the timeout of 1s. Since select proceeds with the first receive that’s ready, we’ll take the timeout case if the operation takes more than the allowed 1s.
		// 这里是使用 select 实现一个超时操作。 res := <- c1 等待结果，<-Time.After 等待超时（1秒钟）以后发送的值。 由于 select 默认处理第一个已准备好的接收操作， 因此如果操作耗时超过了允许的 1 秒的话，将会执行超时 case。
		select {
		case res := <-c1:
			fmt.Println(res)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1")
		}

		// If we allow a longer timeout of 3s, then the receive from c2 will succeed and we’ll print the result.
		// 如果我们允许一个长一点的超时时间：3 秒， 就可以成功的从 c2 接收到值，并且打印出结果。
		c2 := make(chan string, 1)
		go func() {
			time.Sleep(2 * time.Second)
			c2 <- "result 2"
		}()
		select {
		case res := <-c2:
			fmt.Println(res)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout 2")
		}

		// Running this program shows the first operation timing out and the second succeeding.
		// 运行这个程序，首先显示运行超时的操作，然后是成功接收的。
	},
}

func init() {
	cmd.RootCmd.AddCommand(timeoutsCmd)
}
