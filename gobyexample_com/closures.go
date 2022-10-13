package gobyexample_com

import (
	"fmt"
	"go-labs/cmd"

	"github.com/spf13/cobra"
)

// Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.
// Go 支持匿名函数， 并能用其构造 闭包。 匿名函数在你想定义一个不需要命名的内联函数时是很实用的。

// This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.
// intSeq 函数返回一个在其函数体内定义的匿名函数。 返回的函数使用闭包的方式 隐藏 变量 i。 返回的函数 隐藏 变量 i 以形成闭包。
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// closuresCmd represents the closures command
var closuresCmd = &cobra.Command{
	Use:   "go-by-example:closures",
	Short: "https://gobyexample.com/closures",
	Run: func(cmd *cobra.Command, args []string) {

		// We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
		// 我们调用 intSeq 函数，将返回值（一个函数）赋给 nextInt。 这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时，都会更新 i 的值。
		nextInt := intSeq()

		// See the effect of the closure by calling nextInt a few times.
		// 通过多次调用 nextInt 来看看闭包的效果。
		fmt.Println(nextInt())
		fmt.Println(nextInt())
		fmt.Println(nextInt())

		// To confirm that the state is unique to that particular function, create and test a new one.
		// 为了确认这个状态对于这个特定的函数是唯一的，我们重新创建并测试一下。
		newInts := intSeq()
		fmt.Println(newInts())

		// The last feature of functions we’ll look at for now is recursion.
		// 我们马上要学习关于函数的最后一个特性：递归。
	},
}

func init() {
	cmd.RootCmd.AddCommand(closuresCmd)
}
