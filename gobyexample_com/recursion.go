package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// Go supports recursive functions. Here’s a classic example.
// Go 支持 递归。 这里是一个经典的阶乘示例。

// This fact function calls itself until it reaches the base case of fact(0).
// fact 函数在到达 fact(0) 前一直调用自身。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// functionsCmd represents the recursion command
var recursionCmd = &cobra.Command{
	Use:   "go-by-example:recursion",
	Short: "https://gobyexample.com/recursion",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(fact(7))

		var fib func(n int) int

		// Closures can also be recursive, but this requires the closure to be declared with a typed var explicitly before it’s defined.
		// 闭包也可以是递归的，但这需要在定义闭包之前用类型化变量显式声明闭包。
		fib = func(n int) int {
			if n < 2 {
				return n
			}
			return fib(n-1) + fib(n-2)
		}

		// Since fib was previously declared in main, Go knows which function to call with fib here.
		// 由于fib以前是在main中声明的，所以Go知道在这里用fib调用哪个函数。

		fmt.Println(fib(7))

	},
}

func init() {
	cmd.RootCmd.AddCommand(recursionCmd)
}
