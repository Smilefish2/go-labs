package gobyexample_com

import (
	"fmt"
	"github.com/fatih/color"
	"go-labs/cmd"
	"net/http"

	"github.com/spf13/cobra"
)

// Writing a basic HTTP server is easy using the net/http package.
// 使用 net/http 包，我们可以轻松实现一个简单的 HTTP 服务器。

// A fundamental concept in net/http servers is handlers. A handler is an object implementing the http.Handler interface. A common way to write a handler is by using the http.HandlerFunc adapter on functions with the appropriate signature.
// handlers 是 net/http 服务器里面的一个基本概念。 handler 对象实现了 http.Handler 接口。 编写 handler 的常见方法是，在具有适当签名的函数上使用 http.HandlerFunc 适配器。
func hello(w http.ResponseWriter, req *http.Request) {

	// Functions serving as handlers take a http.ResponseWriter and a http.Request as arguments. The response writer is used to fill in the HTTP response. Here our simple response is just “hello\n”.
	// handler 函数有两个参数，http.ResponseWriter 和 http.Request。 response writer 被用于写入 HTTP 响应数据，这里我们简单的返回 “hello\n”。
	fmt.Fprintf(w, "hello\n")
}

// This handler does something a little more sophisticated by reading all the HTTP request headers and echoing them into the response body.
// 这个 handler 稍微复杂一点， 我们需要读取的 HTTP 请求 header 中的所有内容，并将他们输出至 response body。
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// httpServersCmd represents the http-servers command
var httpServersCmd = &cobra.Command{
	Use:   "go-by-example:http-servers",
	Short: "https://gobyexample.com/http-servers",
	Run: func(cmd *cobra.Command, args []string) {

		color.Yellow("注意⚠️")
		color.Red("使用浏览器访问：http://localhost:8090/hello")
		color.Yellow("注意⚠️")

		// We register our handlers on server routes using the http.HandleFunc convenience function. It sets up the default router in the net/http package and takes a function as an argument.
		// 使用 http.HandleFunc 函数，可以方便的将我们的 handler 注册到服务器路由。 它是 net/http 包中的默认路由，接受一个函数作为参数。
		http.HandleFunc("/hello", hello)
		http.HandleFunc("/headers", headers)

		// Finally, we call the ListenAndServe with the port and a handler. nil tells it to use the default router we’ve just set up.
		// 最后，我们调用 ListenAndServe 并带上端口和 handler。 nil 表示使用我们刚刚设置的默认路由器。
		http.ListenAndServe(":8090", nil)

		// Run the server in the background.
		// 后台运行服务器。
		fmt.Println("go run http-servers.go &")

		// Access the /hello route.
		// 访问 /hello 路由。
		fmt.Println("curl localhost:8090/hello")

	},
}

func init() {
	cmd.RootCmd.AddCommand(httpServersCmd)
}
