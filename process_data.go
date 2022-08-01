package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
)

type MPXRes struct {
	DataFileName string      `json:"dataFileName"`
	Data         []MPXRecord `json:"data"`
}

type MPXRecord struct {
	State string      `json:"state"`
	Cases json.Number `json:"cases"`
	Range string      `json:"range"`
	Date  string      `json:"date"`
}

func main() {
	input_file := os.Args[1]
	output_base := os.Args[2]

	// read the whole file at once
	mpx_bytes, err := ioutil.ReadFile(input_file)
	if err != nil {
		panic(err)
	}

	var mpx_records_raw MPXRes
	err = json.Unmarshal(mpx_bytes, &mpx_records_raw)
	if err != nil {
		panic(err)
	}

	// extract date from file name and attach it to all records
	r, _ := regexp.Compile(`(\d+)`)
	extracted_date := r.FindString(mpx_records_raw.DataFileName)

	for i := 0; i < len(mpx_records_raw.Data); i++ {
		mpx_records_raw.Data[i].Date = extracted_date
	}

	a, _ := json.Marshal(mpx_records_raw.Data)

	// write the whole body at once
	output_file := output_base + "-" + extracted_date + ".json"
	err = ioutil.WriteFile(output_file, a, 0644)
	if err != nil {
		panic(err)
	}
}
