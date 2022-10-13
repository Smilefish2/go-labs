package gobyexample_com

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
	"regexp"
)

// Go offers built-in support for regular expressions. Here are some examples of common regexp-related tasks in Go.
// Go 提供了内建的正则表达式支持。 这儿有一些在 Go 中与 regexp 相关的常见用法示例。

// regularExpressionsCmd represents the string formatting command
var regularExpressionsCmd = &cobra.Command{
	Use:   "go-by-example:regular-expressions",
	Short: "https://gobyexample.com/regular-expressions",
	Run: func(cmd *cobra.Command, args []string) {

		// This tests whether a pattern matches a string.
		// 测试一个字符串是否符合一个表达式。
		match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
		fmt.Println(match)

		// Above we used a string pattern directly, but for other regexp tasks you’ll need to Compile an optimized Regexp struct.
		// 上面我们是直接使用字符串，但是对于一些其他的正则任务， 你需要通过 Compile 得到一个优化过的 Regexp 结构体。
		r, _ := regexp.Compile("p([a-z]+)ch")

		// Many methods are available on these structs. Here’s a match test like we saw earlier.
		// 该结构体有很多方法。这是一个类似于我们前面看到的匹配测试。
		fmt.Println(r.MatchString("peach"))

		// This finds the match for the regexp.
		// 查找匹配的字符串。
		fmt.Println(r.FindString("peach punch"))
		// This also finds the first match but returns the start and end indexes for the match instead of the matching text.
		// 这个也是查找首次匹配的字符串， 但是它的返回值是，匹配开始和结束位置的索引，而不是匹配的内容。
		fmt.Println("idx:", r.FindStringIndex("peach punch"))
		// The Submatch variants include information about both the whole-pattern matches and the submatches within those matches. For example this will return information for both p([a-z]+)ch and ([a-z]+).
		// Submatch 返回完全匹配和局部匹配的字符串。 例如，这里会返回 p([a-z]+)ch 和 ([a-z]+) 的信息。
		fmt.Println(r.FindStringSubmatch("peach punch"))

		// Similarly this will return information about the indexes of matches and submatches.
		// 类似的，这个会返回完全匹配和局部匹配位置的索引。
		fmt.Println(r.FindStringSubmatchIndex("peach punch"))
		// The All variants of these functions apply to all matches in the input, not just the first. For example to find all matches for a regexp.
		// 带 All 的这些函数返回全部的匹配项， 而不仅仅是首次匹配项。例如查找匹配表达式全部的项。
		fmt.Println(r.FindAllString("peach punch pinch", -1))
		// These All variants are available for the other functions we saw above as well.
		// All 同样可以对应到上面的所有函数。
		fmt.Println("all:", r.FindAllStringSubmatchIndex(
			"peach punch pinch", -1))

		// Providing a non-negative integer as the second argument to these functions will limit the number of matches.
		// 这些函数接收一个正整数作为第二个参数，来限制匹配次数。
		fmt.Println(r.FindAllString("peach punch pinch", 2))

		// Our examples above had string arguments and used names like MatchString. We can also provide []byte arguments and drop String from the function name.
		// 上面的例子中，我们使用了字符串作为参数， 并使用了 MatchString 之类的方法。 我们也可以将 String 从函数命中去掉，并提供 []byte 的参数。
		fmt.Println(r.Match([]byte("peach")))

		// When creating global variables with regular expressions you can use the MustCompile variation of Compile. MustCompile panics instead of returning an error, which makes it safer to use for global variables.
		// 创建正则表达式的全局变量时，可以使用 Compile 的变体 MustCompile 。 MustCompile 用 panic 代替返回一个错误 ，这样使用全局变量更加安全。
		r = regexp.MustCompile("p([a-z]+)ch")
		fmt.Println("regexp:", r)

		// The regexp package can also be used to replace subsets of strings with other values.
		// regexp 包也可以用来将子字符串替换为为其它值。

		fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

		// The Func variant allows you to transform matched text with a given function.
		// Func 变体允许您使用给定的函数来转换匹配的文本。

		in := []byte("a peach")
		out := r.ReplaceAllFunc(in, bytes.ToUpper)
		fmt.Println(string(out))

		// For a complete reference on Go regular expressions check the regexp package docs.
		// 有关 Go 正则表达式的说明，请参考 regexp 包文档。
	},
}

func init() {
	cmd.RootCmd.AddCommand(regularExpressionsCmd)
}
