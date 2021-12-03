package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"time"
)

// Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals. Here’s an example of a ticker that ticks periodically until we stop it.
// 定时器 是当你想要在未来某一刻执行一次时使用的 - 打点器 则是为你想要以固定的时间间隔重复执行而准备的。 这里是一个打点器的例子，它将定时的执行，直到我们将它停止。

// tickersCmd represents the tickers command
var tickersCmd = &cobra.Command{
	Use:   "go-by-example:tickers",
	Short: "https://gobyexample.com/tickers",
	Run: func(cmd *cobra.Command, args []string) {

		// Tickers use a similar mechanism to timers: a channel that is sent values. Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.
		// 打点器和定时器的机制有点相似：使用一个通道来发送数据。 这里我们使用通道内建的 select，等待每 500ms 到达一次的值。

		ticker := time.NewTicker(500 * time.Millisecond)
		done := make(chan bool)

		go func() {
			for {
				select {
				case <-done:
					return
				case t := <-ticker.C:
					fmt.Println("Tick at", t)
				}
			}
		}()

		// Tickers can be stopped like timers. Once a ticker is stopped it won’t receive any more values on its channel. We’ll stop ours after 1600ms.
		// 打点器可以和定时器一样被停止。 打点器一旦停止，将不能再从它的通道中接收到值。 我们将在运行 1600ms 后停止这个打点器。

		time.Sleep(1600 * time.Millisecond)
		ticker.Stop()
		done <- true
		fmt.Println("Ticker stopped")

		// When we run this program the ticker should tick 3 times before we stop it.
		// 当我们运行这个程序时，打点器会在我们停止它前打点 3 次。

	},
}

func init() {
	cmd.RootCmd.AddCommand(tickersCmd)
}
