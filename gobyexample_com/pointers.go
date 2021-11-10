package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// Go supports pointers, allowing you to pass references to values and records within your program.
// Go 支持 指针， 允许在程序中通过 引用传递 来传递值和数据结构。

// We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr. zeroval has an int parameter, so arguments will be passed to it by value. zeroval will get a copy of ival distinct from the one in the calling function.
// 我们将通过两个函数：zeroval 和 zeroptr 来比较 指针 和 值。 zeroval 有一个 int 型参数，所以使用值传递。 zeroval 将从调用它的那个函数中得到一个实参的拷贝：ival。
func zeroval(ival int) {
	ival = 0
}

// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer. The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address. Assigning a value to a dereferenced pointer changes the value at the referenced address.
// zeroptr 有一个和上面不同的参数：*int，这意味着它使用了 int 指针。 紧接着，函数体内的 *iptr 会 解引用 这个指针，从它的内存地址得到这个地址当前对应的值。 对解引用的指针赋值，会改变这个指针引用的真实地址的值。
func zeroptr(iptr *int) {
	*iptr = 0
}

// pointersCmd represents the pointers command
var pointersCmd = &cobra.Command{
	Use:   "go-by-example:pointers",
	Short: "https://gobyexample.com/pointers",
	Run: func(cmd *cobra.Command, args []string) {

		i := 1
		fmt.Println("initial:", i)

		zeroval(i)
		fmt.Println("zeroval:", i)

		// The &i syntax gives the memory address of i, i.e. a pointer to i.
		// 通过 &i 语法来取得 i 的内存地址，即指向 i 的指针。
		zeroptr(&i)
		fmt.Println("zeroptr:", i)

		// Pointers can be printed too.
		// 指针也是可以被打印的。
		fmt.Println("pointer:", &i)

		// zeroval doesn’t change the i in main, but zeroptr does because it has a reference to the memory address for that variable.
		// zeroval 在 main 函数中不能改变 i 的值， 但是 zeroptr 可以，因为它有这个变量的内存地址的引用。
	},
}

func init() {
	cmd.RootCmd.AddCommand(pointersCmd)
}
