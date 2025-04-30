package main

import (
	"fmt"
	"testing"
)

var (
	table map[string]string = map[string]string{
		"01K": "testfile-1K",
		"04K": "testfile-4K",
		"64K": "testfile-64K",
		"01M": "testfile-1M",
		"04M": "testfile-4M",
		"64M": "testfile-64M",
	}
	keys []string = []string{
		"01K",
		"04K",
		"64K",
		"01M",
		"04M",
		"64M",
	}
)

func Benchmark(b *testing.B) {
	for _, benchName := range keys {
		benchfile := table[benchName]
		b.Run(fmt.Sprintf("krnlRead/%s", benchName), func(b *testing.B) {
			for b.Loop() {
				krnlRead(fmt.Sprintf("./testdata/%s", benchfile))
			}
		})

		b.Run(fmt.Sprintf("mmapRead/%s", benchName), func(b *testing.B) {
			for b.Loop() {
				mmapRead(fmt.Sprintf("./testdata/%s", benchfile))
			}
		})
	}
}
