package main

import (
	"os"
	"syscall"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func process(data *[]byte) int {
	countLetterB := 0
	for character := range *data {
		if character == 'b' {
			countLetterB += 1
		}
	}
	return countLetterB
}

func krnlRead(path string) {
	data, err := os.ReadFile(path)
	check(err)
	process(&data)
}

func mmapRead(path string) {
	fileInfo, err := os.Stat(path)
	check(err)

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	data, err := syscall.Mmap(int(file.Fd()), 0, int(fileInfo.Size()), syscall.PROT_READ, syscall.MAP_SHARED)
	check(err)
	defer func() {
		if err := syscall.Munmap(data); err != nil {
			panic(err)
		}
	}()
	process(&data)

}
