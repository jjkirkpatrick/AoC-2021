package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	horizontal := 0
	vertical := 0
	aim := 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//loop over the file and add each line to a slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//for each line in lines, split line and print the result
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		v, _ := strconv.Atoi(splitLine[1])
		switch splitLine[0] {
		case "forward":
			horizontal += v
			vertical += aim * v
		case "up":
			aim -= v
		case "down":
			aim += v
		}
	}

	fmt.Println("Distance: ", horizontal*vertical)

}
