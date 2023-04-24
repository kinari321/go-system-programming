package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	str := "Hello, world!"
	r := strings.NewReader(str)
	io.CopyN(os.Stdout, r, int64(5))
}
