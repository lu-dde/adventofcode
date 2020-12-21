package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {

	solve := getProblem(os.Args[1])

	if solve == nil {
		log.Fatalf("couldn't find problem: %s", os.Args[1])
	}

	file, err := os.Open(solve.InputFile)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	input := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	problemChannel := make(chan string, len(input))
	solutionChannel := make(chan string, 1)

	for _, i := range input {
		problemChannel <- i
	}

	start := time.Now()

	go solve.Solve(problemChannel, solutionChannel)

	close(problemChannel)

	solution, ok := <-solutionChannel
	close(solutionChannel)
	file.Close()

	if ok {
		log.Printf("2020 Day %s Part %s '%s' in %s", solve.Day, solve.Part, solution, time.Since(start))
	}

}
