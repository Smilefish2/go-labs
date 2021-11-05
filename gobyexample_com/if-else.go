package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// ifElseCmd represents the if-else command
var ifElseCmd = &cobra.Command{
	Use:   "go-by-example:if-else",
	Short: "https://gobyexample.com/if-else",
	Run: func(cmd *cobra.Command, args []string) {

		// Here’s a basic example.
		// 这里是一个基本的例子。
		if 7%2 == 0 {
			fmt.Println("7 is even")
		} else {
			fmt.Println("7 is odd")
		}

		// You can have an if statement without an else.
		// 你可以不要 else 只用 if 语句。
		if 8%4 == 0 {
			fmt.Println("8 is divisible by 4")
		}

		// A statement can precede conditionals; any variables declared in this statement are available in all branches.
		// 在条件语句之前可以有一个声明语句；在这里声明的变量可以在这个语句所有的条件分支中使用。
		if num := 9; num < 0 {
			fmt.Println(num, "is negative")
		} else if num < 10 {
			fmt.Println(num, "has 1 digit")
		} else {
			fmt.Println(num, "has multiple digits")
		}

		// Note that you don’t need parentheses around conditions in Go, but that the braces are required.
		// 注意，在 Go 中，条件语句的圆括号不是必需的，但是花括号是必需的。

		// There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
		// Go 没有三目运算符， 即使是基本的条件判断，依然需要使用完整的 if 语句。
	},
}

func init() {
	cmd.RootCmd.AddCommand(ifElseCmd)
}
