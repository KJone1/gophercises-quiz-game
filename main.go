package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	q := []string{}
	a := []string{}
	correct_answers := 0

	csvFileName := flag.String("c", "questions.csv", "csv file or whatever")
	flag.Parse()

	file, err := os.ReadFile(*csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Questions from: %s\n", *csvFileName)
	newline_separated := strings.Split(string(file), "\r\n")
	for index := range newline_separated {
		comma_separated := strings.Split(newline_separated[index], ",")
		q = append(q, comma_separated[0])
		a = append(a, comma_separated[1])
	}

	for index, scanned := range q {

		fmt.Printf("Question : %s\n", scanned)
		fmt.Print("Your answer: ")
		userin, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		user_answer := strings.TrimRight(userin, "\r\n")
		if user_answer == a[index] {
			correct_answers++
			fmt.Fprintf(os.Stdout, "\033[32mGood job %s is the correct answer. \n-\n\033[0m", user_answer)
		} else {
			for {
				fmt.Fprintf(os.Stdout, "\033[31mWrong Answer try again.\n-\n\033[0m")
				fmt.Printf("Question : %s\n", scanned)
				fmt.Print("Your answer: ")
				userin, err := bufio.NewReader(os.Stdin).ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				user_answer := strings.TrimRight(userin, "\r\n")
				if user_answer == a[index] {
					correct_answers++
					fmt.Fprintf(os.Stdout, "\033[32mGood job %s is the correct answer. \n-\n\033[0m", user_answer)
					break
				}
			}
		}
	}
	fmt.Printf("Game Ended. \nNumber of correct answers: %d.\n", correct_answers)
}
