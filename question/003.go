package question

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
)

// 翻转字符串

// 问题描述
// 请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
// 给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

// 解题思路
// 翻转字符串其实是将一个字符串以中间字符为轴，前后翻转，即将str[len]赋值给str[0],将str[0] 赋值 str[len]。

// 源码参考
func reverString(s string) (string, bool) {
	str := []rune(s)
	l := len(str)
	if l > 5000 {
		return s, false
	}

	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	return string(str), true
}

var q003Cmd = &cobra.Command{
	Use:   "question:003",
	Short: "https://github.com/yongxinz/gopher/blob/main/interview/q003.md",
	Run: func(cmd *cobra.Command, args []string) {
		str := "12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728"
		fmt.Println("输入字符：" + str)
		s, result := reverString(str)
		if result {
			fmt.Println("输出字符：" + s)
		} else {
			fmt.Println("翻转失败")
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(q003Cmd)
}
