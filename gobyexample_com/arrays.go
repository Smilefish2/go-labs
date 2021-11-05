package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// arraysCmd represents the arrays command
var arraysCmd = &cobra.Command{
	Use:   "go-by-example:arrays",
	Short: "https://gobyexample.com/arrays",
	Run: func(cmd *cobra.Command, args []string) {
		// Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for ints means 0s.
		// 这里我们创建了一个刚好可以存放 5 个 int 元素的数组 a。 元素的类型和长度都是数组类型的一部分。 数组默认值是零值，对于 int 数组来说，元素的零值是 0。
		var a [5]int
		fmt.Println("emp:", a)

		// We can set a value at an index using the array[index] = value syntax, and get a value with array[index].
		// 我们可以使用 array[index] = value 语法来设置数组指定位置的值， 或者用 array[index] 得到值。
		a[4] = 100
		fmt.Println("set:", a)
		fmt.Println("get:", a[4])

		// The builtin len returns the length of an array.
		// 内置函数 len 可以返回数组的长度。
		fmt.Println("len:", len(a))

		// Use this syntax to declare and initialize an array in one line.
		// 使用这个语法在一行内声明并初始化一个数组。
		b := [5]int{1, 2, 3, 4, 5}
		fmt.Println("dcl:", b)

		// Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
		// 数组类型是一维的，但是你可以组合构造多维的数据结构。
		var twoD [2][3]int
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				twoD[i][j] = i + j
			}
		}
		fmt.Println("2d: ", twoD)

		// Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.
		// 注意，使用 fmt.Println 打印数组时，会按照 [v1 v2 v3 ...] 的格式打印。

		// You’ll see slices much more often than arrays in typical Go. We’ll look at slices next.
		// 在 Go 程序中，相较于数组，用得更多的是 _切片(slice)_。我们接着来看切片。
	},
}

func init() {
	cmd.RootCmd.AddCommand(arraysCmd)
}
