package question

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
)

var q008Cmd = &cobra.Command{
	Use:   "question:008",
	Short: "https://github.com/yongxinz/gopher/blob/main/interview/q008.md",
	Run: func(cmd *cobra.Command, args []string) {
		s := make([]int, 5)
		s = append(s, 1, 2, 3)
		fmt.Println(s)
	},
}

func init() {
	cmd.RootCmd.AddCommand(q008Cmd)
}
