package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {

	solver := getSolver(os.Args[1])
	testfile := getTestfile(os.Args[1])
	day := os.Args[1][0]
	part := os.Args[1][1]

	file, err := os.Open(testfile)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	problemChannel := make(chan string)
	solutionChannel := make(chan string)

	go solver(problemChannel, solutionChannel)

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
		log.Printf("Day %c Part %c '%s' in %s", day, part, solution, time.Since(start))
	}

}
