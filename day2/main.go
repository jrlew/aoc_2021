package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	var horizontal_position int = 0
	var depth int = 0

	instructions := getInput("./input")
	for _, instruction := range instructions {
		x, y := parseInstruction(instruction)

		horizontal_position += x
		depth += y
	}

	fmt.Println("Part One")
	fmt.Printf("Depth: %d\n", depth)
	fmt.Printf("Horizontal Position: %d\n", horizontal_position)
	fmt.Printf("Depth x Horizontal Position: %d\n\n", depth*horizontal_position)
}

func partTwo() {
	var horizontal_position int = 0
	var depth int = 0
	var aim int = 0

	instructions := getInput("./input")
	for _, instruction := range instructions {
		x, y := parseInstruction(instruction)

		aim += y
		depth += x * aim
		horizontal_position += x
	}

	fmt.Println("Part Two")
	fmt.Printf("Depth: %d\n", depth)
	fmt.Printf("Horizontal Position: %d\n", horizontal_position)
	fmt.Printf("Depth x Horizontal Position: %d\n\n", depth*horizontal_position)
}

func getInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
	}

	return data
}

func parseInstruction(instruction string) (int, int) {
	split := strings.Split(instruction, " ")

	var x int
	var y int

	switch split[0] {
	case "forward":
		x, _ = strconv.Atoi(split[1])
		y = 0
	case "down":
		x = 0
		y, _ = strconv.Atoi(split[1])
	case "up":
		x = 0
		y, _ = strconv.Atoi(split[1])
		y = y * -1
	}

	return x, y
}
