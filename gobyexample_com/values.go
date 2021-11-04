package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"

	"github.com/spf13/cobra"
)

// valuesCmd represents the values command
var valuesCmd = &cobra.Command{
	Use:   "go-by-example:values",
	Short: "https://gobyexample.com/values",
	Run: func(cmd *cobra.Command, args []string) {

		// 字符串连接
		fmt.Println("go" + "lang")
		fmt.Println("go" + " " + "lang")

		// 加减乘除
		fmt.Println("1+1 =", 1+1)
		fmt.Println("1-1 =", 1-1)
		fmt.Println("7.0*3.0 =", 7.0*3.0)
		fmt.Println("7.0/3.0 =", 7.0/3.0)

		// 布尔
		fmt.Println(true && false)
		fmt.Println(true || false)
		fmt.Println(!true)
	},
}

func init() {
	cmd.RootCmd.AddCommand(valuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
