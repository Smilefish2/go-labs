package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
	"math/rand"
	"time"
)

// Go’s math/rand package provides pseudorandom number generation.
// Go 的 math/rand 包提供了伪随机数生成器。

// randomNumbersCmd represents the random-numbers command
var randomNumbersCmd = &cobra.Command{
	Use:   "go-by-example:random-numbers",
	Short: "https://gobyexample.com/random-numbers",
	Run: func(cmd *cobra.Command, args []string) {

		// For example, rand.Intn returns a random int n, 0 <= n < 100.
		// 例如，rand.Intn 返回一个随机的整数 n，且 0 <= n < 100。
		fmt.Print(rand.Intn(100), ",")
		fmt.Print(rand.Intn(100))
		fmt.Println()

		// rand.Float64 returns a float64 f, 0.0 <= f < 1.0.
		// rand.Float64 返回一个64位浮点数 f，且 0.0 <= f < 1.0。
		fmt.Println(rand.Float64())

		// This can be used to generate random floats in other ranges, for example 5.0 <= f' < 10.0.
		// 这个技巧可以用来生成其他范围的随机浮点数， 例如，5.0 <= f < 10.0
		fmt.Print((rand.Float64()*5)+5, ",")
		fmt.Print((rand.Float64() * 5) + 5)
		fmt.Println()

		// The default number generator is deterministic, so it’ll produce the same sequence of numbers each time by default. To produce varying sequences, give it a seed that changes. Note that this is not safe to use for random numbers you intend to be secret, use crypto/rand for those.
		// 默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。 要产生不同的数字序列，需要给定一个不同的种子。 注意，对于想要加密的随机数，使用此方法并不安全， 应该使用 crypto/rand。
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		// Call the resulting rand.Rand just like the functions on the rand package.
		// 调用上面返回的 rand.Rand，就像调用 rand 包中函数一样。
		fmt.Print(r1.Intn(100), ",")
		fmt.Print(r1.Intn(100))
		fmt.Println()

		// If you seed a source with the same number, it produces the same sequence of random numbers.
		// 如果使用相同种子生成的随机数生成器，会生成相同的随机数序列。
		s2 := rand.NewSource(42)
		r2 := rand.New(s2)
		fmt.Print(r2.Intn(100), ",")
		fmt.Print(r2.Intn(100))
		fmt.Println()
		s3 := rand.NewSource(42)
		r3 := rand.New(s3)
		fmt.Print(r3.Intn(100), ",")
		fmt.Print(r3.Intn(100))

		// See the math/rand package docs for references on other random quantities that Go can provide.
		// 有关 Go 提供的其他随机数的信息， 请参阅 math/rand 包文档。
	},
}

func init() {
	cmd.RootCmd.AddCommand(randomNumbersCmd)
}
