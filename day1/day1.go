package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var calibFile string = "in-calib.txt"
	var partATotal int = 0
	var partBTotal int = 0

	ForwardKey := []string{
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	}

	ReverseKey := []string{
		"eno", "1",
		"owt", "2",
		"eerht", "3",
		"ruof", "4",
		"evif", "5",
		"xis", "6",
		"neves", "7",
		"thgie", "8",
		"enin", "9",
	}
	rxDigitString := "\\d"
	digits, _ := regexp.Compile(rxDigitString)

	readFile, err := os.Open(calibFile)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		// Create Reverse String
		var sb strings.Builder
		for index := len(fileScanner.Text()) - 1; index >= 0; index-- {
			sb.WriteByte(fileScanner.Text()[index])
		}

		// Replace all spelled numbers with digits
		replacer := strings.NewReplacer(ForwardKey...)
		partBNumStr := replacer.Replace(fileScanner.Text())

		// Replace all spelled numbers with digits (reverse spelling)
		replaceRev := strings.NewReplacer(ReverseKey...)
		partBNumStrRev := replaceRev.Replace(sb.String())

		// Get first digits
		paFirst := digits.FindAllString(fileScanner.Text(), -1)[0]
		pbFirst := digits.FindAllString(partBNumStr, -1)[0]

		// Get last digits
		paLast := digits.FindAllString(sb.String(), -1)[0]
		pbLast := digits.FindAllString(partBNumStrRev, -1)[0]

		//fmt.Println(fileScanner.Text(), "[", paFirst, paLast, "]  (", pbFirst, pbLast, ")")

		// Add to running totals
		partACal, _ := strconv.Atoi(paFirst + paLast)
		partBCal, _ := strconv.Atoi(pbFirst + pbLast)
		partATotal += partACal
		partBTotal += partBCal
	}
	readFile.Close()

	fmt.Println("----------")
	fmt.Println(" Part A: Calibration Sum: ", partATotal)
	fmt.Println(" Part B: Calibration Sum: ", partBTotal)

}
