package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	correct_answers := 0
	csvFileName := flag.String("c", "questions.csv", "csv file or whatever")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Printf("Questions from: %s\n", *csvFileName)
	scanner := csv.NewReader(file)

	for {
		scanned, err := scanner.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question : %s\n", scanned[0])
		fmt.Print("Your answer: ")
		reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		ans := strings.TrimRight(reader, "\r\n")
		if ans == scanned[1] {
			correct_answers++
			fmt.Printf("Good job %s is the correct answer. \n-\n", ans)
		} else {
			fmt.Printf("Try again.\n-\n")
		}
	}
	defer fmt.Printf("Game Ended. \nNumber of correct answers: %d.\n", correct_answers)
}
