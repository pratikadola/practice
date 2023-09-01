package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	cycle, x, sum := 1, 1, 0

	for _, line := range fileLines {
		ins := strings.Split(line, " ")
		if ins[0] == "noop" {
			cycle++
		} else {
			val, _ := strconv.Atoi(ins[1])
			// fmt.Println(val)
			cycle++
			sum += checkSignalStrength(cycle, x)
			cycle++
			x += val
		}
		sum += checkSignalStrength(cycle, x)

	}
	fmt.Println(sum)
}

func checkSignalStrength(cycle int, val int) int {
	i := 20
	for i <= 220 {
		if i == cycle {
			// fmt.Println(cycle)
			return cycle * val
		}
		i += 40
	}
	return 0
}
