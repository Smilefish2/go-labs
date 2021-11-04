package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// helloWorldCmd represents the hello-world command
var helloWorldCmd = &cobra.Command{
	Use:   "go-by-example:hello-world",
	Short: "https://gobyexample.com/hello-world",
	Run: func(cmd *cobra.Command, args []string) {

		// To run the program, put the code in hello-world.go and use go run.
		// 要运行这个程序，先将将代码放到名为 hello-world.go 的文件中，然后执行 go run。
		//
		// $ go run hello-world.go
		//	 hello world
		// Sometimes we’ll want to build our programs into binaries. We can do this using go build.
		// 如果我们想将程序编译成二进制文件（Windows 平台是 .exe 可执行文件）， 可以通过 go build 来达到目的。
		//
		// $ go build hello-world.go
		// $ ls
		// hello-world    hello-world.go
		// We can then execute the built binary directly.
		// 然后我们可以直接运行这个二进制文件。
		//
		// $ ./hello-world
		// hello world
		// Now that we can run and build basic Go programs, let’s learn more about the language.
		// 现在我们可以运行和编译基础的 Go 程序了， 让我们开始学习更多关于这门语言的知识吧。

		fmt.Println("hello world")
	},
}

func init() {
	cmd.RootCmd.AddCommand(helloWorldCmd)
}
