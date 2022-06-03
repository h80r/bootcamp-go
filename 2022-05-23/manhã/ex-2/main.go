package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./products.csv")
	if err != nil {
		panic(err)
	}

	parseLines(data)
}

func parseLines(data []byte) {
	byteLines := bytes.Split(data, []byte("\n"))

	for i := range byteLines {
		id, price, amount := parseLine(byteLines[i])
		fmt.Printf("%-6s\t\t%20s\t\t%6s\n", id, price, amount)
	}
}

func parseLine(line []byte) (string, string, string) {
	fields := bytes.Split([]byte(line), []byte(";"))

	return string(fields[0]), string(fields[1]), string(fields[2])
}
