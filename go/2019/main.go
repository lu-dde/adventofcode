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

	log.Printf("time taken: %s", time.Since(start))

	if ok {
		log.Print(solution)
	}

}
