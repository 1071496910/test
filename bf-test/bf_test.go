package main

import (
	"github.com/willf/bloom"
	"strconv"
	"testing"
	"unsafe"
)

func BenchmarkBloomFilter(b *testing.B) {
	bf := bloom.New(1000000, 7)
	for i := 0; i < b.N; i++ {
		is := []byte(strconv.Itoa(i))
		bf.Add(is)
		bf.Test(is)
	}

}

func BenchmarkBloomFilter2(b *testing.B) {
	bf := bloom.New(1000000, 7)
	for j := 0; j < b.N; j++ {
		istr := strconv.Itoa(j)
		is := *(*[]byte)(unsafe.Pointer(&(istr)))
		bf.Add(is)
		bf.Test(is)
	}
}
