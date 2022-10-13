package gobyexample_com

import (
	"fmt"
	"go-labs/cmd"

	"github.com/spf13/cobra"
)

// valuesCmd represents the values command
var valuesCmd = &cobra.Command{
	Use:   "go-by-example:values",
	Short: "https://gobyexample.com/values",
	Run: func(cmd *cobra.Command, args []string) {

		// Strings, which can be added together with +.
		// 字符串可以通过 + 连接。
		fmt.Println("go" + "lang")
		fmt.Println("go" + " " + "lang")

		// Integers and floats.
		// 整数和浮点数
		fmt.Println("1+1 =", 1+1)
		fmt.Println("1-1 =", 1-1)
		fmt.Println("7.0*3.0 =", 7.0*3.0)
		fmt.Println("7.0/3.0 =", 7.0/3.0)

		// Booleans, with boolean operators as you’d expect.
		// 布尔型，以及常见的布尔操作。
		fmt.Println(true && false)
		fmt.Println(true || false)
		fmt.Println(!true)
	},
}

func init() {
	cmd.RootCmd.AddCommand(valuesCmd)
}
