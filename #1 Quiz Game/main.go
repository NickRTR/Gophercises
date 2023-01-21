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

func ask(question, answer string) bool {
	fmt.Printf("%s:", question)
	var input string
	fmt.Scan(&input)
	return input == answer
}

func main() {
	data := readCSV()

	correct := 0
	incorrect := 0
	for _, quiz := range data {
		if ask(quiz[0], quiz[1]) {
			correct++
		} else {
			incorrect++
		}
	}

	fmt.Printf("\nResult: %d answers correct, %d answers incorrect.\n", correct, incorrect)
}
