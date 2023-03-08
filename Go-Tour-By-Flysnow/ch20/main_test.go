package main

import "testing"

// 含有单元测试代码的go文件必须以_test.go结尾
// 单元测试的函数名必须以Test开头，是可导出的、公开的函数。
// 测试函数的签名必须接收一个指向testing.T类型的指针，并且不能返回任何值。

// 单元测试文件名_test.go前面的部分最好是被测试的函数所在的go文件的文件名
// 函数名最好是“Test+要测试的函数名”

// 单测覆盖率
// -coverprofile 这个Flag，它可以得到一个单元测试覆盖率文件，运行这行命令还可以同时看到测试覆盖率
// go test -coverprofile=ch20.cover .
// 或者 go test -cover

// 生成覆盖率 html
// go tool cover -html=ch20.cover -o=ch20.html
func TestFibonacci(t *testing.T) {
	fsMap := map[int]int{
		-1: 0,
		0:  0,
		1:  1,
		2:  1,
		3:  2,
		4:  3,
		5:  5,
		6:  8,
		7:  13,
		8:  21,
		9:  34,
	}
	for k, v := range fsMap {
		fib := Fibonacci(k)
		if v == fib {
			t.Logf("PASS, expect=%v, get=%v", v, fib)
		} else {
			t.Errorf("FAIL, expect=%v, get=%v", v, fib)
		}
	}
}

// 基准测试函数必须以Benchmark开头，必须是可导出的。
// 函数的签名必须接收一个指向testing.B类型的指针，并且不能返回任何值。
// 最后的for循环很重要，被测试的代码要放到循环里。b.N是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能。

// -bench这个Flag，它接收一个表达式作为参数，以匹配基准测试的函数，“.”表示运行所有基准测试。
// go test -bench .

// 基准测试的时间默认是1s，如果想让测试运行的时间更长，可以通过-benchtime指定
// go test -bench=. -benchtime=3s .

// BenchmarkFabonacci-6   	14838926	       238.6 ns/op	       0 B/op	       0 allocs/op
// 【BenchmarkFabonacci-6】运行基准测试时对应的GOMAXPROCS的值。
// 【14838926】运行for循环的次数，也就是调用被测试代码的次数
// 【238.6 ns/op】每次需要花费的时间
// 【0 B/op】表示每次操作分配了多少字节的内存
// 【0 allocs/op】表示每次操作分配内存的次数。

func BenchmarkFabonacci(b *testing.B) {
	n := 10
	// 要开启内存统计也比较简单，即通过ReportAllocs()方法开启
	b.ReportAllocs()
	// 通过ResetTimer方法重置计时器（还有StartTimer和StopTimer方法）
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fibonacci(n)
	}
}

// Go语言通过RunParallel方法运行并发基准测试。RunParallel方法会创建多个goroutine，并将b.N分配给这些goroutine执行。
func BenchmarkFabonacciRunParallel(b *testing.B) {
	n := 10
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fibonacci(n)
		}
	})
}
