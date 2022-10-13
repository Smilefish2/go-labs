package gobyexample_com

import (
	"fmt"
	"go-labs/cmd"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// Sometimes we’d like our Go programs to intelligently handle Unix signals. For example, we might want a server to gracefully shutdown when it receives a SIGTERM, or a command-line tool to stop processing input if it receives a SIGINT. Here’s how to handle signals in Go with channels.
// 有时候，我们希望 Go 可以智能的处理 Unix 信号。 例如，我们希望当服务器接收到一个 SIGTERM 信号时，能够优雅退出， 或者一个命令行工具在接收到一个 SIGINT 信号时停止处理输入信息。 我们这里讲的就是在 Go 中如何使用通道来处理信号。

// signalsCmd represents the signals command
var signalsCmd = &cobra.Command{
	Use:   "go-by-example:signals",
	Short: "https://gobyexample.com/signals",
	Run: func(cmd *cobra.Command, args []string) {

		// Go signal notification works by sending os.Signal values on a channel. We’ll create a channel to receive these notifications. Note that this channel should be buffered.
		// Go 通过向一个通道发送 os.Signal 值来发送信号通知。 我们将创建一个通道来接收这些通知（同时再创建一个在程序结束时发送通知的通道）。
		sigs := make(chan os.Signal, 1)

		// signal.Notify registers the given channel to receive notifications of the specified signals.
		// signal.Notify 注册给定的通道，用于接收特定信号。
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		// We could receive from sigs here in the main function, but let’s see how this could also be done in a separate goroutine, to demonstrate a more realistic scenario of graceful shutdown.
		// 这个协程执行一个阻塞的信号接收操作。 当它接收到一个值时，它将打印这个值，然后通知程序可以退出。
		done := make(chan bool, 1)

		// This goroutine executes a blocking receive for signals. When it gets one it’ll print it out and then notify the program that it can finish.
		// 这个协程执行一个阻塞的信号接收操作。 当它接收到一个值时，它将打印这个值，然后通知程序可以退出。
		go func() {

			sig := <-sigs
			fmt.Println()
			fmt.Println(sig)
			done <- true
		}()

		// The program will wait here until it gets the expected signal (as indicated by the goroutine above sending a value on done) and then exit.
		// 程序将在这里进行等待，直到它得到了期望的信号 （也就是上面的协程发送的 done 值），然后退出。
		fmt.Println("awaiting signal")
		<-done
		fmt.Println("exiting")

		// When we run this program it will block waiting for a signal. By typing ctrl-C (which the terminal shows as ^C) we can send a SIGINT signal, causing the program to print interrupt and then exit.
		// 当我们运行这个程序时，它将一直等待一个信号。 通过 ctrl-C（终端显示为 ^C），我们可以发送一个 SIGINT 信号， 这会使程序打印 interrupt 然后退出。
	},
}

func init() {
	cmd.RootCmd.AddCommand(signalsCmd)
}
