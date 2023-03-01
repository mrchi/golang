package main

import (
	"io"
	"os"
	"sync"
	"testing"
)

// 字节切片[]byte 在多场景中用于做缓冲。
// 自定义字节池，重用已经创建好的[]byte来提高对象的使用率，降低内存的申请和GC
type BytePoolCap struct {
	c      chan []byte // 一个channel，用于充当字节缓存池。
	length int         // 使用make函数创建[]byte时的len参数。
	cap    int         // 使用make函数创建[]byte时的cap参数。
}

// 获取字节切片
func (bpc *BytePoolCap) Get() (b []byte) {
	select {
	// 尝试从 channel 中获取一个
	case b = <-bpc.c:
	// 如果不存在则创建一个
	default:
		if bpc.cap > 0 {
			b = make([]byte, bpc.length, bpc.cap)
		} else {
			b = make([]byte, bpc.length)
		}
	}
	return
}

// 用完后放回字节切片
func (bpc *BytePoolCap) Put(b []byte) {
	select {
	// 尝试放回到 channel 中
	case bpc.c <- b:
	// 放不回去就算了
	default:
	}
}

// 工厂函数
func NewBytePoolCap(maxSize int, length int, capacity int) (bp *BytePoolCap) {
	return &BytePoolCap{
		c:      make(chan []byte, maxSize),
		length: length,
		cap:    capacity,
	}
}

//-------------benchmark----------------

// 模拟操作文件的函数
func mockReadFile(b []byte) {
	f, _ := os.Open("main.go")
	for {
		n, err := io.ReadFull(f, b)
		if n == 0 || err == io.EOF {
			break
		}
	}
}

// 自定义字节池 Benchmark 测试函数
func opBytePool(bpc *BytePoolCap) {
	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func(bpc *BytePoolCap) {
			defer wg.Done()
			buffer := bpc.Get()
			defer bpc.Put(buffer)
			mockReadFile(buffer)
		}(bpc)
	}

	wg.Wait()
}

// sync.Pool Benchmark 测试函数
func opSyncPool(sp *sync.Pool) {
	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func(sp *sync.Pool) {
			defer wg.Done()
			buffer := sp.Get().([]byte)
			defer sp.Put(buffer)
			mockReadFile(buffer)
		}(sp)
	}

	wg.Wait()
}

func BenchmarkBytePool(b *testing.B) {
	// BytePoolCap是自定义的，所以它存放的对象类型明确，不用经过一层类型断言转换，还可以自己定制对象池的大小等。
	bpc := NewBytePoolCap(500, 1024, 1024)
	opBytePool(bpc)
}

func BenchmarkSyncPool(b *testing.B) {
	// sync.Pool可以存放任何对象，但需要经过一层类型断言转换
	sp := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}
	opSyncPool(sp)
}

func main() {}
