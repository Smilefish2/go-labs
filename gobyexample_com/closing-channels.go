package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
)

// Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.
// 关闭 一个通道意味着不能再向这个通道发送值了。 该特性可以向通道的接收方传达工作已经完成的信息。

// closingChannelsCmd represents the closing channels command
var closingChannelsCmd = &cobra.Command{
	Use:   "go-by-example:closing-channels",
	Short: "https://gobyexample.com/closing-channels",
	Run: func(cmd *cobra.Command, args []string) {

		// In this example we’ll use a jobs channel to communicate work to be done from the main() goroutine to a worker goroutine. When we have no more jobs for the worker we’ll close the jobs channel.
		// 在这个例子中，我们将使用一个 jobs 通道，将工作内容， 从 main() 协程传递到一个工作协程中。 当我们没有更多的任务传递给工作协程时，我们将 close 这个 jobs 通道。
		jobs := make(chan int, 5)
		done := make(chan bool)

		// Here’s the worker goroutine. It repeatedly receives from jobs with j, more := <-jobs. In this special 2-value form of receive, the more value will be false if jobs has been closed and all values in the channel have already been received. We use this to notify on done when we’ve worked all our jobs.
		// 这是工作协程。使用 j, more := <- jobs 循环的从 jobs 接收数据。 根据接收的第二个值，如果 jobs 已经关闭了， 并且通道中所有的值都已经接收完毕，那么 more 的值将是 false。 当我们完成所有的任务时，会使用这个特性通过 done 通道通知 main 协程。

		go func() {
			for {
				j, more := <-jobs
				if more {
					fmt.Println("received job", j)
				} else {
					fmt.Println("received all jobs")
					done <- true
					return
				}
			}
		}()

		// This sends 3 jobs to the worker over the jobs channel, then closes it.
		// 使用 jobs 发送 3 个任务到工作协程中，然后关闭 jobs。

		for j := 1; j <= 3; j++ {
			jobs <- j
			fmt.Println("sent job", j)
		}
		close(jobs)
		fmt.Println("sent all jobs")

		// We await the worker using the synchronization approach we saw earlier.
		// 使用前面学到的通道同步方法等待任务结束。
		<-done

	},
}

func init() {
	cmd.RootCmd.AddCommand(closingChannelsCmd)
}
