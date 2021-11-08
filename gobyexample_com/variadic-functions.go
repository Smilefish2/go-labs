package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
// 可变参数函数。 在调用时可以传递任意数量的参数。 例如，fmt.Println 就是一个常见的变参函数。

// Here’s a function that will take an arbitrary number of ints as arguments.
// 这个函数接受任意数量的 int 作为参数。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// variadicFunctionsCmd represents the variadic-functions command
var variadicFunctionsCmd = &cobra.Command{
	Use:   "go-by-example:variadic-functions",
	Short: "https://gobyexample.com/variadic-functions",
	Run: func(cmd *cobra.Command, args []string) {

		// Variadic functions can be called in the usual way with individual arguments.
		// 变参函数使用常规的调用方式，传入独立的参数。
		sum(1, 2)
		sum(1, 2, 3)

		// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
		// 如果你有一个含有多个值的 slice，想把它们作为参数使用， 你需要这样调用 func(slice...)。
		nums := []int{1, 2, 3, 4}
		sum(nums...)

		// Another key aspect of functions in Go is their ability to form closures, which we’ll look at next.
		// 接下来我们要看的是 Go 函数的另一个关键特性：闭包。
	},
}

func init() {
	cmd.RootCmd.AddCommand(variadicFunctionsCmd)
}
