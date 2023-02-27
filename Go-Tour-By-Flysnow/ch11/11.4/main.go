package main

import (
	"fmt"
	"sync"
	"time"
)

func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < 3; i++ {
			out <- fmt.Sprintf("配件%d", i)
		}
	}()
	return out
}

func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			time.Sleep(1 * time.Second)
			out <- fmt.Sprintf("组装（%s）", c)
		}
	}()
	return out
}

func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- fmt.Sprintf("打包（%s）", c)
		}
	}()
	return out
}

// 扇入组件（merge函数非常小，而且与业务无关，不能当作一道工序，所以我把它叫作组件。）
// 该merge组件是可以复用的，流水线中的任何工序需要扇入的时候，都可以使用merge组件。
func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}

	wg.Add(len(ins))

	// 扇入
	for _, in := range ins {
		go p(in)
	}

	// 启动一个单独的 goroutine，等全部扇入的 goroutine 结束后关闭 channel
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// 扇入和扇出：扇出的数据流向是发散传递出去，是输出流；扇入的数据流向是汇聚进来，是输入流。

	coms := buy(3)

	// 扇出，多个函数同时消费同一个 goroutine 生产者输出的 channel
	phones1 := build(coms)
	phones2 := build(coms)
	phones3 := build(coms)

	phones := merge(phones1, phones2, phones3) // 扇入
	packs := pack(phones)

	for p := range packs {
		fmt.Println(p)
	}
}
