package gobyexample_com

import (
	"errors"
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C. Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for any other, non-error tasks.
// 符合 Go 语言习惯的做法是使用一个独立、明确的返回值来传递错误信息。 这与 Java、Ruby 使用的异常（exception） 以及在 C 语言中有时用到的重载 (overloaded) 的单返回/错误值有着明显的不同。 Go 语言的处理方式能清楚的知道哪个函数返回了错误，并使用跟其他（无异常处理的）语言类似的方式来处理错误。

// By convention, errors are the last return value and have type error, a built-in interface.
// 按照惯例，错误通常是最后一个返回值并且是 error 类型，它是一个内建的接口。
func f1(arg int) (int, error) {
	if arg == 42 {

		// errors.New constructs a basic error value with the given error message.
		// errors.New 使用给定的错误信息构造一个基本的 error 值。
		return -1, errors.New("can't work with 42")
	}

	// A nil value in the error position indicates that there was no error.
	// 返回错误值为 nil 代表没有错误。
	return arg + 3, nil
}

// It’s possible to use custom types as errors by implementing the Error() method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.
// 你还可以通过实现 Error() 方法来自定义 error 类型。 这里使用自定义错误类型来表示上面例子中的参数错误。
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// In this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob.
		// 在这个例子中，我们使用 &argError 语法来建立一个新的结构体， 并提供了 arg 和 prob 两个字段的值。
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// errorsCmd represents the errors command
var errorsCmd = &cobra.Command{
	Use:   "go-by-example:errors",
	Short: "https://gobyexample.com/errors",
	Run: func(cmd *cobra.Command, args []string) {

		// The two loops below test out each of our error-returning functions. Note that the use of an inline error check on the if line is a common idiom in Go code.
		// 下面的两个循环测试了每一个会返回错误的函数。 注意，在 if 的同一行进行错误检查，是 Go 代码中的一种常见用法。
		for _, i := range []int{7, 42} {
			if r, e := f1(i); e != nil {
				fmt.Println("f1 failed:", e)
			} else {
				fmt.Println("f1 worked:", r)
			}
		}
		for _, i := range []int{7, 42} {
			if r, e := f2(i); e != nil {
				fmt.Println("f2 failed:", e)
			} else {
				fmt.Println("f2 worked:", r)
			}
		}

		// If you want to programmatically use the data in a custom error, you’ll need to get the error as an instance of the custom error type via type assertion.
		// 如果你想在程序中使用自定义错误类型的数据， 你需要通过类型断言来得到这个自定义错误类型的实例。
		_, e := f2(42)
		if ae, ok := e.(*argError); ok {
			fmt.Println(ae.arg)
			fmt.Println(ae.prob)
		}

		// See this great post on the Go blog for more on error handling.
		// 到 Go 官方博客去看这篇很棒的文章， 以获取更多关于错误处理的信息。
		// http://blog.golang.org/2011/07/error-handling-and-go.html
	},
}

func init() {
	cmd.RootCmd.AddCommand(errorsCmd)
}
