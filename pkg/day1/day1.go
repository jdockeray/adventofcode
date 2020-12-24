package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func paired(find int, compare int, numbers []int) int {
	for _, n := range numbers {
		if n+compare == find {
			return n
		}
	}
	return 0
}

func One() {
	file, err := os.Open("input/data.txt")

	numbers := make([]int, 0, 10)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		i, _ := strconv.Atoi(str)

		n := paired(2020, i, numbers)

		if n != 0 {
			fmt.Printf("numbers found they are %d %d and the answer is %d \n", n, i, n*i)
			return
		}
		numbers = append(numbers, i)
	}

	fmt.Println("no numbers found")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func Multiply(compare int, numbers []int) [][]int {
	pairs := make([][]int, 0, 10)
	for _, n := range numbers {
		pairs = append(pairs, []int{compare, n})
	}
	return pairs
}

func Two() {
	file, err := os.Open("input/data.txt")

	numbers := make([]int, 0, 10)
	pairs := make([][]int, 0, 10)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		i, _ := strconv.Atoi(str)
		if len(numbers) > 1 {
			pairs = append(pairs, Multiply(i, numbers)...)

			for _, pair := range pairs {
				x, y := pair[0], pair[1]
				total := x + y

				p := paired(2020, total, numbers)

				if p != 0 {
					fmt.Printf("numbers found they are %d %d %d and the answer is %d", x, y, p, x*y*p)
					return
				}

			}

		}
		numbers = append(numbers, i)
	}

	fmt.Println("no numbers found")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	One()
	fmt.Println("===============================")
	fmt.Println("===============================")
	Two()
}
