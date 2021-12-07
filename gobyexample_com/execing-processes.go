package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

// In the previous example we looked at spawning external processes. We do this when we need an external process accessible to a running Go process. Sometimes we just want to completely replace the current Go process with another (perhaps non-Go) one. To do this we’ll use Go’s implementation of the classic exec function.
// 在前面的例子中，我们了解了生成外部进程的知识， 当我们需要在运行的 Go 流程中访问的外部流程时，便可以执行此操作。 但是有时候，我们只想用其它（也许是非 Go）的进程，来完全替代当前的 Go 进程。 这时，我们可以使用经典的 exec 函数的 Go 的实现。

// execingProcessesCmd represents the execing-processes command
var execingProcessesCmd = &cobra.Command{
	Use:   "go-by-example:execing-processes",
	Short: "https://gobyexample.com/execing-processes",
	Run: func(cmd *cobra.Command, args []string) {

		// For our example we’ll exec ls. Go requires an absolute path to the binary we want to execute, so we’ll use exec.LookPath to find it (probably /bin/ls).
		// 在这个例子中，我们将执行 ls 命令。 Go 要求我们提供想要执行的可执行文件的绝对路径， 所以我们将使用 exec.LookPath 找到它（应该是 /bin/ls）。
		binary, lookErr := exec.LookPath("ls")
		if lookErr != nil {
			panic(lookErr)
		}

		// Exec requires arguments in slice form (as opposed to one big string). We’ll give ls a few common arguments. Note that the first argument should be the program name.
		// Exec 需要的参数是切片的形式的（不是放在一起的一个大字符串）。 我们给 ls 一些基本的参数。注意，第一个参数需要是程序名。
		args2 := []string{"ls", "-a", "-l", "-h"}

		// Exec also needs a set of environment variables to use. Here we just provide our current environment.
		// Exec 同样需要使用环境变量。 这里我们仅提供当前的环境变量。
		env := os.Environ()

		// Here’s the actual syscall.Exec call. If this call is successful, the execution of our process will end here and be replaced by the /bin/ls -a -l -h process. If there is an error we’ll get a return value.
		// 这里是真正的 syscall.Exec 调用。 如果这个调用成功，那么我们的进程将在这里结束，并被 /bin/ls -a -l -h 进程代替。 如果存在错误，那么我们将会得到一个返回值。
		execErr := syscall.Exec(binary, args2, env)
		if execErr != nil {
			panic(execErr)
		}

		// When we run our program it is replaced by ls.
		// 当我们运行程序时，它会替换为 ls。
		fmt.Println("go run execing-processes.go")

		// Note that Go does not offer a classic Unix fork function. Usually this isn’t an issue though, since starting goroutines, spawning processes, and exec’ing processes covers most use cases for fork.
		// 注意 Go 没有提供 Unix 经典的 fork 函数。 一般来说，这没有问题，因为启动协程、生成进程和执行进程， 已经涵盖了 fork 的大多数使用场景。

	},
}

func init() {
	cmd.RootCmd.AddCommand(execingProcessesCmd)
}
