package gobyexample_com

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"time"
)

// Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service. Go elegantly supports rate limiting with goroutines, channels, and tickers.
// 速率限制 是控制服务资源利用和质量的重要机制。 基于协程、通道和打点器，Go 优雅的支持速率限制。

// rateLimitingCmd represents the rate limiting command
var rateLimitingCmd = &cobra.Command{
	Use:   "go-by-example:rate-limiting",
	Short: "https://gobyexample.com/rate-limiting",
	Run: func(cmd *cobra.Command, args []string) {

		// First we’ll look at basic rate limiting. Suppose we want to limit our handling of incoming requests. We’ll serve these requests off a channel of the same name.
		// 首先，我们将看一个基本的速率限制。 假设我们想限制对收到请求的处理，我们可以通过一个渠道处理这些请求。

		requests := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			requests <- i
		}
		close(requests)

		// This limiter channel will receive a value every 200 milliseconds. This is the regulator in our rate limiting scheme.
		// limiter 通道每 200ms 接收一个值。 这是我们任务速率限制的调度器。

		limiter := time.Tick(200 * time.Millisecond)

		// By blocking on a receive from the limiter channel before serving each request, we limit ourselves to 1 request every 200 milliseconds.
		// 通过在每次请求前阻塞 limiter 通道的一个接收， 可以将频率限制为，每 200ms 执行一次请求。
		for req := range requests {
			<-limiter
			fmt.Println("request", req, time.Now())
		}

		// We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit. We can accomplish this by buffering our limiter channel. This burstyLimiter channel will allow bursts of up to 3 events.
		// 有时候我们可能希望在速率限制方案中允许短暂的并发请求，并同时保留总体速率限制。 我们可以通过缓冲通道来完成此任务。 burstyLimiter 通道允许最多 3 个爆发（bursts）事件。

		burstyLimiter := make(chan time.Time, 3)

		// Fill up the channel to represent allowed bursting.
		// 填充通道，表示允许的爆发（bursts）。
		for i := 0; i < 3; i++ {
			burstyLimiter <- time.Now()
		}

		// Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
		// 每 200ms 我们将尝试添加一个新的值到 burstyLimiter中， 直到达到 3 个的限制。
		go func() {
			for t := range time.Tick(200 * time.Millisecond) {
				burstyLimiter <- t
			}
		}()

		// Now simulate 5 more incoming requests. The first 3 of these will benefit from the burst capability of burstyLimiter.
		// 现在，模拟另外 5 个传入请求。 受益于 burstyLimiter 的爆发（bursts）能力，前 3 个请求可以快速完成。
		burstyRequests := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			burstyRequests <- i
		}
		close(burstyRequests)
		for req := range burstyRequests {
			<-burstyLimiter
			fmt.Println("request", req, time.Now())
		}

		// Running our program we see the first batch of requests handled once every ~200 milliseconds as desired.
		// 运行程序，我们看到第一批请求意料之中的大约每 200ms 处理一次。

		// For the second batch of requests we serve the first 3 immediately because of the burstable rate limiting, then serve the remaining 2 with ~200ms delays each.
		// 第二批请求，由于爆发（burstable）速率控制，我们直接连续处理了 3 个请求， 然后以大约每 200ms 一次的速度，处理了剩余的 2 个请求。
	},
}

func init() {
	cmd.RootCmd.AddCommand(rateLimitingCmd)
}
