package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func linesScanner(path string) *bufio.Scanner {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(bufio.ScanLines)
	return scanner
}

func parseLine(line string) (int, int, uint8, string) {
	parts := strings.Split(line, " ")
	nums := strings.Split(parts[0], "-")
	low, _ := strconv.Atoi(nums[0])
	high, _ := strconv.Atoi(nums[1])
	char := parts[1][0]
	pw := parts[2]
	return low, high, char, pw
}

func part1(path string) int {
	scanner := linesScanner(path)
	var countValid = 0
	for scanner.Scan() {
		s := scanner.Text()
		low, high, char, pw := parseLine(s)
		freq := strings.Count(pw, string(char))
		if freq >= low && freq <= high {
			countValid++
		}
	}
	return countValid
}

func part2(path string) int {
	scanner := linesScanner(path)
	var countValid = 0
	for scanner.Scan() {
		s := scanner.Text()
		low, high, char, pw := parseLine(s)
		if (pw[low-1] == char || pw[high-1] == char) && pw[low-1] != pw[high-1] {
			countValid++
		}
	}
	return countValid
}

func main() {
	fmt.Println(part1("aoc2020/day2/day2.txt"))
	fmt.Println(part2("aoc2020/day2/day2.txt"))
}
