package pointer

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
)

// PointerCmd 从指针获取指针指向的
// 当使用&操作符对普通变量进行取地址操作并得到变量的指针后，可以对指针使用*操作符，也就是指针取值，代码如下。
var pointerCmd = &cobra.Command{
	Use:   "pointer:main",
	Short: "http://c.biancheng.net/view/21.html",
	Run: func(cmd *cobra.Command, args []string) {

		// 准备一个字符串类型
		var house = "Malibu Point 10880, 90265"
		// 对字符串取地址, ptr类型为*string
		ptr := &house
		// 打印ptr的类型
		fmt.Printf("ptr type: %T\n", ptr)
		// 打印ptr的指针地址
		fmt.Printf("address: %p\n", ptr)
		// 对指针进行取值操作
		value := *ptr
		// 取值后的类型
		fmt.Printf("value type: %T\n", value)
		// 指针取值后就是指向变量的值
		fmt.Printf("value: %s\n", value)
	},
}

func init() {
	cmd.RootCmd.AddCommand(pointerCmd)
}
