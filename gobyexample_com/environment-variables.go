package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
	"os"
	"strings"
)

// Environment variables are a universal mechanism for conveying configuration information to Unix programs. Let’s look at how to set, get, and list environment variables.
// 环境变量 是一种向 Unix 程序传递配置信息的常见方式。 让我们来看看如何设置、获取以及列出环境变量。

// environmentVariablesCmd represents the environment-variables command
var environmentVariablesCmd = &cobra.Command{
	Use:   "go-by-example:environment-variables",
	Short: "https://gobyexample.com/environment-variables",
	Run: func(cmd *cobra.Command, args []string) {

		// To set a key/value pair, use os.Setenv. To get a value for a key, use os.Getenv. This will return an empty string if the key isn’t present in the environment.
		// 使用 os.Setenv 来设置一个键值对。 使用 os.Getenv获取一个键对应的值。 如果键不存在，将会返回一个空字符串。
		os.Setenv("FOO", "1")
		fmt.Println("FOO:", os.Getenv("FOO"))
		fmt.Println("BAR:", os.Getenv("BAR"))

		// Use os.Environ to list all key/value pairs in the environment. This returns a slice of strings in the form KEY=value. You can strings.SplitN them to get the key and value. Here we print all the keys.
		// 使用 os.Environ 来列出所有环境变量键值对。 这个函数会返回一个 KEY=value 形式的字符串切片。 你可以使用 strings.SplitN 来得到键和值。这里我们打印所有的键。
		fmt.Println()
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			fmt.Println(pair[0])
		}

		// Running the program shows that we pick up the value for FOO that we set in the program, but that BAR is empty.
		// 运行这个程序，显示我们在程序中设置的 FOO 的值， 然而没有设置的 BAR 是空的。

		// The list of keys in the environment will depend on your particular machine.
		// 键的列表是由你的电脑情况而定的。

		// If we set BAR in the environment first, the running program picks that value up.
		// 如果我们在运行前设置了 BAR 的值，那么运行程序将会获取到这个值。
	},
}

func init() {
	cmd.RootCmd.AddCommand(environmentVariablesCmd)
}
