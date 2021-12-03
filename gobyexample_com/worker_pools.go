package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"time"
)

// In this example we’ll look at how to implement a worker pool using goroutines and channels.
// 在这个例子中，我们将看到如何使用协程与通道实现一个_工作池_。

// Here’s the worker, of which we’ll run several concurrent instances. These workers will receive work on the jobs channel and send the corresponding results on results. We’ll sleep a second per job to simulate an expensive task.
// 这是 worker 程序，我们会并发的运行多个 worker。 worker 将在 jobs 频道上接收工作，并在 results 上发送相应的结果。 每个 worker 我们都会 sleep 一秒钟，以模拟一项昂贵的（耗时一秒钟的）任务。

func workerPools(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

// workerPoolsCmd represents the worker pools command
var workerPoolsCmd = &cobra.Command{
	Use:   "go-by-example:worker_pools",
	Short: "https://gobyexample.com/worker-pools",
	Run: func(cmd *cobra.Command, args []string) {

		// In order to use our pool of workers we need to send them work and collect their results. We make 2 channels for this.
		// 为了使用 worker 工作池并且收集其的结果，我们需要 2 个通道。

		const numJobs = 5
		jobs := make(chan int, numJobs)
		results := make(chan int, numJobs)

		// This starts up 3 workers, initially blocked because there are no jobs yet.
		// 这里启动了 3 个 worker， 初始是阻塞的，因为还没有传递任务。

		for w := 1; w <= 3; w++ {
			go workerPools(w, jobs, results)
		}

		// Here we send 5 jobs and then close that channel to indicate that’s all the work we have.
		// 这里我们发送 5 个 jobs， 然后 close 这些通道，表示这些就是所有的任务了。

		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)

		// Finally we collect all the results of the work. This also ensures that the worker goroutines have finished. An alternative way to wait for multiple goroutines is to use a WaitGroup.
		// 最后，我们收集所有这些任务的返回值。 这也确保了所有的 worker 协程都已完成。 另一个等待多个协程的方法是使用WaitGroup。

		for a := 1; a <= numJobs; a++ {
			<-results
		}

		// Our running program shows the 5 jobs being executed by various workers. The program only takes about 2 seconds despite doing about 5 seconds of total work because there are 3 workers operating concurrently.
		// 运行程序，显示 5 个任务被多个 worker 执行。 尽管所有的工作总共要花费 5 秒钟，但该程序只花了 2 秒钟， 因为 3 个 worker 是并行的。
	},
}

func init() {
	cmd.RootCmd.AddCommand(workerPoolsCmd)
}
