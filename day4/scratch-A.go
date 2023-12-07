package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func partA() {
	var puzzleIn string = "puzzleIn.txt"
	//var puzzleIn string = "sample.txt"
	readFile, err := os.Open(puzzleIn)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var winners map[string]int
	var answer int = 0
	for fileScanner.Scan() {
		var count int = 0
		var points int = 0
		winners = make(map[string]int)
		line := strings.SplitAfter(fileScanner.Text(), ": ")
		scratch := strings.Split(line[1], " | ")
		//fmt.Println("wins:", scratch[0])
		//fmt.Println("vals:", scratch[1])
		for _, w := range strings.Fields(scratch[0]) {
			winners[w] = 0
		}
		for _, s := range strings.Fields(scratch[1]) {
			_, ok := winners[s]
			if ok {
				count++
			}
		}
		if count > 0 {
			points += IntPow(2, count-1)
		}
		clear(winners)
		answer += points
	}
	fmt.Println("PartA Answer:", answer)
}

func partB() {
	fmt.Println("PartB:")
	var puzzleIn string = "puzzleIn.txt"
	//var puzzleIn string = "sample.txt"
	readFile, err := os.Open(puzzleIn)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

}

func main() {
	partA()
	partB()
}
