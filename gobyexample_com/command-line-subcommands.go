package gobyexample_com

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"go-labs/cmd"
	"os"

	"github.com/spf13/cobra"
)

// Some command-line tools, like the go tool or git have many subcommands, each with its own set of flags. For example, go build and go get are two different subcommands of the go tool. The flag package lets us easily define simple subcommands that have their own flags.
// go 和 git 这种命令行工具，都有很多的 子命令 。 并且每个工具都有一套自己的 flag，比如： go build 和 go get 是 go 里面的两个不同的子命令。 flag 包让我们可以轻松的为工具定义简单的子命令。

// commandLineSubcommandsCmd represents the command-line-subcommands command
var commandLineSubcommandsCmd = &cobra.Command{
	Use:   "go-by-example:command-line-subcommands",
	Short: "https://gobyexample.com/command-line-subcommands",
	Run: func(cmd *cobra.Command, args []string) {

		color.Yellow("注意⚠️")
		color.Red("️参数冲突，需要另起新文件main包方式执行此例️")
		color.Yellow("注意⚠️")

		// We declare a subcommand using the NewFlagSet function, and proceed to define new flags specific for this subcommand.
		// 我们使用 NewFlagSet 函数声明一个子命令， 然后为这个子命令定义一个专用的 flag。
		fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
		fooEnable := fooCmd.Bool("enable", false, "enable")
		fooName := fooCmd.String("name", "", "name")

		// For a different subcommand we can define different supported flags.
		// 对于不同的子命令，我们可以定义不同的 flag。
		barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
		barLevel := barCmd.Int("level", 0, "level")

		// The subcommand is expected as the first argument to the program.
		// 子命令应作为程序的第一个参数传入。
		if len(os.Args) < 2 {
			fmt.Println("expected 'foo' or 'bar' subcommands")
			os.Exit(1)
		}

		// Check which subcommand is invoked.
		// 检查哪一个子命令被调用了。
		switch os.Args[1] {

		// For every subcommand, we parse its own flags and have access to trailing positional arguments.
		// 每个子命令，都会解析自己的 flag 并允许它访问后续的位置参数。
		case "foo":
			fooCmd.Parse(os.Args[2:])
			fmt.Println("subcommand 'foo'")
			fmt.Println("  enable:", *fooEnable)
			fmt.Println("  name:", *fooName)
			fmt.Println("  tail:", fooCmd.Args())
		case "bar":
			barCmd.Parse(os.Args[2:])
			fmt.Println("subcommand 'bar'")
			fmt.Println("  level:", *barLevel)
			fmt.Println("  tail:", barCmd.Args())
		default:
			fmt.Println("expected 'foo' or 'bar' subcommands")
			os.Exit(1)
		}

		fmt.Println("go build command-line-subcommands.go ")
		// First invoke the foo subcommand.
		// 首先调用 foo 子命令。
		fmt.Println("./command-line-subcommands foo -enable -name=joe a1 a2")

		// Now try bar.
		// 然后试一下 bar 子命令。
		fmt.Println("./command-line-subcommands bar -level 8 a1")

		// But bar won’t accept foo’s flags.
		// 但是 bar 不接受 foo 的 flag（enable）。
		fmt.Println("./command-line-subcommands bar -enable a1")

		// Next we’ll look at environment variables, another common way to parameterize programs.
		// 接下来我们会学习程序获取参数的另一种常见方式——环境变量。
	},
}

func init() {
	cmd.RootCmd.AddCommand(commandLineSubcommandsCmd)
}
