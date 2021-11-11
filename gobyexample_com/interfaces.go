package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"
	"math"

	"github.com/spf13/cobra"
)

// Interfaces are named collections of method signatures.
// 方法签名的集合叫做：_接口(Interfaces)_。

// Here’s a basic interface for geometric shapes.
// 这是一个几何体的基本接口。
type geometry interface {
	area() float64
	perim() float64
}

// For our example we’ll implement this interface on rect and circle types.
// 在这个例子中，我们将为 rect 和 circle 实现该接口。

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to implement all the methods in the interface. Here we implement geometry on rects.
// 要在 Go 中实现一个接口，我们只需要实现接口中的所有方法。 这里我们为 rect 实现了 geometry 接口。
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for circles.
// circle 的实现。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call methods that are in the named interface. Here’s a generic measure function taking advantage of this to work on any geometry.
// 如果一个变量实现了某个接口，我们就可以调用指定接口中的方法。 这儿有一个通用的 measure 函数，我们可以通过它来使用所有的 geometry。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// interfacesCmd represents the interfaces command
var interfacesCmd = &cobra.Command{
	Use:   "go-by-example:interfaces",
	Short: "https://gobyexample.com/interfaces",
	Run: func(cmd *cobra.Command, args []string) {

		r := rect{width: 3, height: 4}
		c := circle{radius: 5}

		// The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
		// 结构体类型 circle 和 rect 都实现了 geometry 接口， 所以我们可以将其实例作为 measure 的参数。
		measure(r)
		measure(c)

		// To learn more about Go’s interfaces, check out this great blog post.
		// 要学习更多关于 Go 接口的知识， 可以看看这篇很棒的博文。
		// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
	},
}

func init() {
	cmd.RootCmd.AddCommand(interfacesCmd)
}
