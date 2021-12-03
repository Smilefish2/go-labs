package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
)

// When using channels as function parameters, you can specify if a channel is meant to only send or receive values. This specificity increases the type-safety of the program.
// 当使用通道作为函数的参数时，你可以指定这个通道是否为只读或只写。 该特性可以提升程序的类型安全。

// This ping function only accepts a channel for sending values. It would be a compile-time error to try to receive on this channel.
// ping 函数定义了一个只能发送数据的（只写）通道。 尝试从这个通道接收数据会是一个编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function accepts one channel for receives (pings) and a second for sends (pongs).
// pong 函数接收两个通道，pings 仅用于接收数据（只读），pongs 仅用于发送数据（只写）。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// channelDirectionsCmd represents the Channel Directions command
var channelDirectionsCmd = &cobra.Command{
	Use:   "go-by-example:channel-directions",
	Short: "https://gobyexample.com/channel-directions",
	Run: func(cmd *cobra.Command, args []string) {

		pings := make(chan string, 1)
		pongs := make(chan string, 1)
		ping(pings, "passed message")
		pong(pings, pongs)
		fmt.Println(<-pongs)

	},
}

func init() {
	cmd.RootCmd.AddCommand(channelDirectionsCmd)
}
