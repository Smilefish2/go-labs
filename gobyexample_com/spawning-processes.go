package gobyexample_com

import (
	"fmt"
	"io"
	"learn-go-with-cli/cmd"
	"os/exec"

	"github.com/spf13/cobra"
)

// Sometimes our Go programs need to spawn other, non-Go processes.
// 有时，我们的 Go 程序需要生成其他的、非 Go 的进程。

// spawningProcessesCmd represents the spawning-processes command
var spawningProcessesCmd = &cobra.Command{
	Use:   "go-by-example:spawning-processes",
	Short: "https://gobyexample.com/spawning-processes",
	Run: func(cmd *cobra.Command, args []string) {

		// We’ll start with a simple command that takes no arguments or input and just prints something to stdout. The exec.Command helper creates an object to represent this external process.
		// 我们将从一个简单的命令开始，没有参数或者输入，仅打印一些信息到标准输出流。 exec.Command 可以帮助我们创建一个对象，来表示这个外部进程。
		dateCmd := exec.Command("date")

		// .Output is another helper that handles the common case of running a command, waiting for it to finish, and collecting its output. If there were no errors, dateOut will hold bytes with the date info.
		// .Output 是另一个帮助函数，常用于处理运行命令、等待命令完成并收集其输出。 如果没有错误，dateOut 将保存带有日期信息的字节。
		dateOut, err := dateCmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println("> date")
		fmt.Println(string(dateOut))

		// Next we’ll look at a slightly more involved case where we pipe data to the external process on its stdin and collect the results from its stdout.
		// 下面我们将看看一个稍复杂的例子， 我们将从外部进程的 stdin 输入数据并从 stdout 收集结果。
		grepCmd := exec.Command("grep", "hello")

		// Here we explicitly grab input/output pipes, start the process, write some input to it, read the resulting output, and finally wait for the process to exit.
		// 这里我们明确的获取输入/输出管道，运行这个进程， 写入一些输入数据、读取输出结果，最后等待程序运行结束。
		grepIn, _ := grepCmd.StdinPipe()
		grepOut, _ := grepCmd.StdoutPipe()
		grepCmd.Start()
		grepIn.Write([]byte("hello grep\ngoodbye grep"))
		grepIn.Close()
		grepBytes, _ := io.ReadAll(grepOut)
		grepCmd.Wait()

		// We omitted error checks in the above example, but you could use the usual if err != nil pattern for all of them. We also only collect the StdoutPipe results, but you could collect the StderrPipe in exactly the same way.
		// 上面的例子中，我们忽略了错误检测， 当然，你也可以使用常见的 if err != nil 方式来进行错误检查。 我们只收集了 StdoutPipe 的结果， 但是你可以使用相同的方法收集 StderrPipe 的结果。
		fmt.Println("> grep hello")
		fmt.Println(string(grepBytes))

		// Note that when spawning commands we need to provide an explicitly delineated command and argument array, vs. being able to just pass in one command-line string. If you want to spawn a full command with a string, you can use bash’s -c option:
		// 注意，在生成命令时，我们需要提供一个明确描述命令和参数的数组，而不能只传递一个命令行字符串。 如果你想使用一个字符串生成一个完整的命令，那么你可以使用 bash 命令的 -c 选项：
		lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
		lsOut, err := lsCmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println("> ls -a -l -h")
		fmt.Println(string(lsOut))

		// The spawned programs return output that is the same as if we had run them directly from the command-line.
		// 生成的程序返回的输出，和我们直接通过命令行运行这些程序的输出是相同的。
	},
}

func init() {
	cmd.RootCmd.AddCommand(spawningProcessesCmd)
}
