package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
	"net"
	"net/url"
)

// URLs provide a uniform way to locate resources. Here’s how to parse URLs in Go.
// URL 提供了统一资源定位方式。 这里展示了在 Go 中是如何解析 URL 的。

// urlParsingCmd represents the url-parsing command
var urlParsingCmd = &cobra.Command{
	Use:   "go-by-example:url-parsing",
	Short: "https://gobyexample.com/url-parsing",
	Run: func(cmd *cobra.Command, args []string) {

		// We’ll parse this example URL, which includes a scheme, authentication info, host, port, path, query params, and query fragment.
		// 我们将解析这个 URL 示例，它包含了一个 scheme、 认证信息、主机名、端口、路径、查询参数以及查询片段。
		s := "postgres://user:pass@host.com:5432/path?k=v#f"

		// Parse the URL and ensure there are no errors.
		// 解析这个 URL 并确保解析没有出错。
		u, err := url.Parse(s)
		if err != nil {
			panic(err)
		}

		// Accessing the scheme is straightforward.
		// 直接访问 scheme。
		fmt.Println(u.Scheme)

		// User contains all authentication info; call Username and Password on this for individual values.
		// User 包含了所有的认证信息， 这里调用 Username 和 Password 来获取单独的值。
		fmt.Println(u.User)
		fmt.Println(u.User.Username())
		p, _ := u.User.Password()
		fmt.Println(p)

		// The Host contains both the hostname and the port, if present. Use SplitHostPort to extract them.
		// Host 包含主机名和端口号（如果存在）。使用 SplitHostPort 提取它们。
		fmt.Println(u.Host)
		host, port, _ := net.SplitHostPort(u.Host)
		fmt.Println(host)
		fmt.Println(port)

		// Here we extract the path and the fragment after the #.
		// 这里我们提取路径和 # 后面的查询片段信息。
		fmt.Println(u.Path)
		fmt.Println(u.Fragment)

		// To get query params in a string of k=v format, use RawQuery. You can also parse query params into a map. The parsed query param maps are from strings to slices of strings, so index into [0] if you only want the first value.
		// 要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。 你也可以将查询参数解析为一个 map。已解析的查询参数 map 以查询字符串为键， 已解析的查询参数会从字符串映射到到字符串的切片， 因此如果您只想要第一个值，则索引为 [0]。
		fmt.Println(u.RawQuery)
		m, _ := url.ParseQuery(u.RawQuery)
		fmt.Println(m)
		fmt.Println(m["k"][0])

		// Running our URL parsing program shows all the different pieces that we extracted.
		// 运行我们的 URL 解析程序， 显示全部我们提取的 URL 的不同数据块。
	},
}

func init() {
	cmd.RootCmd.AddCommand(urlParsingCmd)
}
