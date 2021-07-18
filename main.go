package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

const size = 1024 * 1024 * 100

func read(accumulatorSize, readBufSize int) {

	buf := make([]byte, readBufSize)

	f, _ := os.Open("dump.txt")
	reader := bufio.NewReader(f)
	var n int
	var err error
	var bufWriteError error
	//should we create a buffer of fixed size?
	start := time.Now()
	//allocate 1mb buffer
	buffer := bytes.NewBuffer(make([]byte, accumulatorSize))
	buffer.Reset()

	for {
		n, err = reader.Read(buf)
		if n == 0 {
			if err == io.EOF {
				err = nil
				break
			} else if err != nil {
				break
			}
		}
		_, bufWriteError = buffer.Write(buf[:n])
	}
	if bufWriteError != nil {
		fmt.Println(bufWriteError)
	}
	if err != nil {
		fmt.Println(err)
	}

	took := time.Since(start)
	fmt.Printf("Read: read buffer size %d, write buffer size %d, took %v %d\n", readBufSize, accumulatorSize, took, len(buffer.Bytes()))
}

func write(dumpSize int) {
	s := make([]byte, dumpSize)
	buf := bytes.NewBuffer(s)
	s[0] = 'a'
	s[1] = 'x'
	s[2] = 'o'
	s[len(s)-1] = '\n'
	f, _ := os.Create("dump.txt")
	fw := bufio.NewWriter(f)
	var n int
	var err error
	start := time.Now()
	n, err = fw.Write(buf.Bytes())
	took := time.Since(start)
	fmt.Printf("Write took %v\n", took)
	fw.Flush()
	fmt.Printf("Written %d, size %d, %v\n", n, len(s), err)
}

func main() {
	//adjust according to your system page size
	PAGESIZE := 4
	GB := 1024 * 1024 * 1024
	MB := 1024 * 1024
	KB := 1024
	//write and read 1GB using 1GB buff size
	write(1 * GB)
	read(1*GB, 1*GB)
	//read 1Gb using 1MB buffer
	read(1*MB, 1*GB)
	//read 1Gb using 1KB buffer
	read(1*KB, 1*GB)
	//page size read buffer
	read(1*GB, PAGESIZE*KB)

	//write and read 1MB using 1MB buff size
	write(1 * MB)
	read(1*MB, 1*MB)
	//read 1Gb using 1KB buffer
	read(1*KB, 1*MB)

	//page size read buffer
	read(1*MB, PAGESIZE*KB)

	write(10 * MB)
	read(10*MB, 1*MB)
	//read 1Gb using 1KB buffer
	read(1*KB, 1*MB)

	//page size read buffer
	read(10*MB, PAGESIZE*KB)
}
