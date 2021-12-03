package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"
	"sort"

	"github.com/spf13/cobra"
)

// Go’s sort package implements sorting for builtins and user-defined types. We’ll look at sorting for builtins first.
// Go 的 sort 包实现了内建及用户自定义数据类型的排序功能。 我们先来看看内建数据类型的排序。

// sortingCmd represents the sorting command
var sortingCmd = &cobra.Command{
	Use:   "go-by-example:sorting",
	Short: "https://gobyexample.com/sorting",
	Run: func(cmd *cobra.Command, args []string) {

		// Sort methods are specific to the builtin type; here’s an example for strings. Note that sorting is in-place, so it changes the given slice and doesn’t return a new one.
		// 排序方法是针对内置数据类型的； 这是一个字符串排序的例子。 注意，它是原地排序的，所以他会直接改变给定的切片，而不是返回一个新切片。
		strs := []string{"c", "a", "b"}
		sort.Strings(strs)
		fmt.Println("Strings:", strs)

		// An example of sorting ints.
		// 一个 int 排序的例子。
		ints := []int{7, 2, 4}
		sort.Ints(ints)
		fmt.Println("Ints:   ", ints)

		// We can also use sort to check if a slice is already in sorted order.
		// 我们也可以使用 sort 来检查一个切片是否为有序的。
		s := sort.IntsAreSorted(ints)
		fmt.Println("Sorted: ", s)

		// Running our program prints the sorted string and int slices and true as the result of our AreSorted test.
		// 运行程序，打印排序好的字符串和整型切片， 以及数组是否有序的检查结果：true。
	},
}

func init() {
	cmd.RootCmd.AddCommand(sortingCmd)
}
