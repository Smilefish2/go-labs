package gobyexample_com

import (
	"fmt"
	"go-labs/cmd"
	"time"

	"github.com/spf13/cobra"
)

// A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the Unix epoch. Here’s how to do it in Go.
// 一般程序会有获取 Unix 时间 的秒数，毫秒数，或者微秒数的需求。来看看如何用 Go 来实现。

// epochCmd represents the epoch command
var epochCmd = &cobra.Command{
	Use:   "go-by-example:epoch",
	Short: "https://gobyexample.com/epoch",
	Run: func(cmd *cobra.Command, args []string) {

		// Use time.Now with Unix or UnixNano to get elapsed time since the Unix epoch in seconds or nanoseconds, respectively.
		// 分别使用 time.Now 的 Unix 和 UnixNano， 来获取从 Unix 纪元起，到现在经过的秒数和纳秒数。
		now := time.Now()
		secs := now.Unix()
		nanos := now.UnixNano()
		fmt.Println(now)

		// Note that there is no UnixMillis, so to get the milliseconds since epoch you’ll need to manually divide from nanoseconds.
		// 注意 UnixMillis 是不存在的，所以要得到毫秒数的话， 你需要手动的从纳秒转化一下。
		millis := nanos / 1000000
		fmt.Println(secs)
		fmt.Println(millis)
		fmt.Println(nanos)

		// You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.
		// 你也可以将 Unix 纪元起的整数秒或者纳秒转化到相应的时间。
		fmt.Println(time.Unix(secs, 0))
		fmt.Println(time.Unix(0, nanos))

		// Next we’ll look at another time-related task: time parsing and formatting.
		// 下面我们将看看另一个时间相关的任务：时间解析与格式化。
	},
}

func init() {
	cmd.RootCmd.AddCommand(epochCmd)
}
