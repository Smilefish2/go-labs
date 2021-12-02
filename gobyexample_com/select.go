package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"
	"time"

	"github.com/spf13/cobra"
)

// Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.
// Go 的 选择器（select） 让你可以同时等待多个通道操作。 将协程、通道和选择器结合，是 Go 的一个强大特性。

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "go-by-example:select",
	Short: "https://gobyexample.com/select",
	Run: func(cmd *cobra.Command, args []string) {

		// For our example we’ll select across two channels.
		// 在这个例子中，我们将从两个通道中选择。
		c1 := make(chan string)
		c2 := make(chan string)

		// Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.
		// 各个通道将在一定时间后接收一个值， 通过这种方式来模拟并行的协程执行（例如，RPC 操作）时造成的阻塞（耗时）。
		go func() {
			time.Sleep(1 * time.Second)
			c1 <- "one"
		}()
		go func() {
			time.Sleep(2 * time.Second)
			c2 <- "two"
		}()

		// We’ll use select to await both of these values simultaneously, printing each one as it arrives.
		// 我们使用 select 关键字来同时等待这两个值， 并打印各自接收到的值。
		for i := 0; i < 2; i++ {
			select {
			case msg1 := <-c1:
				fmt.Println("received", msg1)
			case msg2 := <-c2:
				fmt.Println("received", msg2)
			}
		}

		// We receive the values "one" and then "two" as expected.
		// 跟预期的一样，我们首先接收到值 "one"，然后是 "two"。

		// Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.
		// 注意，程序总共仅运行了两秒左右。因为 1 秒 和 2 秒的 Sleeps 是并发执行的，
	},
}

func init() {
	cmd.RootCmd.AddCommand(selectCmd)
}
