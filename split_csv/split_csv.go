package splitcsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func read_csv(path string) ([][]string, []string) {
	// open file
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	var header []string
	// var concatenate string

	// read csv records values using csv.Reader
	var records [][]string

	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	i := 0
	for {

		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		// fmt.Printf("%+v\n", rec)

		if i == 0 {
			header = append(header, rec...)
			i++
		}

		records = append(records, rec)
	}
	return records, header
}

func split_len_part(length int, filePart int) int {
	return int(math.Ceil(float64(length) / float64(filePart)))
}

func writcsv(total_len, split_len int, fileNameCsv string, filePart int, records [][]string, header []string) error {
	j := 1
	beginning := 0
	end := split_len
	filename, _ := getFileName(fileNameCsv)

	for i := filePart; i != 0; i-- {
		csvFile, err := os.Create("./split_data/" + filename + "_part" + strconv.Itoa(j) + ".csv")
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("failed creating file: %s", err))
		}
		defer csvFile.Close()

		w := csv.NewWriter(csvFile)
		defer w.Flush()

		// Using WriteAll
		var data [][]string

		// make header after part 1
		if i != filePart {
			data = append(data, header)
		}
		if total_len < split_len {
			split_len = split_len - (split_len - total_len)
		}
		if j != 1 {
			data = append(data, records[(beginning+1):(end+1)]...)
		} else {
			data = append(data, records[beginning:(end+1)]...)

		}
		w.WriteAll(data)

		beginning = end
		end = beginning + split_len
		j++
	}

	return nil
}

// getFileName extracts name and extension from path
func getFileName(path string) (string, string) {
	filenameArr := strings.Split(filepath.Base(path), ".")
	if len(filenameArr) == 2 {
		return filenameArr[0], filenameArr[1]
	}

	return filenameArr[0], ""
}

func Split_Csv(path *string, filePart *int) (int, error) {
	records, header := read_csv(*path)
	split_length := split_len_part(len(records), int(*filePart))

	os.RemoveAll("./split_data/")
	err2 := os.MkdirAll("split_data/", os.ModePerm)
	if err2 != nil {
		msg := fmt.Sprintf("Couldn't create directory 'split_data/' : %v", err2)
		log.Fatalf(msg)

	}

	err := writcsv(len(records), split_length, filepath.Base(*path), *filePart, records, header)
	files, _ := os.ReadDir("./split_data/")
	return len(files), err
}
