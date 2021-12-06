package gobyexample_com

import (
	"crypto/sha1"
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
)

// SHA1 hashes are frequently used to compute short identities for binary or text blobs. For example, the git revision control system uses SHA1s extensively to identify versioned files and directories. Here’s how to compute SHA1 hashes in Go.
// SHA1 散列（hash） 经常用于生成二进制文件或者文本块的短标识。 例如，git 版本控制系统 大量的使用了 SHA1 来标识受版本控制的文件和目录。 这是 Go 中如何进行 SHA1 散列计算的例子。

// Go implements several hash functions in various crypto/* packages.
// Go 在多个 crypto/* 包中实现了一系列散列函数。

// sha1HashesCmd represents the sha1-hashes command
var sha1HashesCmd = &cobra.Command{
	Use:   "go-by-example:sha1-hashes",
	Short: "https://gobyexample.com/sha1-hashes",
	Run: func(cmd *cobra.Command, args []string) {

		s := "sha1 this string"

		// The pattern for generating a hash is sha1.New(), sha1.Write(bytes), then sha1.Sum([]byte{}). Here we start with a new hash.
		// 产生一个散列值的方式是 sha1.New()，sha1.Write(bytes)， 然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
		h := sha1.New()

		// Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
		// 写入要处理的字节。如果是一个字符串， 需要使用 []byte(s) 将其强制转换成字节数组。
		h.Write([]byte(s))

		// This gets the finalized hash result as a byte slice. The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.
		// Sum 得到最终的散列值的字符切片。Sum 接收一个参数， 可以用来给现有的字符切片追加额外的字节切片：但是一般都不需要这样做。
		bs := h.Sum(nil)

		// SHA1 values are often printed in hex, for example in git commits. Use the %x format verb to convert a hash results to a hex string.
		// SHA1 值经常以 16 进制输出，例如在 git commit 中。 我们这里也使用 %x 来将散列结果格式化为 16 进制字符串。

		fmt.Println(s)
		fmt.Printf("%x\n", bs)

		// Running the program computes the hash and prints it in a human-readable hex format.
		// 运行程序计算散列值，并以可读的 16 进制格式输出。

		// You can compute other hashes using a similar pattern to the one shown above. For example, to compute MD5 hashes import crypto/md5 and use md5.New().
		// 你可以使用和上面相似的方式来计算其他形式的散列值。 例如，计算 MD5 散列，引入 crypto/md5 并使用 md5.New() 方法。

		// Note that if you need cryptographically secure hashes, you should carefully research hash strength!
		// 注意，如果你需要密码学上的安全散列，你需要仔细的研究一下 加密散列函数。
	},
}

func init() {
	cmd.RootCmd.AddCommand(sha1HashesCmd)
}
