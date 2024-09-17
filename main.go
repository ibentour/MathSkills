package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Average(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return int(math.Round(float64(sum) / float64(len(numbers))))
}

func Median(numbers []int) int {
	ln := len(numbers)
	sort.Ints(numbers)

	if len(numbers)%2 != 0 {
		return int(math.Round(float64(numbers[ln/2])))
	}

	mid := numbers[(ln/2)-1 : (ln/2)+1]
	return int(math.Round(float64(Average(mid))))
}

func Variance(numbers []int, average int) int {
	var sumSquares float64
	for _, num := range numbers {
		sumSquares += math.Pow(float64(num) - float64(average), 2)
	}
	return int(math.Round(sumSquares / float64(len(numbers))))
}

func Standard(variance int) int {
	return int(math.Round(math.Sqrt(float64(variance))))
}

func main() {
	file := os.Args[1:]
	if len(file) != 1 || file[0] != "data.txt" {
		fmt.Println("[USAGE]! : go run your-program.go data.txt")
		return
	}

	data, err := os.ReadFile(file[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	numStr := strings.Split(string(data), "\n")
	numbers := []int{}
	for _, line := range numStr {
		num, err := strconv.Atoi(line)
		if err != nil {
			if line == "" {
				continue
			}
			fmt.Println("[Error]: ", err)
			return
		}
		numbers = append(numbers, num)
	}

	average := Average(numbers)
	median := Median(numbers)
	variance := Variance(numbers, average)
	standard := Standard(variance)
	fmt.Printf("Average: %d\n", average)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", variance)
	fmt.Printf("Standard Deviation: %d\n", standard)
}
