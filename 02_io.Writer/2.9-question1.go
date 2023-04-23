package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "数字: %d\n文字列: %s\n浮動小数点数: %f\n", 3, "hoge", 3.14)
}
