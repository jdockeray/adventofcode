package day5

import (
	"bufio"
	"fmt"
	"github.com/jdockeray/adventofcode/internal/day2"
	"os"
	"strconv"
	"strings"
)

func ConvertToBinary(z string, o string) func(str string) string {
	return func(str string) string {
		s := strings.ReplaceAll(str, z, "0")
		s = strings.ReplaceAll(s, o, "1")
		return s
	}
}

var convertRow = ConvertToBinary("F", "B")
var convertCol = ConvertToBinary("L", "R")

func ParseRowCols(rowcols string) (row string, col string) {
	row = rowcols[0:7]
	col = rowcols[7:]
	return row, col
}

func n(s string) int {
	i, _ := strconv.ParseInt(s, 2, len(s)+1)
	return int(i)
}

func GetId(rowCols string) int {
	row, col := ParseRowCols(rowCols)
	r := n(convertRow(row))
	c := n(convertCol(col))
	return (r * 8) + c
}

func Day5() {
	res, _ := day2.FetchData(os.Getenv("FILE_SERVER") + "/day5.txt")
	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)
	max := 0
	min := 0
	seats := map[int]bool{}

	for scanner.Scan() {
		str := scanner.Text()

		id := GetId(str)
		seats[id] = true
		if max == 0 && min == 0 {
			// first run of the loop we set max and min to the first value
			max = id
			min = id
		}

		if id > max {
			max = id
		}
		if id < min {
			min = id
		}
	}

	for i := min; i <= max; i++ {
		seat := seats[i]

		if !seat {
			fmt.Println("========")
			fmt.Println("Your seat is:")
			fmt.Println(i)
			fmt.Println("========")
		}
	}

	fmt.Println("The highest value is:")
	fmt.Println(max)
}
