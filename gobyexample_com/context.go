package gobyexample_com

import (
	"fmt"
	"github.com/fatih/color"
	"learn-go-with-cli/cmd"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// In the previous example we looked at setting up a simple HTTP server. HTTP servers are useful for demonstrating the usage of context.Context for controlling cancellation. A Context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines.
// 在前面的示例中，我们研究了配置简单的 HTTP 服务器。 HTTP 服务器对于演示 context.Context 的用法很有用的， context.Context 被用于控制 cancel。 Context 跨 API 边界和协程携带了：deadline、取消信号以及其他请求范围的值。

func helloContext(w http.ResponseWriter, req *http.Request) {

	// A context.Context is created for each request by the net/http machinery, and is available with the Context() method.
	// net/http 机制为每个请求创建了一个 context.Context， 并且可以通过 Context() 方法获取并使用它。
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Wait for a few seconds before sending a reply to the client. This could simulate some work the server is doing. While working, keep an eye on the context’s Done() channel for a signal that we should cancel the work and return as soon as possible.
	// 等待几秒钟，然后再将回复发送给客户端。 这可以模拟服务器正在执行的某些工作。 在工作时，请密切关注 context 的 Done() 通道的信号， 一旦收到该信号，表明我们应该取消工作并尽快返回。
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		// The context’s Err() method returns an error that explains why the Done() channel was closed.
		// context 的 Err() 方法返回一个错误， 该错误说明了 Done 通道关闭的原因。
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

// contextCmd represents the context command
var contextCmd = &cobra.Command{
	Use:   "go-by-example:context",
	Short: "https://gobyexample.com/context",
	Run: func(cmd *cobra.Command, args []string) {

		color.Yellow("注意⚠️")
		color.Red("使用浏览器访问：http://localhost:8090/hello")
		color.Yellow("注意⚠️")

		// As before, we register our handler on the “/hello” route, and start serving.
		// 跟前面一样，我们在 /hello 路由上注册 handler，然后开始提供服务。

		http.HandleFunc("/hello", helloContext)
		http.ListenAndServe(":8090", nil)

		// Run the server in the background.
		// 后台运行服务器。
		fmt.Println("go run context-in-http-servers.go &")

		// Simulate a client request to /hello, hitting Ctrl+C shortly after starting to signal cancellation.
		// 模拟客户端发出 /hello 请求， 在服务端开始处理后，按下 Ctrl+C 以发出取消信号。
		fmt.Println("curl localhost:8090/hello")
	},
}

func init() {
	cmd.RootCmd.AddCommand(contextCmd)
}
