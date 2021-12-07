package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"
	"os"

	"github.com/spf13/cobra"
)

// Command-line arguments are a common way to parameterize execution of programs. For example, go run hello.go uses run and hello.go arguments to the go program.
// 命令行参数 是指定程序运行参数的一个常见方式。例如，go run hello.go， 程序 go 使用了 run 和 hello.go 两个参数。

// commandLineArgumentsCmd represents the command-line-arguments command
var commandLineArgumentsCmd = &cobra.Command{
	Use:   "go-by-example:command-line-arguments",
	Short: "https://gobyexample.com/command-line-arguments",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("run: go run main.go go-by-example:command-line-arguments a b c d")

		// os.Args provides access to raw command-line arguments. Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.
		// os.Args 提供原始命令行参数访问功能。 注意，切片中的第一个参数是该程序的路径， 而 os.Args[1:]保存了程序全部的参数。
		argsWithProg := os.Args
		argsWithoutProg := os.Args[1:]

		// You can get individual args with normal indexing.
		// 你可以使用标准的下标方式取得单个参数的值。
		arg := os.Args[3]

		fmt.Println(argsWithProg)
		fmt.Println(argsWithoutProg)
		fmt.Println(arg)

		// To experiment with command-line arguments it’s best to build a binary with go build first.
		// 要实验命令行参数，最好先使用 go build 编译一个可执行二进制文件

		// Next we’ll look at more advanced command-line processing with flags.
		// 下面我们要看看更高级的使用标记的命令行处理方法。
	},
}

func init() {
	cmd.RootCmd.AddCommand(commandLineArgumentsCmd)
}
