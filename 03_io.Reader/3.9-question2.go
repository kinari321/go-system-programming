package main

import (
	"crypto/rand"
	"io"
	"os"
)

/*
Q3.2:テスト用の適当なサイズのファイルを作成
ファイルを作成してランダムな内容で埋めてみましょう。
crypto/rand パッケージをimportする
*/

func main() {
	file, err := os.Create("cryptoRandom.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := rand.Reader
	limitReader := io.LimitReader(reader, 1024)

	buffer := make([]byte, 1024)
	_, err = io.ReadFull(limitReader, buffer)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(buffer)
	if err != nil {
		panic(err)
	}
}
