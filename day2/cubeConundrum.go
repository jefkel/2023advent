package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const maxRed int = 12
	const maxGreen int = 13
	const maxBlue int = 14

	var cubeCounter map[string]int
	cubeCounter = make(map[string]int)
	var maxCubes map[string]int
	maxCubes = make(map[string]int)

	var puzzleIn string = "puzzleIn.day2"
	readFile, err := os.Open(puzzleIn)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Set Total validGames Counter
	var validGameTotal int = 0
	var powerSum int = 0

	for fileScanner.Scan() {
		var newln string = fileScanner.Text()
		var gameIndex int
		var junk string

		gameInfo := strings.Split(newln, ":")
		fmt.Sscanf(gameInfo[0], "%s %d", &junk, &gameIndex)
		//gameIndex := rxGetNum.FindAllString(gameInfo[0], -1)[0]

		var gameCheck bool = true
		maxCubes["red"] = 0
		maxCubes["green"] = 0
		maxCubes["blue"] = 0

		// loop on bag pulls
		for _, s := range strings.Split(gameInfo[1], ";") {
			// Reset colour counter hash
			cubeCounter["red"] = 0
			cubeCounter["green"] = 0
			cubeCounter["blue"] = 0
			// Add CubeCounts to appropriate colour
			for _, ss := range strings.Split(s, ",") {
				var cubes int
				var colour string
				fmt.Sscanf(ss, "%d %s", &cubes, &colour)
				cubeCounter[colour] += cubes
				//update max for each colour
				if cubeCounter[colour] > maxCubes[colour] {
					maxCubes[colour] = cubeCounter[colour]
				}
			}
			if cubeCounter["red"] > maxRed || cubeCounter["green"] > maxGreen || cubeCounter["blue"] > maxBlue {
				gameCheck = false
			}
		}
		// Check Validity of game
		if gameCheck {
			validGameTotal += gameIndex
		}
		powerSum += (maxCubes["red"] * maxCubes["green"] * maxCubes["blue"])

	}

	fmt.Println("\nGameIndex Total:", validGameTotal)
	fmt.Println("\nPower Sum Total:", powerSum)
}
