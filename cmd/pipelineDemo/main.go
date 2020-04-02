package main

import (
	"bufio"
	"fmt"
	"os"
	"pipeline/pipeline"
)

func main() {
	const filename = "large.in"
	const n = 100000000
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)

	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	p = pipeline.ReaderSource(reader, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func mergeDemo() {
	p := pipeline.Merge(pipeline.InMenSort(pipeline.ArraySource(3, 2, 6, 7, 4)), pipeline.InMenSort(pipeline.ArraySource(1, 2, 9, 5, 3)))
	// 用range发送方一定要close
	for v := range p {
		fmt.Println(v)
	}
}