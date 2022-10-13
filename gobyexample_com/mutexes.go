package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-labs/cmd"
	"sync"
)

// In the previous example we saw how to manage simple counter state using atomic operations. For more complex state we can use a mutex to safely access data across multiple goroutines.
// 在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器。 对于更加复杂的情况，我们可以使用一个互斥锁 来在 Go 协程间安全的访问数据。

// Container holds a map of counters; since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronize access. Note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer.
// 容器包含计数器的地图；因为我们想从多个goroutine中同时更新它，所以我们添加了一个互斥来同步访问。请注意，不能复制互斥体，所以如果传递此结构，则应通过指针完成。
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Lock the mutex before accessing counters; unlock it at the end of the function using a defer statement.
// 在访问计数器之前锁定互斥锁；在函数末尾使用DEBER语句将其解锁。

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

// mutexesdCmd represents the mutexes command
var mutexesdCmd = &cobra.Command{
	Use:   "go-by-example:mutexes",
	Short: "https://gobyexample.com/mutexes",
	Run: func(cmd *cobra.Command, args []string) {

		// Note that the zero value of a mutex is usable as-is, so no initialization is required here.
		// 请注意，互斥量的零值是可用的，因此这里不需要初始化。
		c := Container{
			counters: map[string]int{"a": 0, "b": 0},
		}

		var wg sync.WaitGroup

		// This function increments a named counter in a loop.
		// 此函数递增循环中的命名计数器。
		doIncrement := func(name string, n int) {
			for i := 0; i < n; i++ {
				c.inc(name)
			}
			wg.Done()
		}

		// Run several goroutines concurrently; note that they all access the same Container, and two of them access the same counter.
		// 同时运行多个goroutine；请注意，它们都访问相同的容器，其中两个访问相同的计数器。

		wg.Add(3)
		go doIncrement("a", 10000)
		go doIncrement("a", 10000)
		go doIncrement("b", 10000)

		// Wait a for the goroutines to finish
		// 等待一段时间，等待goroutines完成
		wg.Wait()
		fmt.Println(c.counters)

		// Running the program shows that the counters updated as expected.
		// 运行该程序表明计数器已按预期更新。

		// Next we’ll look at implementing this same state management task using only goroutines and channels.
		// 接下来，我们将只使用goroutine和channels来实现相同的状态管理任务。
	},
}

func init() {
	cmd.RootCmd.AddCommand(mutexesdCmd)
}
