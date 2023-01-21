package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func readCSV() [][]string {
	f, err := os.Open(file)
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

var file string
var timer int

func init() {
	var (
		defaultFile     = "problems.csv"
		fileDescription = "Change the path to the quiz CSV, the default is \"problems.csv\""
	)
	flag.StringVar(&file, "file", defaultFile, fileDescription)
	flag.StringVar(&file, "f", file, fileDescription+" (shorthand)")

	var (
		defaultTimer     = 30
		timerDescription = "Change the Timer duration, the default is set to 30 seconds"
	)
	flag.IntVar(&timer, "timer", defaultTimer, timerDescription)
	flag.IntVar(&timer, "t", timer, timerDescription+" (shorthand)")

	flag.Parse()
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
