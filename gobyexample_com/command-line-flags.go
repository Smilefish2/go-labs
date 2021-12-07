package gobyexample_com

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// Command-line flags are a common way to specify options for command-line programs. For example, in wc -l the -l is a command-line flag.
// 命令行标志 是命令行程序指定选项的常用方式。例如，在 wc -l 中， 这个 -l 就是一个命令行标志。

// Go provides a flag package supporting basic command-line flag parsing. We’ll use this package to implement our example command-line program.
// Go 提供了一个 flag 包，支持基本的命令行标志解析。 我们将用这个包来实现我们的命令行程序示例。

// commandLineFlagsCmd represents the command-line-flags command
var commandLineFlagsCmd = &cobra.Command{
	Use:   "go-by-example:command-line-flags",
	Short: "https://gobyexample.com/command-line-flags",
	Run: func(cmd *cobra.Command, args []string) {

		color.Yellow("注意⚠️")
		color.Red("️参数冲突，需要另起新文件main包方式执行此例️")
		color.Yellow("注意⚠️")

		// Basic flag declarations are available for string, integer, and boolean options. Here we declare a string flag word with a default value "foo" and a short description. This flag.String function returns a string pointer (not a string value); we’ll see how to use this pointer below.
		// 基本的标记声明仅支持字符串、整数和布尔值选项。 这里我们声明一个默认值为 "foo" 的字符串标志 word 并带有一个简短的描述。 这里的 flag.String 函数返回一个字符串指针（不是一个字符串值）， 在下面我们会看到是如何使用这个指针的。
		wordPtr := flag.String("word", "foo", "a string")

		// This declares numb and fork flags, using a similar approach to the word flag.
		// 使用和声明 word 标志相同的方法来声明 numb 和 fork 标志。
		numbPtr := flag.Int("numb", 42, "an int")
		forkPtr := flag.Bool("fork", false, "a bool")

		// It’s also possible to declare an option that uses an existing var declared elsewhere in the program. Note that we need to pass in a pointer to the flag declaration function.
		// 用程序中已有的参数来声明一个标志也是可以的。 注意在标志声明函数中需要使用该参数的指针。
		var svar string
		flag.StringVar(&svar, "svar", "bar", "a string var")

		// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
		// 所有标志都声明完成以后，调用 flag.Parse() 来执行命令行解析。
		flag.Parse()

		// Here we’ll just dump out the parsed options and any trailing positional arguments. Note that we need to dereference the pointers with e.g. *wordPtr to get the actual option values.
		// 这里我们将仅输出解析的选项以及后面的位置参数。 注意，我们需要使用类似 *wordPtr 这样的语法来对指针解引用， 从而得到选项真正的值。
		fmt.Println("word:", *wordPtr)
		fmt.Println("numb:", *numbPtr)
		fmt.Println("fork:", *forkPtr)
		fmt.Println("svar:", svar)
		fmt.Println("tail:", flag.Args())

		// To experiment with the command-line flags program it’s best to first compile it and then run the resulting binary directly.
		// 测试这个程序前，最好将这个程序编译成二进制文件，然后再运行这个程序。
		fmt.Println("go build command-line-flags.go")

		// Try out the built program by first giving it values for all flags.
		// 首先以给所有标志赋值的方式，尝试运行构建的程序。
		fmt.Println("./command-line-flags -word=opt -numb=7 -fork -svar=flag")

		// Note that if you omit flags they automatically take their default values.
		// 注意，如果你省略一个标志，那么这个标志的值自动的设定为他的默认值。
		fmt.Println("./command-line-flags -word=opt")

		// Trailing positional arguments can be provided after any flags.
		// 尾随的位置参数可以出现在任何标志后面。
		fmt.Println("./command-line-flags -word=opt a1 a2 a3")

		// Note that the flag package requires all flags to appear before positional arguments (otherwise the flags will be interpreted as positional arguments).
		// 注意，flag 包需要所有的标志出现位置参数之前（否则，这个标志将会被解析为位置参数）。
		fmt.Println("./command-line-flags -word=opt a1 a2 a3 -numb=7")

		// Use -h or --help flags to get automatically generated help text for the command-line program.
		// 使用 -h 或者 --help 标志来得到自动生成的这个命令行程序的帮助文本。
		fmt.Println("./command-line-flags -h")

		// If you provide a flag that wasn’t specified to the flag package, the program will print an error message and show the help text again.
		// 如果你提供了一个没有使用 flag 包声明的标志， 程序会输出一个错误信息，并再次显示帮助文本。
		fmt.Println("./command-line-flags -wat")
	},
}

func init() {
	cmd.RootCmd.AddCommand(commandLineFlagsCmd)
}
