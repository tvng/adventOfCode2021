package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	colNumber = 12
)

var (
	numOfLines     = 0
	sumPerCol      = make([]int, colNumber)
	gamma, epsilon int
)

func main() {
	file, err := os.Open("day3input.txt")
	if err != nil {
		fmt.Errorf("Error reading input: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for idx, bit := range scanner.Bytes() { // reads ascii
			sumPerCol[idx] += int(bit - 48)
		}
		numOfLines++
	}

	for _, colNumber := range sumPerCol {
		gamma = (gamma << 1)
		epsilon = (epsilon << 1)

		if colNumber >= (numOfLines / 2) { //majority
			gamma++
		} else {
			epsilon++
		}
	}
	fmt.Printf("Result is : %d * %d = %d", gamma, epsilon, gamma*epsilon)
}
