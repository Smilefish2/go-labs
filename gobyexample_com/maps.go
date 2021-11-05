package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// mapsCmd represents the maps command
var mapsCmd = &cobra.Command{
	Use:   "go-by-example:maps",
	Short: "https://gobyexample.com/maps",
	Run: func(cmd *cobra.Command, args []string) {

		// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
		// map 是 Go 内建的关联数据类型 （在一些其他的语言中也被称为 哈希(hash) 或者 字典(dict) ）。

		// To create an empty map, use the builtin make: make(map[key-type]val-type).
		// 要创建一个空 map，需要使用内建函数 make：make(map[key-type]val-type)。
		m := make(map[string]int)

		// Set key/value pairs using typical name[key] = val syntax.
		// 使用典型的 name[key] = val 语法来设置键值对。
		m["k1"] = 7
		m["k2"] = 13

		// Printing a map with e.g. fmt.Println will show all of its key/value pairs.
		// 打印 map。例如，使用 fmt.Println 打印一个 map，会输出它所有的键值对。
		fmt.Println("map:", m)

		// Get a value for a key with name[key].
		// 使用 name[key] 来获取一个键的值。
		v1 := m["k1"]
		fmt.Println("v1: ", v1)

		// The builtin len returns the number of key/value pairs when called on a map.
		// 内建函数 len 可以返回一个 map 的键值对数量。
		fmt.Println("len:", len(m))

		// The builtin delete removes key/value pairs from a map.
		// 内建函数 delete 可以从一个 map 中移除键值对。
		delete(m, "k2")
		fmt.Println("map:", m)

		// The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.
		// 当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。 这可以用来消除 键不存在 和 键的值为零值 产生的歧义， 例如 0 和 ""。这里我们不需要值，所以用 空白标识符(blank identifier) _ 将其忽略。
		_, prs := m["k2"]
		fmt.Println("prs:", prs)

		// You can also declare and initialize a new map in the same line with this syntax.
		// 你也可以通过右边的语法在一行代码中声明并初始化一个新的 map。
		n := map[string]int{"foo": 1, "bar": 2}
		fmt.Println("map:", n)

		// Note that maps appear in the form map[k:v k:v] when printed with fmt.Println.
		// 注意，使用 fmt.Println 打印一个 map 的时候， 是以 map[k:v k:v] 的格式输出的。
	},
}

func init() {
	cmd.RootCmd.AddCommand(mapsCmd)
}
