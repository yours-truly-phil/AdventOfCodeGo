package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func Parse(path string) []int {
	file, _ := os.Open(path)
	reader := bufio.NewReader(file)
	ints, _ := ReadInts(reader)
	sort.Ints(ints)
	return ints
}

func twoNumsThatSumTo(ints []int, total int) (int, int) {
	for i := range ints {
		a := ints[i]
		if a >= total/2 {
			break
		}
		b := total - ints[i]
		idx := sort.SearchInts(ints, b)
		if ints[idx] == b {
			return a, b
		}
	}
	return -1, -1
}

func threeNumsThatSumTo(ints []int, total int) (int, int, int) {
	for i := range ints {
		a := ints[i]
		if a >= total/3 {
			break
		}
		diff := total - a
		for j := i + 1; j < len(ints); j++ {
			b := ints[j]
			c := diff - b
			if b >= diff/2 || c < b {
				break
			}
			idx := sort.SearchInts(ints, c)
			if ints[idx] == c {
				return a, b, c
			}
		}
	}
	return -1, -1, -1
}

func part1(path string) {
	ints := Parse(path)
	a, b := twoNumsThatSumTo(ints, 2020)
	fmt.Println("part1 =", a*b)
}

func part2(path string) {
	ints := Parse(path)
	a, b, c := threeNumsThatSumTo(ints, 2020)
	fmt.Println("part2 =", a*b*c)
}

func main() {
	part1("aoc2020/day1/day1.txt")
	part2("aoc2020/day1/day1.txt")
}
