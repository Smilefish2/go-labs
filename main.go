package main

import (
	"learn-go-with-cli/cmd"
	_ "learn-go-with-cli/gobyexample_com"
	_ "learn-go-with-cli/pointer"
	_ "learn-go-with-cli/question"
)

func main() {
	cmd.Execute()
}
