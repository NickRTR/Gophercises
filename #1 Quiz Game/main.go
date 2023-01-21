package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var file string
var duration int

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
	flag.IntVar(&duration, "timer", defaultTimer, timerDescription)
	flag.IntVar(&duration, "t", defaultTimer, timerDescription+" (shorthand)")

	flag.Parse()
}

func main() {
	correct := 0
	data := readCSV()

	timer := time.NewTimer(time.Duration(*&duration) * time.Second)

problemLoop:
	for i, p := range data {
		fmt.Printf("Problem #%d: %s = ", i+1, p[0])
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == p[1] {
				correct++
			}
		}
	}

	fmt.Printf("\nResult: %d answers correct, %d answers incorrect.\n", correct, len(data)-correct)
}

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
