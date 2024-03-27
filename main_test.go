package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"

	splitcsv "split_csv/core"
)

type Args struct {
	// gvars     map[string]interface{}
	csv_path  string
	file_part int
}
type Tests []struct {
	name     string
	args     Args
	expected string
}

func read_csv(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		err_slice := [][]string{{err.Error()}}
		return err_slice
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, _ := reader.ReadAll()
	return records
}

// Test_CsvNewLine Unit Testing Block
func Test_CsvNewLine(t *testing.T) {
	tests := Tests{
		{
			name: "conserved new line",
			args: Args{
				csv_path:  "./testing_data/test_new_line.csv",
				file_part: 2,
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitcsv.Split_Csv(&tt.args.csv_path, &tt.args.file_part)

			original_csv := read_csv(tt.args.csv_path)
			split_csv := read_csv("./split_data/test_new_line_part1.csv")

			if original_csv[1][0] != split_csv[1][0] {
				t.Error(fmt.Printf("Line %d: %v, %v\n", 0, original_csv[1], split_csv[1]))
			}
		})
	}
}

// Test_CsvLeading0 Unit Testing Block
func Test_CsvLeading0(t *testing.T) {
	tests := Tests{
		{
			name: "conserved leading 0",
			args: Args{
				csv_path:  "./testing_data/test_leading0.csv",
				file_part: 2,
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitcsv.Split_Csv(&tt.args.csv_path, &tt.args.file_part)

			original_csv := read_csv(tt.args.csv_path)
			split_csv := read_csv("./split_data/test_leading0_part1.csv")

			if original_csv[1][0] != split_csv[1][0] {
				t.Error(fmt.Printf("Line %d: %v, %v\n", 0, original_csv[1], split_csv[1]))
			}
		})
	}
}

// Test_CsvConserveQuote Unit Testing Block
func Test_CsvConserveQuote(t *testing.T) {
	tests := Tests{
		{
			name: "conserved \" and '",
			args: Args{
				csv_path:  "./testing_data/test_conserveQuote.csv",
				file_part: 2,
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitcsv.Split_Csv(&tt.args.csv_path, &tt.args.file_part)

			original_csv := read_csv(tt.args.csv_path)
			split_csv := read_csv("./split_data/test_conserveQuote_part1.csv")

			if original_csv[1][0] != split_csv[1][0] {
				t.Error(fmt.Printf("Line %d: %v, %v\n", 0, original_csv[1], split_csv[1]))
			}
		})
	}
}

// Test_CsvFileParts Unit Testing Block
func Test_CsvFileParts(t *testing.T) {
	tests := Tests{
		{
			name: "test file parts is equal as define",
			args: Args{
				csv_path:  "./testing_data/test_fileParts.csv",
				file_part: 10,
			},
			expected: "10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitcsv.Split_Csv(&tt.args.csv_path, &tt.args.file_part)

			files, _ := os.ReadDir("./split_data/")
			expected, err := strconv.Atoi(tt.expected)
			if err != nil {
				t.Error("Error in expected convertion.")
			}

			if len(files) != expected {
				t.Error("File parts is not equal as define.")
			}
		})
	}
}
