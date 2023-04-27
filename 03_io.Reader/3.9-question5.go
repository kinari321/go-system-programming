package main

import (
	"io"
	"os"
	"strings"
)

/*
Q3.5:CopyN関数を使ってみる
*/

func main() {
	str := "Hello, world!"
	r := strings.NewReader(str)
	io.CopyN(os.Stdout, r, int64(5))
}
