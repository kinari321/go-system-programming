package main

import (
	"io"
	"os"
)

/*
Q3.1:ファイルのコピー
古いファイル(old.txt)を新しいファイル(new.txt)にコピーしてみましょう。
*/

func main() {
	oldFile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	io.Copy(newFile, oldFile)
}
