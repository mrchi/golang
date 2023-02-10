package main

import (
	"fmt"
	"gadget"
)

// 一个接口是特定值预期具有的一组方法。
// 一个类型可以满足多个接口，一个接口（通常应该）可以有多个类型满足它。
// 把接口定义在调用的包中会更灵活
type Device interface {
	Play(string)
	Stop()
}

// 也能将函数的参数定义为接口类型。
func playList(device Device, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	// 一个接口类型的变量能够保存任何满足接口的类型的值
	var device Device
	mixtype := []string{"不想长大", "波斯猫", "中国话"}

	device = gadget.TapePlayer{}
	playList(device, mixtype)
	fmt.Println()

	// 可以将具有其他方法的类型赋值给接口类型。
	device = gadget.TapeRecorder{}
	playList(device, mixtype)
	fmt.Println()

	// 但只能调用接口定义的方法，其他方法不可以
	// 报错：device.Record()
}
