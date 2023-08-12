package main

import (
	"flag"
	"fmt"

	splitcsv "split_csv/split_csv"
)

func main() {
	pathPtr := flag.String("p", "./testing_data/test_fileParts.csv", "path string")
	filePartPtr := flag.Int("fp", 10, "number of file splitting")
	flag.Parse()

	fileLength, err := splitcsv.Split_Csv(pathPtr, filePartPtr)

	fmt.Printf("Success creating: %d\nHave error: %v\n", fileLength, err)
}
