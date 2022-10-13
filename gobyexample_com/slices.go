package gobyexample_com

import (
	"fmt"
	"go-labs/cmd"

	"github.com/spf13/cobra"
)

// slicesCmd represents the slices command
var slicesCmd = &cobra.Command{
	Use:   "go-by-example:slices",
	Short: "https://gobyexample.com/slices",
	Run: func(cmd *cobra.Command, args []string) {

		// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
		// Slice 是 Go 中一个重要的数据类型，它提供了比数组更强大的序列交互方式。

		// Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).
		// 与数组不同，slice 的类型仅由它所包含的元素的类型决定（与元素个数无关）。 要创建一个长度不为 0 的空 slice，需要使用内建函数 make。 这里我们创建了一个长度为 3 的 string 类型的 slice（初始值为零值）。
		s := make([]string, 3)
		fmt.Println("emp:", s)

		// We can set and get just like with arrays.
		// 我们可以和数组一样设置和得到值

		s[0] = "a"
		s[1] = "b"
		s[2] = "c"
		fmt.Println("set:", s)
		fmt.Println("get:", s[2])

		// len returns the length of the slice as expected.
		// len 返回 slice 的长度

		fmt.Println("len:", len(s))

		// In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.
		// 除了基本操作外，slice 支持比数组更丰富的操作。比如 slice 支持内建函数 append， 该函数会返回一个包含了一个或者多个新值的 slice。 注意由于 append 可能返回一个新的 slice，我们需要接收其返回值。
		s = append(s, "d")
		s = append(s, "e", "f")
		fmt.Println("apd:", s)

		// Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
		// slice 还可以 copy。这里我们创建一个空的和 s 有相同长度的 slice——c， 然后将 s 复制给 c。

		c := make([]string, len(s))
		copy(c, s)
		fmt.Println("cpy:", c)

		// Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].
		// slice 支持通过 slice[low:high] 语法进行“切片”操作。 例如，右边的操作可以得到一个包含元素 s[2]、s[3] 和 s[4] 的 slice。
		l := s[2:5]
		fmt.Println("sl1:", l)

		// This slices up to (but excluding) s[5].
		// 这个 slice 包含从 s[0] 到 s[5]（不包含 5）的元素。
		l = s[:5]
		fmt.Println("sl2:", l)

		// And this slices up from (and including) s[2].
		// 这个 slice 包含从 s[2]（包含 2）之后的元素。
		l = s[2:]
		fmt.Println("sl3:", l)

		// We can declare and initialize a variable for slice in a single line as well.
		// 我们可以在一行代码中声明并初始化一个 slice 变量。

		t := []string{"g", "h", "i"}
		fmt.Println("dcl:", t)

		// Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
		// Slice 可以组成多维数据结构。内部的 slice 长度可以不一致，这一点和多维数组不同。

		twoD := make([][]int, 3)
		for i := 0; i < 3; i++ {
			innerLen := i + 1
			twoD[i] = make([]int, innerLen)
			for j := 0; j < innerLen; j++ {
				twoD[i][j] = i + j
			}
		}
		fmt.Println("2d: ", twoD)

		// Note that while slices are different types than arrays, they are rendered similarly by fmt.Println.
		// 注意，slice 和数组是不同的类型，但它们通过 fmt.Println 打印的输出结果是类似的。
		// http://blog.golang.org/2011/01/go-slices-usage-and-internals.html

		// Check out this great blog post by the Go team for more details on the design and implementation of slices in Go.
		// 看看这个由 Go 团队撰写的一篇很棒的博文，了解更多关于 Go 中 slice 的设计和实现细节。

		// Now that we’ve seen arrays and slices we’ll look at Go’s other key builtin data structure: maps.
		// 现在，我们已经学习了数组和 slice，接下来我们将学习 Go 中的另一个重要的内建数据类型：map。

	},
}

func init() {
	cmd.RootCmd.AddCommand(slicesCmd)
}
