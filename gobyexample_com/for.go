package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// forCmd represents the for command
var forCmd = &cobra.Command{
	Use:   "go-by-example:for",
	Short: "https://gobyexample.com/for",
	Run: func(cmd *cobra.Command, args []string) {

		// The most basic type, with a single condition.
		// 最基础的方式，单个循环条件。
		i := 1
		for i <= 3 {
			fmt.Println(i)
			i = i + 1
		}

		// A classic initial/condition/after for loop.
		// 经典的初始/条件/后续 for 循环。
		for j := 7; j <= 9; j++ {
			fmt.Println(j)
		}

		// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
		// 不带条件的 for 循环将一直重复执行， 直到在循环体内使用了 break 或者 return 跳出循环。
		for {
			fmt.Println("loop")
			break
		}

		// You can also continue to the next iteration of the loop.
		// 你也可以使用 continue 直接进入下一次循环。
		for n := 0; n <= 5; n++ {
			if n%2 == 0 {
				continue
			}
			fmt.Println(n)
		}

		// We’ll see some other for forms later when we look at range statements, channels, and other data structures.
		// 我们在后续教程中学习 range 语句，channels 以及其他数据结构时， 还会看到一些 for 的其它用法。

	},
}

func init() {
	cmd.RootCmd.AddCommand(forCmd)
}
