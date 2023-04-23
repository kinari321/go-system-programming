package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}

	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	file, err := os.Create("./output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	multiWriter := io.MultiWriter(gzipWriter, file)

	encoder := json.NewEncoder(multiWriter)
	err = encoder.Encode(source)
	if err != nil {
		panic(err)
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
