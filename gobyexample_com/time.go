package gobyexample_com

import (
	"fmt"
	"learn-go-with-cli/cmd"
	"time"

	"github.com/spf13/cobra"
)

// Go offers extensive support for times and durations; here are some examples.
// Go 为时间（time）和时间段（duration）提供了大量的支持；这儿有是一些例子。

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "go-by-example:time",
	Short: "https://gobyexample.com/time",
	Run: func(cmd *cobra.Command, args []string) {

		p := fmt.Println

		// We’ll start by getting the current time.
		// 从获取当前时间时间开始。
		now := time.Now()
		p(now)

		// You can build a time struct by providing the year, month, day, etc. Times are always associated with a Location, i.e. time zone.
		// 通过提供年月日等信息，你可以构建一个 time。 时间总是与 Location 有关，也就是时区。
		then := time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
		p(then)

		// You can extract the various components of the time value as expected.
		// 你可以提取出时间的各个组成部分。

		p(then.Year())
		p(then.Month())
		p(then.Day())
		p(then.Hour())
		p(then.Minute())
		p(then.Second())
		p(then.Nanosecond())
		p(then.Location())

		// The Monday-Sunday Weekday is also available.
		// 支持通过 Weekday 输出星期一到星期日。

		p(then.Weekday())

		// These methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively.
		// 这些方法用来比较两个时间，分别测试一下是否为之前、之后或者是同一时刻，精确到秒。

		p(then.Before(now))
		p(then.After(now))
		p(then.Equal(now))

		// The Sub methods returns a Duration representing the interval between two times.
		// 方法 Sub 返回一个 Duration 来表示两个时间点的间隔时间。

		diff := now.Sub(then)
		p(diff)

		// We can compute the length of the duration in various units.
		// 我们可以用各种单位来表示时间段的长度。
		p(diff.Hours())
		p(diff.Minutes())
		p(diff.Seconds())
		p(diff.Nanoseconds())

		// You can use Add to advance a time by a given duration, or with a - to move backwards by a duration.
		// 你可以使用 Add 将时间后移一个时间段，或者使用一个 - 来将时间前移一个时间段。
		p(then.Add(diff))
		p(then.Add(-diff))

		// Next we’ll look at the related idea of time relative to the Unix epoch.
		// 接下来，我们将研究与 Unix 纪元相关的概念。

	},
}

func init() {
	cmd.RootCmd.AddCommand(timeCmd)
}
