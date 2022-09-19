package cmd

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "shell查询交互模式",
	Run: func(cmd *cobra.Command, args []string) {
		p := prompt.New(
			executor,
			completer,
			prompt.OptionPrefix(">>> "),
		)
		p.Run()
	},
}

func executor(s string) {
	s = strings.TrimSpace(s)

	switch s {
	case "search":
		return
	case "exit", "quit":
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	args := strings.Split(s, " ")
	if len(args) > 2 {
		color.Errorln("错误，参数太多")
		return
	}

	//fmt.Println(args)

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("/bin/sh", "-c", ex+" "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}

	return

}

func completer(in prompt.Document) []prompt.Suggest {

	if in.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}

	var promptSuggests []prompt.Suggest
	for _, cmd := range rootCmd.Commands() {
		switch cmd.Name() {
		case "help", "completion":
			continue
		}
		promptSuggests = append(promptSuggests, prompt.Suggest{Text: cmd.Use, Description: cmd.Short})
	}

	args := strings.Split(in.TextBeforeCursor(), " ")
	var secondLevelCommand = []string{args[0]}
	secondLevelCmd, _, secondLevelErr := rootCmd.Find(secondLevelCommand)
	if secondLevelErr != nil {
		return prompt.FilterHasPrefix(promptSuggests, in.GetWordBeforeCursor(), true)
	}

	argsCount := len(args)
	switch argsCount {
	case 1, 2:
		promptSuggests = []prompt.Suggest{}
		for _, cmd := range secondLevelCmd.Commands() {
			promptSuggests = append(promptSuggests, prompt.Suggest{Text: cmd.Use, Description: cmd.Short})
		}

		if argsCount == 2 {
			var threeLevelCommand = []string{args[1]}
			threeLevelCmd, _, _ := secondLevelCmd.Find(threeLevelCommand)
			if threeLevelCmd.Use == args[1] {
				promptSuggests = []prompt.Suggest{}
			}
		}

	default:
		promptSuggests = []prompt.Suggest{}
	}

	return prompt.FilterHasPrefix(promptSuggests, in.GetWordBeforeCursor(), true)
}
