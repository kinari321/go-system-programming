package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()
	writer.Write([]string{"Alice", "22"})
	writer.Write([]string{"Bob", "31"})
}
