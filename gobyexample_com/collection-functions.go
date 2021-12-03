package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"strings"
)

// We often need our programs to perform operations on collections of data, like selecting all items that satisfy a given predicate or mapping all items to a new collection with a custom function.
// 我们经常需要程序对数据集合执行操作， 例如选择满足给定条件的全部 item， 或通过自定义函数将全部 item 映射到一个新的集合。

// In some languages it’s idiomatic to use generic data structures and algorithms. Go does not support generics; in Go it’s common to provide collection functions if and when they are specifically needed for your program and data types.
// 在某些语言中，习惯使用通用数据结构和算法。Go不支持泛型；在Go中，如果程序和数据类型特别需要收集函数，则通常会提供这些函数。

// Here are some example collection functions for slices of strings. You can use these examples to build your own functions. Note that in some cases it may be clearest to just inline the collection-manipulating code directly, instead of creating and calling a helper function.
// 下面是一些用于字符串片段的示例集合函数。您可以使用这些示例构建自己的函数。请注意，在某些情况下，直接内联集合操作代码可能是最清楚的，而不是创建和调用助手函数。

// Index returns the first index of the target string t, or -1 if no match is found.
// Index 返回目标字符串 t 在 vs 中第一次出现位置的索引， 或者在没有匹配值时返回 -1。
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include returns true if the target string t is in the slice.
// Include 如果目标字符串 t 存在于切片 vs 中，则返回 true。
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any returns true if one of the strings in the slice satisfies the predicate f.
// Any 如果切片 vs 中的任意一个字符串满足条件 f，则返回 true。
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all of the strings in the slice satisfy the predicate f.
// All 如果切片 vs 中的所有字符串都满足条件 f，则返回 true。
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter 返回一个新的切片，新切片由原切片 vs 中满足条件 f 的字符串构成。
// Filter returns a new slice containing all strings in the slice that satisfy the predicate f.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying the function f to each string in the original slice.
// Map 返回一个新的切片，新切片的字符串由原切片 vs 中的字符串经过函数 f 映射后得到。
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// collectionFunctionsCmd represents the collection functions command
var collectionFunctionsCmd = &cobra.Command{
	Use:   "go-by-example:collection-functions",
	Short: "https://gobyexample.com/collection-functions",
	Run: func(cmd *cobra.Command, args []string) {

		// 试试各种组合函数。
		var strs = []string{"peach", "apple", "pear", "plum"}

		fmt.Println(Index(strs, "pear"))

		fmt.Println(Include(strs, "grape"))

		fmt.Println(Any(strs, func(v string) bool {
			return strings.HasPrefix(v, "p")
		}))

		fmt.Println(All(strs, func(v string) bool {
			return strings.HasPrefix(v, "p")
		}))

		fmt.Println(Filter(strs, func(v string) bool {
			return strings.Contains(v, "e")
		}))

		fmt.Println(Map(strs, strings.ToUpper))

		// 上面的例子都是用的匿名函数，当前，你也可以使用正确类型的命名函数
	},
}

func init() {
	cmd.RootCmd.AddCommand(collectionFunctionsCmd)
}
