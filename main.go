package main

import "fmt"

func main() {
	n := 20

	results := make(chan Board)
	board := MakeBoard(n)

	go Crawl(board, results)
	for result := range results {
		fmt.Println(result.ToString())
		fmt.Println("")
	}
}
