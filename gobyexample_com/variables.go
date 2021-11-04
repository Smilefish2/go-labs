package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// variablesCmd represents the variables command
var variablesCmd = &cobra.Command{
	Use:   "go-by-example:variables",
	Short: "https://gobyexample.com/variables",
	Run: func(cmd *cobra.Command, args []string) {

		// var declares 1 or more variables.
		// var 声明 1 个或者多个变量。
		var a = "initial"
		fmt.Println(a)

		// You can declare multiple variables at once.
		// 你可以一次性声明多个变量。
		var b, c int = 1, 2
		fmt.Println(b, c)

		// Go will infer the type of initialized variables.
		// Go 会自动推断已经有初始值的变量的类型。
		var d = true
		fmt.Println(d)

		// Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
		// 声明后却没有给出对应的初始值时，变量将会初始化为 零值 。 例如，int 的零值是 0。
		var e int
		fmt.Println(e)

		// The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
		// := 语法是声明并初始化变量的简写， 例如 var f string = "short" 可以简写为右边这样。
		f := "apple"
		fmt.Println(f)
	},
}

func init() {
	cmd.RootCmd.AddCommand(variablesCmd)
}
