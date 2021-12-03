package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
)

// Go makes it possible to recover from a panic, by using the recover built-in function. A recover can stop a panic from aborting the program and let it continue with execution instead.
// Go通过使用recover内置功能，可以从恐慌中恢复。恢复可以阻止死机中止程序，让程序继续执行。

// An example of where this can be useful: a server wouldn’t want to crash if one of the client connections exhibits a critical error. Instead, the server would want to close that connection and continue serving other clients. In fact, this is what Go’s net/http does by default for HTTP servers.
// 这一点很有用的一个例子是：如果某个客户端连接出现严重错误，服务器不会希望崩溃。相反，服务器希望关闭该连接并继续为其他客户端提供服务。事实上，这就是Go的net/http默认为http服务器所做的。

// This function panics.
// 这一功能令人恐慌。
func mayPanic() {
	panic("a problem")
}

// recoverCmd represents the recover command
var recoverCmd = &cobra.Command{
	Use:   "go-by-example:recover",
	Short: "https://gobyexample.com/recover",
	Run: func(cmd *cobra.Command, args []string) {

		// recover must be called within a deferred function. When the enclosing function panics, the defer will activate and a recover call within it will catch the panic.
		// 必须在延迟函数中调用recover。当封闭函数崩溃时，延迟将被激活，其中的recover调用将捕获崩溃。

		// The return value of recover is the error raised in the call to panic.
		// recover的返回值是在对panic的调用中引发的错误。
		defer func() {
			if r := recover(); r != nil {

				fmt.Println("Recovered. Error:\n", r)
			}
		}()

		mayPanic()

		// This code will not run, because mayPanic panics. The execution of main stops at the point of the panic and resumes in the deferred closure.
		// 这段代码不会运行，因为mayPanic会恐慌。main的执行在恐慌点停止，并在延迟关闭时恢复。
		fmt.Println("After mayPanic()")
	},
}

func init() {
	cmd.RootCmd.AddCommand(recoverCmd)
}
