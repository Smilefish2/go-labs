package question

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
	"strings"
)

// 判断字符串中字符是否全都不同

// 问题描述
// 请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
// 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。

// 解题思路
// 这里有几个重点，第一个是ASCII字符，ASCII字符字符一共有256个，其中128个是常用字符，可以在键盘上输入。128之后的是键盘上无法找到的。
// 然后是全部不同，也就是字符串中的字符没有重复的，再次，不准使用额外的储存结构，且字符串小于等于3000。
// 如果允许其他额外储存结构，这个题目很好做。如果不允许的话，可以使用golang内置的方式实现。

// 源码参考
// 通过strings.Count 函数判断：
func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}

	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

// 通过strings.Index和strings.LastIndex函数判断：
func isUniqueString2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}

	for k, v := range s {
		if v > 127 {
			return false
		}

		if strings.Index(s, string(v)) != k {
			return false
		}
	}

	return true
}

//源码解析
//
// 以上两种方法都可以实现这个算法。
// 第一个方法使用的是golang内置方法strings.Count,可以用来判断在一个字符串中包含的另外一个字符串的数量。
// 第二个方法使用的是golang内置方法strings.Index和strings.LastIndex，用来判断指定字符串在另外一个字符串的索引位置，分别是第一次发现位置和最后发现位置。

var q002Cmd = &cobra.Command{
	Use:   "question:002",
	Short: "https://github.com/yongxinz/gopher/blob/main/interview/q002.md",
	Run: func(cmd *cobra.Command, args []string) {
		strUnique := "12AB34CD56EF78GH"
		strRepeat := "12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728"

		fmt.Println(isUniqueString(strUnique))
		fmt.Println(isUniqueString(strRepeat))

		fmt.Println(isUniqueString2(strUnique))
		fmt.Println(isUniqueString2(strRepeat))
	},
}

func init() {
	cmd.RootCmd.AddCommand(q002Cmd)
}
