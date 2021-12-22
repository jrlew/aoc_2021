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
}

func partOne() {
	diags := getInput("./input")
	mostCommonBits := getMostCommonBits(diags)
	fmt.Printf("%v\n", mostCommonBits)
	gamma, epsilon := parseRates(mostCommonBits)

	fmt.Printf("%d", gamma*epsilon)
}

func parseDiags(diags []string) [][]int {
	var parsedDiags [][]int = [][]int{}
	for _, diag := range diags {
		split := strings.Split(diag, "")
		var diag []int = []int{}
		for _, char := range split {
			i, _ := strconv.Atoi(char)
			diag = append(diag, i)
		}
		parsedDiags = append(parsedDiags, diag)
	}
	return parsedDiags
}

func getMostCommonBits(diags []string) map[string]map[string]int {
	bits := map[string]map[string]int{
		"0":  map[string]int{"0": 0, "1": 0},
		"1":  map[string]int{"0": 0, "1": 0},
		"2":  map[string]int{"0": 0, "1": 0},
		"3":  map[string]int{"0": 0, "1": 0},
		"4":  map[string]int{"0": 0, "1": 0},
		"5":  map[string]int{"0": 0, "1": 0},
		"6":  map[string]int{"0": 0, "1": 0},
		"7":  map[string]int{"0": 0, "1": 0},
		"8":  map[string]int{"0": 0, "1": 0},
		"9":  map[string]int{"0": 0, "1": 0},
		"10": map[string]int{"0": 0, "1": 0},
		"11": map[string]int{"0": 0, "1": 0},
	}

	for _, diag := range diags {
		for j, bit := range diag {
			jndex := strconv.Itoa(j)
			bits[jndex][string(bit)] = bits[jndex][string(bit)] + 1
		}
	}

	return bits
}

func parseRates(bits map[string]map[string]int) (gamma int64, epsilon int64) {
	var gammaRate string = ""
	var epsilonRate string = ""

	for _, i := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"} {
		fmt.Printf("%v\n", i)
		fmt.Printf("%v\n\n", bits[i])

		if bits[i]["0"] > bits[i]["1"] {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
	}
	fmt.Printf("%d\n", gammaRate)
	fmt.Printf("%d\n", epsilonRate)
	gammaInt, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilonRate, 2, 64)
	fmt.Printf("%d\n", gammaInt)
	fmt.Printf("%d\n", epsilonInt)

	return gammaInt, epsilonInt
}

func strToBinary(s string, base int) []byte {

	var b []byte

	for _, c := range s {
		b = strconv.AppendInt(b, int64(c), base)
	}

	return b
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
