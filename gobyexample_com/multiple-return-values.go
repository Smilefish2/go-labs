package gobyexample_com

import (
	"fmt"
	"go-labs/cmd"

	"github.com/spf13/cobra"
)

// Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.
// Go 原生支持 _多返回值_。 这个特性在 Go 语言中经常用到，例如用来同时返回一个函数的结果和错误信息。

// The (int, int) in this function signature shows that the function returns 2 ints.
// (int, int) 在这个函数中标志着这个函数返回 2 个 int。
func vals() (int, int) {
	return 3, 7
}

// multipleReturnValuesCmd represents the multiple-return-values command
var multipleReturnValuesCmd = &cobra.Command{
	Use:   "go-by-example:multiple-return-values",
	Short: "https://gobyexample.com/multiple-return-values",
	Run: func(cmd *cobra.Command, args []string) {

		// Here we use the 2 different return values from the call with multiple assignment.
		// 这里我们通过 多赋值 操作来使用这两个不同的返回值。
		a, b := vals()
		fmt.Println(a)
		fmt.Println(b)

		// If you only want a subset of the returned values, use the blank identifier _.
		// 如果你仅仅需要返回值的一部分的话，你可以使用空白标识符 _。
		_, c := vals()
		fmt.Println(c)

		// Accepting a variable number of arguments is another nice feature of Go functions; we’ll look at this next.
		// 我们接下来要学习的是 Go 函数另一个很好的特性：变参函数。
	},
}

func init() {
	cmd.RootCmd.AddCommand(multipleReturnValuesCmd)
}
