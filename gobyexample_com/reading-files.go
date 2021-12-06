package gobyexample_com

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"learn-go-with-cli/cmd"
	"os"
)

// Reading and writing files are basic tasks needed for many Go programs. First we’ll look at some examples of reading files.
// 读写文件在很多程序中都是必须的基本任务。 首先我们来看一些读文件的例子。

// Reading files requires checking most calls for errors. This helper will streamline our error checks below.
// 读取文件需要经常进行错误检查， 这个帮助方法可以精简下面的错误检查过程。
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// readingFilesCmd represents the reading-files command
var readingFilesCmd = &cobra.Command{
	Use:   "go-by-example:reading-files",
	Short: "https://gobyexample.com/reading-files",
	Run: func(cmd *cobra.Command, args []string) {

		// Perhaps the most basic file reading task is slurping a file’s entire contents into memory.
		// 最基本的文件读取任务或许就是将文件内容读取到内存中。
		dat, err := os.ReadFile("/tmp/dat")
		check(err)
		fmt.Print(string(dat))

		// You’ll often want more control over how and what parts of a file are read. For these tasks, start by Opening a file to obtain an os.File value.
		// 您通常会希望对文件的读取方式和内容进行更多控制。 对于这个任务，首先使用 Open 打开一个文件，以获取一个 os.File 值。
		f, err := os.Open("/tmp/dat")
		check(err)

		// Read some bytes from the beginning of the file. Allow up to 5 to be read but also note how many actually were read.
		// 从文件的开始位置读取一些字节。 最多允许读取 5 个字节，但还要注意实际读取了多少个。
		b1 := make([]byte, 5)
		n1, err := f.Read(b1)
		check(err)
		fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

		// You can also Seek to a known location in the file and Read from there.
		// 你也可以 Seek 到一个文件中已知的位置，并从这个位置开始读取。
		o2, err := f.Seek(6, 0)
		check(err)
		b2 := make([]byte, 2)
		n2, err := f.Read(b2)
		check(err)
		fmt.Printf("%d bytes @ %d: ", n2, o2)
		fmt.Printf("%v\n", string(b2[:n2]))

		// The io package provides some functions that may be helpful for file reading. For example, reads like the ones above can be more robustly implemented with ReadAtLeast.
		// 例如，io 包提供了一个更健壮的实现 ReadAtLeast，用于读取上面那种文件。
		o3, err := f.Seek(6, 0)
		check(err)
		b3 := make([]byte, 2)
		n3, err := io.ReadAtLeast(f, b3, 2)
		check(err)
		fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

		// There is no built-in rewind, but Seek(0, 0) accomplishes this.
		// 没有内建的倒带，但是 Seek(0, 0) 实现了这一功能。
		_, err = f.Seek(0, 0)
		check(err)

		// The bufio package implements a buffered reader that may be useful both for its efficiency with many small reads and because of the additional reading methods it provides.
		// bufio 包实现了一个缓冲读取器，这可能有助于提高许多小读操作的效率，以及它提供了很多附加的读取函数。
		r4 := bufio.NewReader(f)
		b4, err := r4.Peek(5)
		check(err)
		fmt.Printf("5 bytes: %s\n", string(b4))

		// Close the file when you’re done (usually this would be scheduled immediately after Opening with defer).
		// 任务结束后要关闭这个文件 （通常这个操作应该在 Open 操作后立即使用 defer 来完成）。
		f.Close()

	},
}

func init() {
	cmd.RootCmd.AddCommand(readingFilesCmd)
}
