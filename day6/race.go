package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var puzzleIn string = "puzzleIn.txt"
	//var puzzleIn string = "sample.txt"
	readFile, err := os.Open(puzzleIn)
	if err != nil {
		fmt.Println(err)
	}
	//	var lines []string
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	times := strings.Split(fileScanner.Text(), ": ")
	var r [2][5]int
	for i, val := range strings.Fields(times[1]) {
		r[0][i], _ = strconv.Atoi(val)
		fmt.Println(r[0][i])
	}

	fileScanner.Scan()
	dist := strings.Split(fileScanner.Text(), ": ")
	for i, val := range strings.Fields(dist[1]) {
		r[1][i], _ = strconv.Atoi(val)
	}

	w := 1

	for i := 0; i < 4; i++ {
		n := 0
		for a := 1; a < r[0][i]; a++ {
			y := (r[0][i] - a) * a
			if y > r[1][i] {
				n++
			}
		}
		w *= n
		fmt.Println("Time:", r[0][i], "num", n)
	}
	fmt.Println("Answer: ", w)

	p2time, _ := strconv.Atoi(strings.Replace(times[1], " ", "", -1))
	p2dist, _ := strconv.Atoi(strings.Replace(dist[1], " ", "", -1))
	var nn int64 = 0
	for a := 1; a < p2time; a++ {
		y := (p2time - a) * a
		if y > p2dist {
			nn++
		}
	}
	fmt.Print("Part2: ", nn)

}
