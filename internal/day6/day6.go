package day6

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func SetSum(set map[string]int, size int) int {
	sum := 0
	for _, val := range set {
		if val >= size {
			sum = sum + 1
		}
	}
	return sum
}

func Day6() {
	resp, err := http.Get(os.Getenv("FILE_SERVER") + "/day6.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	sum := 0
	size := 0
	strSet := map[string]int{}

	for scanner.Scan() {
		str := scanner.Text()

		for _, r := range str {
			val, exists := strSet[string(r)]
			if exists {
				strSet[string(r)] = val + 1
			} else {
				strSet[string(r)] = 1
			}
		}

		if len(str) == 0 {

			sum = sum + SetSum(strSet, size)

			// reset
			strSet = map[string]int{}
			size = 0
		} else {

			size = size + 1
		}
	}

	// this is to cover the end of file, i feel like there must be a better way of doing this
	sum = sum + SetSum(strSet, size)

	fmt.Println(sum)

}
