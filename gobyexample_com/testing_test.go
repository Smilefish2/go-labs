package gobyexample_com

import (
	"fmt"
	"testing"
)

// We’ll be testing this simple implementation of an integer minimum. Typically, the code we’re testing would be in a source file named something like intutils.go, and the test file for it would then be named intutils_test.go.
// 我们要测试下面这个简单的函数——返回最小值。 一般地，需要被测试的代码应该在类似于 intutils.go 的文件下， 其对应的测试文件应该被命名为 intutils_test.go。
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// A test is created by writing a function with a name beginning with Test.
// 通常编写一个名称以 Test 开头的函数来创建测试。
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error* will report test failures but continue executing the test. t.Fatal* will report test failures and stop the test immediately.
		// t.Error* 会报告测试失败的信息，然后继续运行测试。 t.Fail* 会报告测试失败的信息，然后立即终止测试。
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Writing tests can be repetitive, so it’s idiomatic to use a table-driven style, where test inputs and expected outputs are listed in a table and a single loop walks over them and performs the test logic.
// 单元测试可以重复，所以会经常使用 表驱动 风格编写单元测试， 表中列出了输入数据，预期输出，使用循环，遍历并执行测试逻辑。
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// t.Run enables running “subtests”, one for each table entry. These are shown separately when executing go test -v.
	// t.Run 可以运行一个 “subtests” 子测试，一个子测试对应表中一行数据。 运行 go test -v 时，他们会分开显示。
	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmark tests typically go in _test.go files and are named beginning with Benchmark. The testing runner executes each benchmark function several times, increasing b.N on each run until it collects a precise measurement.
// 基准测试通常位于_test.go文件中，并以基准开始命名。测试运行程序多次执行每个基准函数，每次运行时增加b.N，直到收集到精确的测量值。

func BenchmarkIntMin(b *testing.B) {

	// Typically the benchmark runs a function we’re benchmarking in a loop b.N times.
	// 通常，基准测试运行一个函数，我们在循环b.N次中进行基准测试。
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
