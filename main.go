package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//startTime := time.Now() // Record the start time

	file, err := os.Open("pwned-passwords.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 32*1024*1024), bufio.MaxScanTokenSize) // Set buffer to 32MB

	maxSum := 0
	var maxSumHashes []string

	for scanner.Scan() {
		line := scanner.Text()
		idx := strings.Index(line, ":")
		hash := line[:idx]

		hashDigitsOnly := strings.Map(func(r rune) rune {
			if r >= '0' && r <= '9' {
				return r
			}
			return -1
		}, hash)

		digitSum := sumDigits(hashDigitsOnly)

		if digitSum > maxSum {
			maxSum = digitSum
			maxSumHashes = []string{hash}
		} else if digitSum == maxSum {
			maxSumHashes = append(maxSumHashes, hash)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	for _, hash := range maxSumHashes {
		fmt.Printf("%s:%d\n", hash, maxSum)
	}

	//endTime := time.Now()
	//executionTime := endTime.Sub(startTime)
	//fmt.Println("Execution Time:", executionTime)
}

func sumDigits(s string) int {
	sum := 0
	for _, c := range s {
		if c > '0' && c <= '9' {
			sum += int(c - '0')
		}
	}
	return sum
}
