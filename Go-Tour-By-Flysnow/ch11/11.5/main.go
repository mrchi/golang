package main

import (
	"fmt"
	"time"
)

func washVegetables() <-chan string {
	out := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		out <- "洗好的菜"
	}()
	return out
}

func boilWater() <-chan string {
	out := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		out <- "100 度开水"
	}()
	return out
}

func main() {
	// Future模式可以理解为未来模式。
	// 主协程不用等待子协程返回的结果，可以先去做其他事情，等未来需要子协程结果的时候再来取，如果子协程还没有返回结果，就一直等待。

	// Future模式下的协程和普通协程的最大区别是可以返回结果，而这个结果会在未来的某个时间点使用。
	// 所以在未来获取这个结果的操作必须是一个阻塞的操作，要一直等到获取结果为止。

	// 如果你的大任务可以拆解为一个个独立并发执行的小任务，并且可以通过这些小任务的结果得出最终大任务的结果，就可以使用Future模式。

	vegetablesCh := washVegetables()
	waterCh := boilWater()

	fmt.Println("任务已安排，休息一会")
	time.Sleep(2 * time.Second)
	fmt.Println("休息好了，等等看")

	v := <-vegetablesCh
	w := <-waterCh
	fmt.Printf("都准备好了，我有 %s 和 %s\n", v, w)
}
