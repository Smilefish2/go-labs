package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"os"
)

// Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
// Defer 用于确保程序在执行完成后，会调用某个函数，一般是执行清理工作。 Defer 的用途跟其他语言的 ensure 或 finally 类似。

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

// It’s important to check for errors when closing a file, even in a deferred function.
// 关闭文件时，进行错误检查是非常重要的， 即使在 defer 函数中也是如此。
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// deferCmd represents the defer command
var deferCmd = &cobra.Command{
	Use:   "go-by-example:defer",
	Short: "https://gobyexample.com/defer",
	Run: func(cmd *cobra.Command, args []string) {

		// Suppose we wanted to create a file, write to it, and then close when we’re done. Here’s how we could do that with defer.
		// 假设我们想要创建一个文件、然后写入数据、最后在程序结束时关闭该文件。 这里展示了如何通过 defer 来做到这一切。

		// Immediately after getting a file object with createFile, we defer the closing of that file with closeFile. This will be executed at the end of the enclosing function (main), after writeFile has finished.
		// 在 createFile 后立即得到一个文件对象， 我们使用 defer 通过 closeFile 来关闭这个文件。 这会在封闭函数（main）结束时执行，即 writeFile 完成以后。

		f := createFile("/tmp/defer.txt")
		defer closeFile(f)
		writeFile(f)

		// Running the program confirms that the file is closed after being written.
		// 执行程序，确认写入数据后，文件已关闭。

	},
}

func init() {
	cmd.RootCmd.AddCommand(deferCmd)
}
