package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	previous := 0
	for scanner.Scan() {
		if previous == 0 {
			previous, _ = strconv.Atoi(scanner.Text())
		} else {
			currentValue, _ := strconv.Atoi(scanner.Text())
			if currentValue > previous {
				count++
			}
			previous = currentValue
		}
	}
	fmt.Println(count)
}
