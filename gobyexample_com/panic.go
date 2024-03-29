package gobyexample_com

import (
	"github.com/spf13/cobra"
	"go-labs/cmd"
	"os"
)

// A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.
// panic 意味着有些出乎意料的错误发生。 通常我们用它来表示程序正常运行中不应该出现的错误， 或者我们不准备优雅处理的错误。

// panicCmd represents the panic command
var panicCmd = &cobra.Command{
	Use:   "go-by-example:panic",
	Short: "https://gobyexample.com/panic",
	Run: func(cmd *cobra.Command, args []string) {

		// We’ll use panic throughout this site to check for unexpected errors. This is the only program on the site designed to panic.
		// 我们将使用 panic 来检查这个站点上预期之外的错误。 而该站点上只有一个程序：触发 panic。
		panic("a problem")

		// A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle. Here’s an example of panicking if we get an unexpected error when creating a new file.
		// panic 的一种常见用法是：当函数返回我们不知道如何处理（或不想处理）的错误值时，中止操作。 如果创建新文件时遇到意外错误该如何处理？这里有一个很好的 panic 示例。
		_, err := os.Create("/tmp/file")
		if err != nil {
			panic(err)
		}

		// 运行程序将会导致 panic： 输出一个错误消息和协程追踪信息，并以非零的状态退出程序。
		// 注意，与某些使用 exception 处理错误的语言不同， 在 Go 中，通常会尽可能的使用返回值来标示错误。

	},
}

func init() {
	cmd.RootCmd.AddCommand(panicCmd)
}
