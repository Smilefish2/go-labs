package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"os/exec"
)

// Unit testing is an important part of writing principled Go programs. The testing package provides the tools we need to write unit tests and the go test command runs tests.
// 想要写出好的 Go 程序，单元测试是很重要的一部分。 testing 包为提供了编写单元测试所需的工具，写好单元测试后，我们可以通过 go test 命令运行测试。

// For the sake of demonstration, this code is in package main, but it could be any package. Testing code typically lives in the same package as the code it tests.
// 为方便演示，例子的代码位于 main 包，实际上，单元测试的代码可以位于任何包下。 测试代码通常与需要被测试的代码位于同一个包下。

// testingCmd represents the testing command
var testingCmd = &cobra.Command{
	Use:   "go-by-example:testing",
	Short: "https://gobyexample.com/testing",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("run: cd gobyexample_com")
		fmt.Println("run: go test -v")
		fmt.Println("run: go test -bench=.")
		fmt.Println("run: cd ..")

		// Run all tests in the current project in verbose mode.
		// 以详细模式运行当前项目中的所有测试。

		// Run all benchmarks in the current project. All tests are run prior to benchmarks. The bench flag filters benchmark function names with a regexp.
		// 运行当前项目中的所有基准测试。所有测试都在基准测试之前运行。bench标志使用regexp筛选基准函数名。

		cmdShell := "cd gobyexample_com && go test -v && go test -bench=. && cd .."
		out, err := exec.Command("bash", "-c", cmdShell).Output()
		if err != nil {
			fmt.Printf("Failed to execute command: %s", cmdShell)
		}
		fmt.Println(string(out))
	},
}

func init() {
	cmd.RootCmd.AddCommand(testingCmd)
}
