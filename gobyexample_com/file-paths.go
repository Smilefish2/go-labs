package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
	"path/filepath"
	"strings"
)

// The filepath package provides functions to parse and construct file paths in a way that is portable between operating systems; dir/file on Linux vs. dir\file on Windows, for example.
// filepath 包为 文件路径 ，提供了方便的跨操作系统的解析和构建函数； 比如：Linux 下的 dir/file 和 Windows 下的 dir\file 。

// filePathsCmd represents the file-paths command
var filePathsCmd = &cobra.Command{
	Use:   "go-by-example:file-paths",
	Short: "https://gobyexample.com/file-paths",
	Run: func(cmd *cobra.Command, args []string) {

		// Join should be used to construct paths in a portable way. It takes any number of arguments and constructs a hierarchical path from them.
		// 应使用 Join 来构建可移植(跨操作系统)的路径。 它接收任意数量的参数，并参照传入顺序构造一个对应层次结构的路径。
		p := filepath.Join("dir1", "dir2", "filename")
		fmt.Println("p:", p)

		// You should always use Join instead of concatenating /s or \s manually. In addition to providing portability, Join will also normalize paths by removing superfluous separators and directory changes.
		// 您应该总是使用 Join 代替手动拼接 / 和 \。 除了可移植性，Join 还会删除多余的分隔符和目录，使得路径更加规范。
		fmt.Println(filepath.Join("dir1//", "filename"))
		fmt.Println(filepath.Join("dir1/../dir1", "filename"))

		// Dir and Base can be used to split a path to the directory and the file. Alternatively, Split will return both in the same call.
		// Dir 和 Base 可以被用于分割路径中的目录和文件。 此外，Split 可以一次调用返回上面两个函数的结果。
		fmt.Println("Dir(p):", filepath.Dir(p))
		fmt.Println("Base(p):", filepath.Base(p))

		// We can check whether a path is absolute.
		// 判断路径是否为绝对路径。
		fmt.Println(filepath.IsAbs("dir/file"))
		fmt.Println(filepath.IsAbs("/dir/file"))

		filename := "config.json"

		// Some file names have extensions following a dot. We can split the extension out of such names with Ext.
		// 某些文件名包含了扩展名（文件类型）。 我们可以用 Ext 将扩展名分割出来。
		ext := filepath.Ext(filename)
		fmt.Println(ext)

		// To find the file’s name with the extension removed, use strings.TrimSuffix.
		// 想获取文件名清除扩展名后的值，请使用 strings.TrmSuffix。
		fmt.Println(strings.TrimSuffix(filename, ext))

		// Rel finds a relative path between a base and a target. It returns an error if the target cannot be made relative to base.
		// Rel 寻找 basepath 与 targpath 之间的相对路径。 如果相对路径不存在，则返回错误。
		rel, err := filepath.Rel("a/b", "a/b/t/file")
		if err != nil {
			panic(err)
		}
		fmt.Println(rel)

		rel, err = filepath.Rel("a/b", "a/c/t/file")
		if err != nil {
			panic(err)
		}
		fmt.Println(rel)

	},
}

func init() {
	cmd.RootCmd.AddCommand(filePathsCmd)
}
