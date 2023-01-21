package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCSV() [][]string {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func main() {
	data := readCSV()
	fmt.Println(data)
}
