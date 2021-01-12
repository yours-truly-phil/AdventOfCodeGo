package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(path string) int {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(bufio.ScanLines)
	var countValid = 0
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, " ")
		nums := strings.Split(parts[0], "-")
		low, _ := strconv.Atoi(nums[0])
		high, _ := strconv.Atoi(nums[1])
		char := string(parts[1][0])
		freq := strings.Count(parts[2], char)
		if freq >= low && freq <= high {
			countValid++
		}
	}
	return countValid
}

func part2(path string) {

}

func main() {
	fmt.Println(part1("aoc2020/day2/day2.txt"))
}
