package gobyexample_com

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
)

// Go provides built-in support for base64 encoding/decoding.
// Go 提供了对 base64 编解码的内建支持。

// This syntax imports the encoding/base64 package with the b64 name instead of the default base64. It’ll save us some space below.
// 这个语法引入了 encoding/base64 包， 并使用别名 b64 代替默认的 base64。这样可以节省点空间。

// base64EncodingCmd represents the base64-encoding command
var base64EncodingCmd = &cobra.Command{
	Use:   "go-by-example:base64-encoding",
	Short: "https://gobyexample.com/base64-encoding",
	Run: func(cmd *cobra.Command, args []string) {

		// Here’s the string we’ll encode/decode.
		// 这是要编解码的字符串。
		data := "abc123!?$*&()'-=@~"

		// Go supports both standard and URL-compatible base64. Here’s how to encode using the standard encoder. The encoder requires a []byte so we convert our string to that type.
		// Go 同时支持标准 base64 以及 URL 兼容 base64。 这是使用标准编码器进行编码的方法。 编码器需要一个 []byte，因此我们将 string 转换为该类型。
		sEnc := b64.StdEncoding.EncodeToString([]byte(data))
		fmt.Println(sEnc)

		// Decoding may return an error, which you can check if you don’t already know the input to be well-formed.
		// 解码可能会返回错误，如果不确定输入信息格式是否正确， 那么，你就需要进行错误检查了。
		sDec, _ := b64.StdEncoding.DecodeString(sEnc)
		fmt.Println(string(sDec))

		fmt.Println()

		// This encodes/decodes using a URL-compatible base64 format.
		// 使用 URL base64 格式进行编解码。
		uEnc := b64.URLEncoding.EncodeToString([]byte(data))
		fmt.Println(uEnc)
		uDec, _ := b64.URLEncoding.DecodeString(uEnc)
		fmt.Println(string(uDec))

		// The string encodes to slightly different values with the standard and URL base64 encoders (trailing + vs -) but they both decode to the original string as desired.
		// 标准 base64 编码和 URL base64 编码的 编码字符串存在稍许不同（后缀为 + 和 -）， 但是两者都可以正确解码为原始字符串。
	},
}

func init() {
	cmd.RootCmd.AddCommand(base64EncodingCmd)
}
