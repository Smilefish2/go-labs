package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// Go supports methods defined on struct types.
// Go 支持为结构体类型定义方法(methods) 。

type rect struct {
	width, height int
}

// This area method has a receiver type of *rect.
// 这里的 area 是一个拥有 *rect 类型接收器(receiver)的方法。
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
// 可以为值类型或者指针类型的接收者定义方法。 这是一个值类型接收者的例子。
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// methodsCmd represents the methods command
var methodsCmd = &cobra.Command{
	Use:   "go-by-example:methods",
	Short: "https://gobyexample.com/methods",
	Run: func(cmd *cobra.Command, args []string) {
		r := rect{width: 10, height: 5}

		// Here we call the 2 methods defined for our struct.
		// 这里我们调用上面为结构体定义的两个方法。
		fmt.Println("area: ", r.area())
		fmt.Println("perim:", r.perim())

		// Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
		// 调用方法时，Go 会自动处理值和指针之间的转换。 想要避免在调用方法时产生一个拷贝，或者想让方法可以修改接受结构体的值， 你都可以使用指针来调用方法。
		rp := &r
		fmt.Println("area: ", rp.area())
		fmt.Println("perim:", rp.perim())

		// Next we’ll look at Go’s mechanism for grouping and naming related sets of methods: interfaces.
		// 接下来，我们将学习 Go 对方法集进行命名和分组的另一机制：接口。
	},
}

func init() {
	cmd.RootCmd.AddCommand(methodsCmd)
}
