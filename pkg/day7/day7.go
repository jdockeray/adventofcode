package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Bag struct {
	count int
	name  string
}

type BagWrapper struct {
	children []Bag
	name     string
}

type BagMachine struct {
	bags map[string]BagWrapper
}

func (machine *BagMachine) addBag(wrapper BagWrapper) {
	if machine.bags == nil {
		machine.bags = map[string]BagWrapper{}
	}
	machine.bags[wrapper.name] = wrapper
}

func slug(str string) string {
	str = strings.Trim(str, " ")
	return strings.ReplaceAll(str, " ", "-")
}

func (machine *BagMachine) Contains(bagWrapper BagWrapper, slug string) bool {

	if len(bagWrapper.children) == 0 {
		return false
	}

	found := false
	for _, child := range bagWrapper.children {
		if child.name == slug {
			return true
		}
		childWrapper := machine.bags[child.name]

		if machine.Contains(childWrapper, slug) {
			fmt.Println(childWrapper)
			found = true
		}
	}
	return found
}

func (machine *BagMachine) Count(slug string) int {
	bagWrapper := machine.bags[slug]
	sum := 0
	for _, child := range bagWrapper.children {
		sum += child.count + (child.count * machine.Count(child.name))
	}

	return sum
}

func (machine *BagMachine) Available(slug string) int {
	count := 0
	for _, bagWrapper := range machine.bags {
		if machine.Contains(bagWrapper, slug) {
			count = count + 1
		}
	}
	return count
}

func ParseLine(line string) BagWrapper {
	line = strings.ReplaceAll(line, "bags", "")
	line = strings.ReplaceAll(line, "bag", "")
	line = strings.Trim(line, "\n")
	line = strings.Trim(line, ".")
	line = strings.Trim(line, " ")

	split := strings.Split(line, "contain")

	name := slug(split[0])

	bags := make([]Bag, 0, 10)

	if strings.Contains(split[1], "no other") {
		return BagWrapper{
			name:     name,
			children: bags,
		}
	}

	innerSplit := strings.Split(split[1], ",")

	for _, s := range innerSplit {
		s = strings.Trim(s, " ")
		c, _ := strconv.Atoi(s[0:1])
		n := slug(s[2:]) // this only works if the number of bags is less then 10
		bags = append(bags, Bag{
			count: c,
			name:  n,
		})
	}

	return BagWrapper{
		name:     name,
		children: bags,
	}
}

func main() {
	resp, err := http.Get("http://127.0.0.1:8081/data/day7.debug.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	bagMachine := BagMachine{}
	for scanner.Scan() {
		str := scanner.Text()
		bagMachine.addBag(ParseLine(str))
	}
	fmt.Println(bagMachine.Available("shiny-gold"))
	fmt.Println(bagMachine.Count("shiny-gold"))
}
