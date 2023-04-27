package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

/*
Q3.4:zip ファイルをウェブサーバーからダウンロード
ウェブサーバーにブラウザでアクセスしたらファイルが zip ダウンロードされるようにする
*/
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	example1, err := zipWriter.Create("example.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(example1, strings.NewReader("example1"))

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
