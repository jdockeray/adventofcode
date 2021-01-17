package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func MinMax(array []int) (int, int) {
	var max = array[0]
	var min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func FindWeakness(search int, numbers []int) int {
	sum := 0
	length := 0
	for i, n := range numbers {
		sum = sum + n
		length = length + 1
		for sum > search {
			length = length - 1
			sum = sum - numbers[i-length]
		}
		if sum == search {
			lower := i - (length - 1)
			upper := i + 1
			min, max := MinMax(numbers[lower:upper])
			return min + max
		}

	}
	return 0
}

func FindSum(number int, numbers []int) bool {
	sums := make(map[int]bool)
	for index, left := range numbers {
		for i := index + 1; i < len(numbers); i++ {
			right := numbers[i]
			sums[left+right] = true
		}
	}
	exists := sums[number]

	return exists
}

func Find(numbers []int, blockSize int) int {
	for i := blockSize; i < len(numbers); i++ {
		active := numbers[i]
		if !FindSum(active, numbers[i-blockSize:i]) {
			return active
		}
	}
	return 0
}

func main() {
	resp, err := http.Get("http://127.0.0.1:8081/data/day9.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	numbers := make([]int, 0, 10)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		numbers = append(numbers, n)
	}

	fmt.Println(FindWeakness(Find(numbers, 25), numbers))

}
