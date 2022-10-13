package gobyexample_com

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
	"os"
	"strings"
)

// A line filter is a common type of program that reads input on stdin, processes it, and then prints some derived result to stdout. grep and sed are common line filters.
// 行过滤器（line filter） 是一种常见的程序类型， 它读取 stdin 上的输入，对其进行处理，然后将处理结果打印到 stdout。 grep 和 sed 就是常见的行过滤器。

// Here’s an example line filter in Go that writes a capitalized version of all input text. You can use this pattern to write your own Go line filters.
// 这里是一个使用 Go 编写的行过滤器示例，它将所有的输入文字转化为大写的版本。 你可以使用这个模式来写一个你自己的 Go 行过滤器。

// lineFiltersCmd represents the line-filters command
var lineFiltersCmd = &cobra.Command{
	Use:   "go-by-example:line-filters",
	Short: "https://gobyexample.com/line-filters",
	Run: func(cmd *cobra.Command, args []string) {

		// Wrapping the unbuffered os.Stdin with a buffered scanner gives us a convenient Scan method that advances the scanner to the next token; which is the next line in the default scanner.
		// 用带缓冲的 scanner 包装无缓冲的 os.Stdin， 这为我们提供了一种方便的 Scan 方法， 将 scanner 前进到下一个 令牌（默认为：下一行）。

		scanner := bufio.NewScanner(os.Stdin)

		// Text returns the current token, here the next line, from the input.
		// Text 返回当前的 token，这里指的是输入的下一行。
		for scanner.Scan() {

			// Write out the uppercased line.
			// 输出转换为大写后的行。
			ucl := strings.ToUpper(scanner.Text())

			fmt.Println(ucl)
		}

		// Check for errors during Scan. End of file is expected and not reported by Scan as an error.
		// 检查 Scan 的错误。 文件结束符（EOF）是可以接受的，它不会被 Scan 当作一个错误。
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

		// To try out our line filter, first make a file with a few lowercase lines.
		// 试一下我们的行过滤器，首先创建一个有小写行的文件。

		// Then use the line filter to get uppercase lines.
		// 然后使用行过滤器来得到大写的行。
	},
}

func init() {
	cmd.RootCmd.AddCommand(lineFiltersCmd)
}
