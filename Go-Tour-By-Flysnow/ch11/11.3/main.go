package main

import "fmt"

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

func main() {
	// 流水线模式也称为Pipeline模式，传递的数据称为数据流

	// 流水线由一道道工序构成，每道工序通过channel把数据传递到下一个工序
	// 每道工序一般会对应一个函数，函数里有协程和channel，协程一般用于处理数据并把它放入channel中，整个函数会返回这个channel以供下一道工序使用。
	// 最终要有一个组织者把这些工序串起来，这样就形成了一个完整的流水线，对于数据来说就是数据流。
	coms := buy(3)
	phones := build(coms)
	packs := pack(phones)

	for p := range packs {
		fmt.Println(p)
	}
}
