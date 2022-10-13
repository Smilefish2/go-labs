package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
	"os"
	"path/filepath"
)

// Throughout program execution, we often want to create data that isn’t needed after the program exits. Temporary files and directories are useful for this purpose since they don’t pollute the file system over time.
// 在程序运行时，我们经常创建一些运行时用到，程序结束后就不再使用的数据。 临时目录和文件 对于上面的情况很有用，因为它不会随着时间的推移而污染文件系统。

// temporaryFilesAndDirectoriesCmd represents the temporary-files-and-directories command
var temporaryFilesAndDirectoriesCmd = &cobra.Command{
	Use:   "go-by-example:temporary-files-and-directories",
	Short: "https://gobyexample.com/temporary-files-and-directories",
	Run: func(cmd *cobra.Command, args []string) {

		// The easiest way to create a temporary file is by calling os.CreateTemp. It creates a file and opens it for reading and writing. We provide "" as the first argument, so os.CreateTemp will create the file in the default location for our OS.
		// 创建临时文件最简单的方法是调用 ioutil.TempFile 函数。 它会创建并打开文件，我们可以对文件进行读写。 函数的第一个参数传 ""，ioutil.TempFile 会在操作系统的默认位置下创建该文件。
		f, err := os.CreateTemp("", "sample")
		check(err)

		// Display the name of the temporary file. On Unix-based OSes the directory will likely be /tmp. The file name starts with the prefix given as the second argument to os.CreateTemp and the rest is chosen automatically to ensure that concurrent calls will always create different file names.
		// 打印临时文件的名称。 文件名以 ioutil.TempFile 函数的第二个参数作为前缀， 剩余的部分会自动生成，以确保并发调用时，生成不重复的文件名。 在类 Unix 操作系统下，临时目录一般是 /tmp。
		fmt.Println("Temp file name:", f.Name())

		// Clean up the file after we’re done. The OS is likely to clean up temporary files by itself after some time, but it’s good practice to do this explicitly.
		// defer 删除该文件。 尽管操作系统会自动在某个时间清理临时文件，但手动清理是一个好习惯。
		defer os.Remove(f.Name())

		// We can write some data to the file.
		// 我们可以向文件写入一些数据。
		_, err = f.Write([]byte{1, 2, 3, 4})
		check(err)

		// If we intend to write many temporary files, we may prefer to create a temporary directory. os.MkdirTemp’s arguments are the same as CreateTemp’s, but it returns a directory name rather than an open file.
		// 如果需要写入多个临时文件，最好是为其创建一个临时 目录 。 ioutil.TempDir 的参数与 TempFile 相同， 但是它返回的是一个 目录名 ，而不是一个打开的文件。
		dname, err := os.MkdirTemp("", "sampledir")
		check(err)
		fmt.Println("Temp dir name:", dname)

		// Now we can synthesize temporary file names by prefixing them with our temporary directory.
		// 现在，我们可以通过拼接临时目录和临时文件合成完整的临时文件路径，并写入数据。
		defer os.RemoveAll(dname)

		fname := filepath.Join(dname, "file1")
		err = os.WriteFile(fname, []byte{1, 2}, 0666)
		check(err)

	},
}

func init() {
	cmd.RootCmd.AddCommand(temporaryFilesAndDirectoriesCmd)
}
