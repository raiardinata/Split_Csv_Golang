package main

import (
	"flag"
	"fmt"

	split_core "split_csv/core"
)

func main() {
	pathPtr := flag.String("p", "", "Fill with your string path")
	filePartPtr := flag.Int("fp", 0, "Fill with how many file you want to split")
	flag.Parse()

	fileLength, err := split_core.Split_Csv(pathPtr, filePartPtr)

	fmt.Printf("Success creating: %d\nHave error: %v\n", fileLength, err)
}
