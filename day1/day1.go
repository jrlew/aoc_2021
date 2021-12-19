package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []int = make([]int, 0)
	// var prev int = -42
	var increase int = 0
	var decrease int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		data = append(data, curr)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	datalen := len(data)
	for i := 0; i < datalen-3; i++ {
		currFrame := data[i] + data[i+1] + data[i+2]
		nextFrame := data[i+1] + data[i+2] + data[i+3]

		// fmt.Printf("i: %d - currFrame: %d - nextFrame: %d - increase: %t\n", i, currFrame, nextFrame, nextFrame > currFrame)

		if currFrame > nextFrame {
			decrease++
		} else if currFrame < nextFrame {
			increase++
		} else {
			fmt.Printf("i: %d - no change\n", i)
		}
	}

	// x x x x x x x

	// if prev != -42 {
	// 	if prev > curr {
	// 		decrease++
	// 	} else {
	// 		increase++
	// 	}
	// }
	// prev = curr

	fmt.Println(increase)
	fmt.Println(decrease)
}
