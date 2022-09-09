package question

import (
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
)

// 常见语法题目 一

var q007Cmd = &cobra.Command{
	Use:   "question:007",
	Short: "https://github.com/yongxinz/gopher/blob/main/interview/q007.md",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cmd.RootCmd.AddCommand(q007Cmd)
}
