package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("3-9-3.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	example1, err := zipWriter.Create("example.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(example1, strings.NewReader("example1"))

	example2, err := zipWriter.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(example2, strings.NewReader("example2"))

}
