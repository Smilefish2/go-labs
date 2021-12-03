package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"sync"
	"time"
)

// To wait for multiple goroutines to finish, we can use a wait group.
// 想要等待多个协程完成，我们可以使用 *wait group*。

// This is the function we’ll run in every goroutine.
// 这是我们将在每个goroutine中运行的函数。

func workerGroups(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	// 睡眠来模拟一项昂贵的任务。
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// waitGroupsCmd represents the wait groups command
var waitGroupsCmd = &cobra.Command{
	Use:   "go-by-example:waitgroups",
	Short: "https://gobyexample.com/waitgroups",
	Run: func(cmd *cobra.Command, args []string) {

		// This WaitGroup is used to wait for all the goroutines launched here to finish. Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
		// // 这个 WaitGroup 被用于等待该函数开启的所有协程。
		var wg sync.WaitGroup

		// Launch several goroutines and increment the WaitGroup counter for each.
		// 开启几个协程，并为其递增 WaitGroup 的计数器。
		for i := 1; i <= 5; i++ {
			wg.Add(1)

			// Avoid re-use of the same i value in each goroutine closure. See the FAQ for more details.
			// 避免在每个goroutine闭包中重复使用相同的i值。有关更多详细信息，请参阅常见问题解答。
			i := i

			// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done. This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
			// 将worker调用包装在一个闭包中，确保告诉WaitGroup该worker已完成。这样，工作者本身就不必知道其执行中涉及的并发原语。
			go func() {
				defer wg.Done()
				workerGroups(i)
			}()
		}

		// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
		// 阻塞，直到 WaitGroup 计数器恢复为 0，即所有协程的工作都已经完成。
		wg.Wait()

		// Note that this approach has no straightforward way to propagate errors from workers. For more advanced use cases, consider using the errgroup package.
		// 请注意，这种方法没有直接的方法来传播工作人员的错误。对于更高级的用例，考虑使用ErrGeo包。

		// The order of workers starting up and finishing is likely to be different for each invocation.
		// 对于每个调用，工人启动和完成的顺序可能不同。
	},
}

func init() {
	cmd.RootCmd.AddCommand(waitGroupsCmd)
}
