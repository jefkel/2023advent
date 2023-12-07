package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    var arr [3]string

	//var puzzleIn string = "puzzleIn.txt"
	var puzzleIn string = "sample.txt"
	readFile, err := os.Open(puzzleIn)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Get first line
	fileScanner.Scan()
    arr[0] := fileScanner.Text()
	
for fileScanner.Scan() {
	fileScanner.Text()
}
