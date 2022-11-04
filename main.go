package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func ask_question(question string) string {
	fmt.Printf("Question : %s= ?\n", question)
	fmt.Print("Your answer: ")
	userin, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	user_answer := strings.TrimRight(userin, "\r\n")
	return user_answer
}
func main() {
	q := []string{}
	a := []string{}
	correct_answers := 0

	csv_file_name := flag.String("c", "questions.csv", "csv file or whatever")
	max_retries := flag.Int("r", 3, "number of retries on wrong answer")
	flag.Parse()

	file, err := os.ReadFile(*csv_file_name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Questions from: %s\n", *csv_file_name)
	newline_separated := strings.Split(string(file), "\r\n")
	for index := range newline_separated {
		comma_separated := strings.Split(newline_separated[index], ",")
		q = append(q, comma_separated[0])
		a = append(a, comma_separated[1])
	}

	for index, scanned := range q {
		user_answer := ask_question(scanned)
		if user_answer == a[index] {
			correct_answers++
			fmt.Fprintf(os.Stdout, "\033[32mGood job %s is the correct answer. \n\033[0m-\n", user_answer)
		} else {
			retries := 0
			for retries < *max_retries {
				retries++
				fmt.Fprintf(os.Stdout, "\033[31mWrong Answer %d tries remaining.\n\033[0m-\n", 3-retries)
				user_answer := ask_question(scanned)
				if user_answer == a[index] {
					correct_answers++
					fmt.Fprintf(os.Stdout, "\033[32mGood job %s is the correct answer. \n\033[0m-\n", user_answer)
					break
				}
			}
			if retries == *max_retries {
				fmt.Fprintf(os.Stdout, "\033[31mWrong Answer.\n\033[0m-\n")
			}
		}
	}
	fmt.Printf("Game Ended. \nNumber of correct answers: %d.\n", correct_answers)
}
