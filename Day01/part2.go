package main

import (
	"fmt"
	"os"
)

func main() {
	//setup
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var lines []int
	for {
		var line int
		_, err := fmt.Fscanf(file, "%d\n", &line)
		if err != nil {
			break
		}

		lines = append(lines, line)
	}

	count := 0
	previous := 0
	for i := 0; i < len(lines); i++ {
		window := lines[i : i+3]
		if len(window) == 3 {
			if previous == 0 {
				previous = window[0] + window[1] + window[2]
			} else if window[0]+window[1]+window[2] > previous {
				count++
			}
			previous = window[0] + window[1] + window[2]
		}
	}

	fmt.Println(count)
}
