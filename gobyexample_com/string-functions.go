package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	s "strings"
)

// The standard library’s strings package provides many useful string-related functions. Here are some examples to give you a sense of the package.
// 标准库的 strings 包提供了很多有用的字符串相关的函数。 这儿有一些用来让你对 strings 包有一个初步了解的例子。

// We alias fmt.Println to a shorter name as we’ll use it a lot below.
// 我们给 fmt.Println 一个较短的别名， 因为我们随后会大量的使用它。
var p = fmt.Println

// stringFunctionsCmd represents the string functions command
var stringFunctionsCmd = &cobra.Command{
	Use:   "go-by-example:string-functions",
	Short: "https://gobyexample.com/string-functions",
	Run: func(cmd *cobra.Command, args []string) {

		// Here’s a sample of the functions available in strings. Since these are functions from the package, not methods on the string object itself, we need pass the string in question as the first argument to the function. You can find more functions in the strings package docs.
		// 这是一些 strings 中有用的函数例子。 由于它们都是包的函数，而不是字符串对象自身的方法， 这意味着我们需要在调用函数时，将字符串作为第一个参数进行传递。 你可以在 strings 包文档中找到更多的函数。

		p("Contains:  ", s.Contains("test", "es"))
		p("Count:     ", s.Count("test", "t"))
		p("HasPrefix: ", s.HasPrefix("test", "te"))
		p("HasSuffix: ", s.HasSuffix("test", "st"))
		p("Index:     ", s.Index("test", "e"))
		p("Join:      ", s.Join([]string{"a", "b"}, "-"))
		p("Repeat:    ", s.Repeat("a", 5))
		p("Replace:   ", s.Replace("foo", "o", "0", -1))
		p("Replace:   ", s.Replace("foo", "o", "0", 1))
		p("Split:     ", s.Split("a-b-c-d-e", "-"))
		p("ToLower:   ", s.ToLower("TEST"))
		p("ToUpper:   ", s.ToUpper("test"))
		p()

		// Not part of strings, but worth mentioning here, are the mechanisms for getting the length of a string in bytes and getting a byte by index.
		// 虽然不是 strings 的函数，但仍然值得一提的是， 获取字符串长度（以字节为单位）以及通过索引获取一个字节的机制。

		p("Len: ", len("hello"))
		p("Char:", "hello"[1])

		// Note that len and indexing above work at the byte level. Go uses UTF-8 encoded strings, so this is often useful as-is. If you’re working with potentially multi-byte characters you’ll want to use encoding-aware operations. See strings, bytes, runes and characters in Go for more information.
		// 注意，上面的 len 以及索引工作在字节级别上。 Go 使用 UTF-8 编码字符串，因此通常按原样使用。 如果您可能使用多字节的字符，则需要使用可识别编码的操作。 详情请参考 strings, bytes, runes and characters in Go。

	},
}

func init() {
	cmd.RootCmd.AddCommand(stringFunctionsCmd)
}
