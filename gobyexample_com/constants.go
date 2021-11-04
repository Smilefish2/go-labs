package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"
	"math"

	"github.com/spf13/cobra"
)

// constantsCmd represents the constants command
var constantsCmd = &cobra.Command{
	Use:   "go-by-example:constants",
	Short: "https://gobyexample.com/constants",
	Run: func(cmd *cobra.Command, args []string) {

		// const declares a constant value.
		// const 用于声明一个常量。
		const s string = "constant"
		fmt.Println(s)

		// A const statement can appear anywhere a var statement can.
		// const 语句可以出现在任何 var 语句可以出现的地方
		const n = 500000000

		// Constant expressions perform arithmetic with arbitrary precision.
		// 常数表达式可以执行任意精度的运算
		const d = 3e20 / n
		fmt.Println(d)

		// A numeric constant has no type until it’s given one, such as by an explicit conversion.
		// 数值型常量没有确定的类型，直到被给定某个类型，比如显式类型转化。
		fmt.Println(int64(d))

		// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
		// 一个数字可以根据上下文的需要（比如变量赋值、函数调用）自动确定类型。 举个例子，这里的 math.Sin 函数需要一个 float64 的参数，n 会自动确定类型。
		fmt.Println(math.Sin(n))
	},
}

func init() {
	cmd.RootCmd.AddCommand(constantsCmd)
}
