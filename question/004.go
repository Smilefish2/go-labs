package question

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"strings"
)

// 判断两个给定的字符串排序后是否一致

// 问题描述

// 给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
// 这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。
// 保证两串的长度都小于等于5000。

// 解题思路
// 首先要保证字符串长度小于5000。之后只需要一次循环遍历s1中的字符在s2是否都存在即可。

// 源码参考
func isRegroup(s1, s2 string) bool {
	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))

	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}

	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}

	return true
}

// 源码解析
// 这里还是使用golang内置方法 strings.Count 来判断字符是否一致。

var q004Cmd = &cobra.Command{
	Use:   "question:004",
	Short: "https://github.com/yongxinz/gopher/blob/main/interview/q004.md",
	Run: func(cmd *cobra.Command, args []string) {
		s1 := " 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnoqrstuvwxyz"
		s2 := "6832Fjyx7tMEzQbHAwhBoOIRrqlZ 9NsCDGkL5S0gvPWKVXmi4afce1nYJUTud"

		fmt.Println("s1：" + s1)
		fmt.Println("s2：" + s2)

		fmt.Println(isRegroup(s1, s2))
	},
}

func init() {
	cmd.RootCmd.AddCommand(q004Cmd)
}
