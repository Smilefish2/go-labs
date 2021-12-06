package gobyexample_com

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"os"
)

// Writing files in Go follows similar patterns to the ones we saw earlier for reading.
// 在 Go 中，写文件与我们前面看过的读文件方法类似。

// writingFilesCmd represents the writing-files command
var writingFilesCmd = &cobra.Command{
	Use:   "go-by-example:writing-files",
	Short: "https://gobyexample.com/writing-files",
	Run: func(cmd *cobra.Command, args []string) {

		// To start, here’s how to dump a string (or just bytes) into a file.
		// 开始！这里展示了如何写入一个字符串（或者只是一些字节）到一个文件。
		d1 := []byte("hello\ngo\n")
		err := os.WriteFile("/tmp/dat1", d1, 0644)
		check(err)

		// For more granular writes, open a file for writing.
		// 对于更细粒度的写入，先打开一个文件。
		f, err := os.Create("/tmp/dat2")
		check(err)

		// It’s idiomatic to defer a Close immediately after opening a file.
		// 打开文件后，一个习惯性的操作是：立即使用 defer 调用文件的 Close。
		defer f.Close()

		// You can Write byte slices as you’d expect.
		// 您可以按期望的那样 Write 字节切片。
		d2 := []byte{115, 111, 109, 101, 10}
		n2, err := f.Write(d2)
		check(err)
		fmt.Printf("wrote %d bytes\n", n2)

		// A WriteString is also available.
		// WriteString 也是可用的。
		n3, err := f.WriteString("writes\n")
		check(err)
		fmt.Printf("wrote %d bytes\n", n3)

		// Issue a Sync to flush writes to stable storage.
		// 调用 Sync 将缓冲区的数据写入硬盘。
		f.Sync()

		// bufio provides buffered writers in addition to the buffered readers we saw earlier.
		// 与我们前面看到的带缓冲的 Reader 一样，bufio 还提供了的带缓冲的 Writer。
		w := bufio.NewWriter(f)
		n4, err := w.WriteString("buffered\n")
		check(err)
		fmt.Printf("wrote %d bytes\n", n4)

		// Use Flush to ensure all buffered operations have been applied to the underlying writer.
		// 使用 Flush 来确保，已将所有的缓冲操作应用于底层 writer。
		w.Flush()

		// Try running the file-writing code.
		// 运行这段文件写入代码。

		// Then check the contents of the written files.
		// 然后检查写入文件的内容。

		// Next we’ll look at applying some of the file I/O ideas we’ve just seen to the stdin and stdout streams.
		// 我们刚刚看到了文件 I/O 思想， 接下来，我们看看它在 stdin 和 stdout 流中的应用。
	},
}

func init() {
	cmd.RootCmd.AddCommand(writingFilesCmd)
}
