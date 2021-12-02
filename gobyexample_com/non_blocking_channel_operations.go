package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
)

// Basic sends and receives on channels are blocking. However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
// 常规的通过通道发送和接收数据是阻塞的。 然而，我们可以使用带一个 default 子句的 select 来实现 非阻塞 的发送、接收，甚至是非阻塞的多路 select。

// nonBlockingChannelOperationsCmd represents the Non-Blocking Channel Operations command
var nonBlockingChannelOperationsCmd = &cobra.Command{
	Use:   "go-by-example:non_blocking_channel_operations",
	Short: "https://gobyexample.com/non-blocking-channel-operations",
	Run: func(cmd *cobra.Command, args []string) {

		messages := make(chan string)
		signals := make(chan bool)

		// Here’s a non-blocking receive. If a value is available on messages then select will take the <-messages case with that value. If not it will immediately take the default case.
		// 这是一个非阻塞接收的例子。 如果在 messages 中存在，然后 select 将这个值带入 <-messages case 中。 否则，就直接到 default 分支中。
		select {
		case msg := <-messages:
			fmt.Println("received message", msg)
		default:
			fmt.Println("no message received")
		}

		// A non-blocking send works similarly. Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver. Therefore the default case is selected.
		// 一个非阻塞发送的例子，代码结构和上面接收的类似。 msg 不能被发送到 message 通道，因为这是 个无缓冲区通道，并且也没有接收者，因此， default 会执行。

		msg := "hi"
		select {
		case messages <- msg:
			fmt.Println("sent message", msg)
		default:
			fmt.Println("no message sent")
		}

		// We can use multiple cases above the default clause to implement a multi-way non-blocking select. Here we attempt non-blocking receives on both messages and signals.
		// 我们可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。 这里我们试图在 messages 和 signals 上同时使用非阻塞的接收操作。

		select {
		case msg := <-messages:
			fmt.Println("received message", msg)
		case sig := <-signals:
			fmt.Println("received signal", sig)
		default:
			fmt.Println("no activity")
		}

	},
}

func init() {
	cmd.RootCmd.AddCommand(nonBlockingChannelOperationsCmd)
}
