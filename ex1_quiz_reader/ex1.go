package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var DEFAULT_PROBLEM = "problems.csv"

func evaluateAnswer(score *int, userAns, realAns string) {
	if cleanUp(userAns) == cleanUp(realAns) {
		*score = *score + 1
	}
}

func cleanUp(str string) string {
	return strings.TrimSpace(strings.ToUpper(str))
}

func main() {

	problemPtr := flag.String("pSet", DEFAULT_PROBLEM, "problem csv file")
	timeLimit := flag.Int("t", 30, "Time Limit")
	flag.Parse()

	f, err := os.Open(*problemPtr)
	if err != nil {
		log.Fatal("Unable to read input file " + *problemPtr)
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	score, totalQuestion := 0, 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		fmt.Print("Q: " + record[0] + " A: ")
		answerChan := make(chan string)
		go func() {
			var userAnswer string
			totalQuestion += 1
			_, err = fmt.Scanln(&userAnswer)
			answerChan <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Println("")
			fmt.Println("time is up! ")
			return
		case userAnswer := <-answerChan:
			evaluateAnswer(&score, userAnswer, record[1])
		}

	}

	fmt.Println("Your final score is ", score, "/", totalQuestion)

}
