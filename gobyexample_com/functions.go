package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// Functions are central in Go. We’ll learn about functions with a few different examples.
// 函数 是 Go 的核心。我们将通过一些不同的例子来进行学习它。

// Here’s a function that takes two ints and returns their sum as an int.
// 这里是一个函数，接受两个 int 并且以 int 返回它们的和
func plus(a int, b int) int {

	// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
	// Go 需要明确的 return，也就是说，它不会自动 return 最后一个表达式的值
	return a + b
}

// When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
// 当多个连续的参数为同样类型时， 可以仅声明最后一个参数的类型，忽略之前相同类型参数的类型声明。
func plusPlus(a, b, c int) int {
	return a + b + c
}

// functionsCmd represents the functions command
var functionsCmd = &cobra.Command{
	Use:   "go-by-example:functions",
	Short: "https://gobyexample.com/functions",
	Run: func(cmd *cobra.Command, args []string) {

		// Call a function just as you’d expect, with name(args).
		// 通过 函数名(参数列表) 来调用函数，
		res := plus(1, 2)
		fmt.Println("1+2 =", res)

		// There are several other features to Go functions. One is multiple return values, which we’ll look at next.
		// Go 函数还有很多其他的特性。 其中一个就是多值返回，它也是我们接下来要接触的。
		res = plusPlus(1, 2, 3)
		fmt.Println("1+2+3 =", res)
	},
}

func init() {
	cmd.RootCmd.AddCommand(functionsCmd)
}
