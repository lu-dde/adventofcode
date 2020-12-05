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

	problemChannel := make(chan string, 2000)
	solutionChannel := make(chan string)

	go solve.Solve(problemChannel, solutionChannel)

	start := time.Now()

	for scanner.Scan() {
		line := scanner.Text()
		problemChannel <- line
	}
	close(problemChannel)

	file.Close()

	solution, ok := <-solutionChannel
	close(solutionChannel)

	if ok {
		log.Printf("2020 Day %s Part %s '%s' in %s", solve.Day, solve.Part, solution, time.Since(start))
	}

}
