package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"strconv"
)

// Parsing numbers from strings is a basic but common task in many programs; here’s how to do it in Go.
// 从字符串中解析数字在很多程序中是一个基础常见的任务， 而在 Go 中，是这样处理的。

// The built-in package strconv provides the number parsing.
// 内建的 strconv 包提供了数字解析能力。

// numberParsingCmd represents the number-parsing command
var numberParsingCmd = &cobra.Command{
	Use:   "go-by-example:for",
	Short: "https://gobyexample.com/for",
	Run: func(cmd *cobra.Command, args []string) {

		// With ParseFloat, this 64 tells how many bits of precision to parse.
		// 使用 ParseFloat，这里的 64 表示解析的数的位数。

		f, _ := strconv.ParseFloat("1.234", 64)
		fmt.Println(f)

		// For ParseInt, the 0 means infer the base from the string. 64 requires that the result fit in 64 bits.
		// 在使用 ParseInt 解析整型数时， 例子中的参数 0 表示自动推断字符串所表示的数字的进制。 64 表示返回的整型数是以 64 位存储的。
		i, _ := strconv.ParseInt("123", 0, 64)
		fmt.Println(i)

		// ParseInt will recognize hex-formatted numbers.
		// ParseInt 会自动识别出字符串是十六进制数。

		d, _ := strconv.ParseInt("0x1c8", 0, 64)
		fmt.Println(d)

		// A ParseUint is also available.
		// ParseUint 也是可用的。

		u, _ := strconv.ParseUint("789", 0, 64)
		fmt.Println(u)

		// Atoi is a convenience function for basic base-10 int parsing.
		// Atoi 是一个基础的 10 进制整型数转换函数。

		k, _ := strconv.Atoi("135")
		fmt.Println(k)

		// Parse functions return an error on bad input.
		// 在输入错误时，解析函数会返回一个错误。
		_, e := strconv.Atoi("wat")
		fmt.Println(e)

		// Next we’ll look at another common parsing task: URLs.
		// 下面我们将了解一下另一个常见的解析任务：URL 解析。
	},
}

func init() {
	cmd.RootCmd.AddCommand(numberParsingCmd)
}
