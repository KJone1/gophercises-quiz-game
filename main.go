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
		fmt.Print("Answer: ")
		reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		ans := strings.TrimRight(reader, "\r\n")
		if ans == scanned[1] {
			fmt.Printf("Good job %s the correct answer. \n-\n", ans)
		} else {
			fmt.Printf("try again\n")
		}
	}
}
