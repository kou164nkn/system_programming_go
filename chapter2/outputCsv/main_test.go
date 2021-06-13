package main_test

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"

	"github.com/kou164nkn/system_programming_go/chapter2/outputCsv"
)

func TestCreateCsvAndContent(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string

		fileName       string
		records        [][]string
		expectFileName string
		expectErr      bool
	}{
		{
			name:     "success",
			fileName: "./testdata/actualFile.csv",
			records: [][]string{
				{"name", "test1", "test2", "total"},
				{"Tom", "70", "80", "150"},
				{"Mary", "90", "85", "175"},
			},
			expectFileName: "./testdata/expectFile.csv",
			expectErr:      false,
		},
	}

	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			switch err := main.CreateCsvAndContent(tt.fileName, tt.records); true {
			case err == nil && tt.expectErr:
				t.Error("expected error did not occur")
			case err != nil && !tt.expectErr:
				t.Error("unexpected error: ", err)
			}

			expectData := readCsv(t, tt.expectFileName)
			actualData := readCsv(t, tt.fileName)

			if !reflect.DeepEqual(expectData, actualData) {
				t.Errorf("CreateCsvAndContent want %v but %v", expectData, actualData)
			}
		})
		removeActualFile(t, tt.fileName)
	}
}

func readCsv(t *testing.T, fileName string) [][]string {
	t.Helper()
	file, err := os.Open(fileName)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	return records
}

func removeActualFile(t *testing.T, fileName string) {
	t.Helper()
	if err := os.Remove(fileName); err != nil {
		t.Fatal(err)
	}
}
