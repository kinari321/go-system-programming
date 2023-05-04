package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	// リトライのループ
	for {
		var err error
		// 未コネクション. エラーでリトライ
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		// POSTメソッド
		request, err := http.NewRequest(
			"POST",
			"http://localhost:8888",
			strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		// レスポンス. タイムアウトでリトライ
		response, err := http.ReadResponse(
			bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
	}

}
