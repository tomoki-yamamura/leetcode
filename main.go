package main

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	writer := os.Stdout
	reader := strings.NewReader("123456789012345678901234567890")
	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	myCopyN(writer, reader, size)
}

func myCopyN(w io.Writer, r io.Reader, length int) {
	reader := io.LimitReader(r, int64(length))
	io.Copy(w, reader)
}