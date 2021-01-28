package day10

import (
	"bufio"
	"fmt"
	"github.com/jdockeray/adventofcode/pkg/tree"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

func Calculate(adapters []int) int {
	sort.Ints(adapters)
	differences := struct {
		one   int
		three int
	}{
		one:   0,
		three: 0,
	}
	compare := 0
	for index, adapter := range adapters {
		difference := adapter - compare
		if difference == 1 {
			differences.one += 1
		} else {
			differences.three += 1
		}
		compare = adapters[index]
	}
	differences.three += 1 // built-in voltage adapter

	return differences.one * differences.three
}

func Calculate2(adapters []int) int {
	sort.Ints(adapters)
	return tree.CountLeaves(tree.BuildTree(adapters))
}

func Day10() {
	resp, err := http.Get(os.Getenv("FILE_SERVER") + "/day10.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	adapters := make([]int, 0, 10)

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		adapter, err := strconv.Atoi(scanner.Text())
		if err == nil {
			adapters = append(adapters, adapter)
		}
	}
	//fmt.Println(Calculate(adapters))
	fmt.Println("========")
	fmt.Println(Calculate2(adapters))
	fmt.Println("========")

}
