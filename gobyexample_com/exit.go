package gobyexample_com

import (
	"fmt"
	"go-labs/cmd"
	"os"

	"github.com/spf13/cobra"
)

// Use os.Exit to immediately exit with a given status.
// 使用 os.Exit 可以立即以给定的状态退出程序。

// exitCmd represents the exit command
var exitCmd = &cobra.Command{
	Use:   "go-by-example:exit",
	Short: "https://gobyexample.com/exit",
	Run: func(cmd *cobra.Command, args []string) {

		// defers will not be run when using os.Exit, so this fmt.Println will never be called.
		// 当使用 os.Exit 时 defer 将不会 被执行， 所以这里的 fmt.Println 将永远不会被调用。
		defer fmt.Println("!")

		// Exit with status 3.
		// 退出并且退出状态为 3。
		os.Exit(3)

		// Note that unlike e.g. C, Go does not use an integer return value from main to indicate exit status. If you’d like to exit with a non-zero status you should use os.Exit.
		// 注意，不像例如 C 语言，Go 不使用在 main 中返回一个整数来指明退出状态。 如果你想以非零状态退出，那么你就要使用 os.Exit。

		// If you run exit.go using go run, the exit will be picked up by go and printed.
		// 如果你使用 go run 来运行 exit.go，那么退出状态将会被 go 捕获并打印。
		fmt.Println("go run exit.go")

		// By building and executing a binary you can see the status in the terminal.
		// 通过编译并执行一个二进制文件的方式，你可以在终端中查看退出状态。

		// Note that the ! from our program never got printed.
		// 注意，程序中的 ! 永远不会被打印出来。
	},
}

func init() {
	cmd.RootCmd.AddCommand(exitCmd)
}
