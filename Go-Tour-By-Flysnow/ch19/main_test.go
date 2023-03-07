package main

import "testing"
import "fmt"
import "strings"
import "bytes"

const QUOTE = "十年生死两茫茫，不思量，自难忘"

// “+”拼接，它是在运行时计算的
func StringPlus(p []string) string {
	var s string
	for i := 0; i < len(p); i++ {
		s += p[i]
	}
	return s
}

// fmt拼接方式借助fmt.Sprint系列函数进行拼接
func StringFmt(p []interface{}) string {
	return fmt.Sprint(p...)
}

// join拼接方式是利用strings.Join函数进行拼接的
func StringJoin(p []string) string {
	return strings.Join(p, "")
}

// buffer拼接也用得很多，它是使用bytes.Buffer进行拼接
func StringBuffer(p []string) string {
	var b bytes.Buffer
	for i := 0; i < len(p); i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

// builder拼接
// 为了改进buffer拼接的性能，从Go 1.10版本开始，增加了一个Builder类型，用于提升字符串拼接的性能
func StringBuilder(p []string, cap int) string {
	var b strings.Builder
	// 一次扩容到位，减少因为频繁扩容导致的内存分配
	b.Grow(cap)
	for i := 0; i < len(p); i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

func initStringS(N int) []string {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = QUOTE
	}
	return s
}

// 该函数返回的是[]interface{}，这是专门为StringFmt(p []interface{})拼接函数准备的，减少了类型之间的转换。
func initStringI(N int) []interface{} {
	s := make([]interface{}, N)
	for i := 0; i < N; i++ {
		s[i] = QUOTE
	}
	return s
}

// -------------------- 10 --------------------

func BenchmarkStringPlus10(b *testing.B) {
	p := initStringS(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringPlus(p)
	}
}

func BenchmarkStringFmt10(b *testing.B) {
	p := initStringI(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringFmt(p)
	}
}

func BenchmarkStringJoin10(b *testing.B) {
	p := initStringS(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringJoin(p)
	}
}

func BenchmarkStringBuffer10(b *testing.B) {
	p := initStringS(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuffer(p)
	}
}

func BenchmarkStringBuilder10(b *testing.B) {
	p := initStringS(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuilder(p, 10*len(QUOTE))
	}
}

// -------------------- 100 --------------------

func BenchmarkStringPlus100(b *testing.B) {
	p := initStringS(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringPlus(p)
	}
}

func BenchmarkStringFmt100(b *testing.B) {
	p := initStringI(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringFmt(p)
	}
}

func BenchmarkStringJoin100(b *testing.B) {
	p := initStringS(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringJoin(p)
	}
}

func BenchmarkStringBuffer100(b *testing.B) {
	p := initStringS(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuffer(p)
	}
}

func BenchmarkStringBuilder100(b *testing.B) {
	p := initStringS(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuilder(p, 100*len(QUOTE))
	}
}

// -------------------- 1000 --------------------

func BenchmarkStringPlus1000(b *testing.B) {
	p := initStringS(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringPlus(p)
	}
}

func BenchmarkStringFmt1000(b *testing.B) {
	p := initStringI(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringFmt(p)
	}
}

func BenchmarkStringJoin1000(b *testing.B) {
	p := initStringS(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringJoin(p)
	}
}

func BenchmarkStringBuffer1000(b *testing.B) {
	p := initStringS(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuffer(p)
	}
}

func BenchmarkStringBuilder1000(b *testing.B) {
	p := initStringS(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuilder(p, 1000*len(QUOTE))
	}
}

// 如果有现成的数组、切片，那么可以直接使用join拼接，但是如果没有，并且追求灵活性拼接，则还是选择builder拼接。
