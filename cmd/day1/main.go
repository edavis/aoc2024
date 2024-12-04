package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(handle *os.File) {
	var left []int
	var right []int

	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])

		left = append(left, l)
		right = append(right, r)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
		return
	}

	slices.Sort(left)
	slices.Sort(right)

	var delta float64

	for idx, lv := range left {
		rv := right[idx]
		delta += math.Abs(float64(lv) - float64(rv))
	}

	fmt.Printf("%f\n", delta)
}

func part2(handle *os.File) {
	var left []int
	var right []int

	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])

		left = append(left, l)
		right = append(right, r)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
		return
	}

	slices.Sort(left)
	slices.Sort(right)

	var score int
	for _, lv := range left {
		var mult int

		for _, rv := range right {
			if rv > lv {
				break
			}

			if lv == rv {
				mult += 1
			}
		}
		score += lv * mult
	}

	fmt.Printf("%d\n", score)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("missing input data\n")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return
	}
	defer file.Close()

	part1(file)
	file.Seek(0, 0)
	part2(file)
}
