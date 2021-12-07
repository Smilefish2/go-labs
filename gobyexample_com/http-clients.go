package gobyexample_com

import (
	"bufio"
	"fmt"
	"learn-go-with-cli/cmd"
	"net/http"

	"github.com/spf13/cobra"
)

// The Go standard library comes with excellent support for HTTP clients and servers in the net/http package. In this example we’ll use it to issue simple HTTP requests.
// Go 标准库的 net/http 包为 HTTP 客户端和服务端提供了出色的支持。 在这个例子中，我们将使用它发送简单的 HTTP 请求。

// httpClientsCmd represents the http-clients command
var httpClientsCmd = &cobra.Command{
	Use:   "go-by-example:http-clients",
	Short: "https://gobyexample.com/http-clients",
	Run: func(cmd *cobra.Command, args []string) {

		// Issue an HTTP GET request to a server. http.Get is a convenient shortcut around creating an http.Client object and calling its Get method; it uses the http.DefaultClient object which has useful default settings.
		// 向服务端发送一个 HTTP GET 请求。 http.Get 是创建 http.Client 对象并调用其 Get 方法的快捷方式。 它使用了 http.DefaultClient 对象及其默认设置。
		resp, err := http.Get("http://gobyexample.com")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Print the HTTP response status.
		// 打印 HTTP response 状态.
		fmt.Println("Response status:", resp.Status)

		// Print the first 5 lines of the response body.
		// 打印 response body 前面 5 行的内容。
		scanner := bufio.NewScanner(resp.Body)
		for i := 0; scanner.Scan() && i < 5; i++ {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

	},
}

func init() {
	cmd.RootCmd.AddCommand(httpClientsCmd)
}
