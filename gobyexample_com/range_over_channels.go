package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
)

// In a previous example we saw how for and range provide iteration over basic data structures. We can also use this syntax to iterate over values received from a channel.
// 在前面的例子中， 我们讲过 for 和 range 为基本的数据结构提供了迭代的功能。 我们也可以使用这个语法来遍历的从通道中取值。

// rangeOverChannelsCmd represents the range over channels command
var rangeOverChannelsCmd = &cobra.Command{
	Use:   "go-by-example:range_over_channels",
	Short: "https://gobyexample.com/range-over-channels",
	Run: func(cmd *cobra.Command, args []string) {

		// We’ll iterate over 2 values in the queue channel.
		// 我们将遍历在 queue 通道中的两个值。
		queue := make(chan string, 2)
		queue <- "one"
		queue <- "two"
		close(queue)

		// This range iterates over each element as it’s received from queue. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
		// range 迭代从 queue 中得到每个值。 因为我们在前面 close 了这个通道，所以，这个迭代会在接收完 2 个值之后结束。

		for elem := range queue {
			fmt.Println(elem)
		}

		// This example also showed that it’s possible to close a non-empty channel but still have the remaining values be received.
		// 这个例子也让我们看到，一个非空的通道也是可以关闭的， 并且，通道中剩下的值仍然可以被接收到。
	},
}

func init() {
	cmd.RootCmd.AddCommand(rangeOverChannelsCmd)
}
