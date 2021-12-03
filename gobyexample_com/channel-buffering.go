package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.
// 默认情况下，通道是 无缓冲 的，这意味着只有对应的接收（<- chan） 通道准备好接收时，才允许进行发送（chan <-）。 有缓冲通道 允许在没有对应接收者的情况下，缓存一定数量的值。

// channelBufferingCmd represents the channel Buffering command
var channelBufferingCmd = &cobra.Command{
	Use:   "go-by-example:channel-buffering",
	Short: "https://gobyexample.com/channel-buffering",
	Run: func(cmd *cobra.Command, args []string) {

		// Here we make a channel of strings buffering up to 2 values.
		// 这里我们 make 了一个字符串通道，最多允许缓存 2 个值。
		messages := make(chan string, 3)

		// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
		// 由于此通道是有缓冲的， 因此我们可以将这些值发送到通道中，而无需并发的接收。
		messages <- "buffered"
		messages <- "channel"
		messages <- "channel3"

		// Later we can receive these two values as usual.
		// 然后我们可以正常接收这两个值。
		fmt.Println(<-messages)
		fmt.Println(<-messages)
		fmt.Println(<-messages)
	},
}

func init() {
	cmd.RootCmd.AddCommand(channelBufferingCmd)
}
