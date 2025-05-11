package main

import (
	"io"
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

func krnlRead(path string) int {
	data, err := os.ReadFile(path)
	check(err)
	acc := process(&data)
	return acc
}

func mmapRead(path string) int {
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
	acc := process(&data)
	return acc
}

func bachRead(path string) int {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	offset := 0
	acc := 0
	data := make([]byte, 4 * 1024)
	n, err := file.ReadAt(data, int64(offset))
	for err != io.EOF {
		acc += process(&data)
		offset += n
		n, err = file.ReadAt(data, int64(offset))
	}
	return acc
}