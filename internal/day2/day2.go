package day2

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Policy struct {
	term string
	min  int
	max  int
}

type PasswordPolicy struct {
	policy   Policy
	password string
}

func getMinMax(str string) (int, int) {
	minmax := strings.Split(str, "-")
	min, _ := strconv.Atoi(minmax[0])
	max, _ := strconv.Atoi(minmax[1])
	return min, max
}

func parsePassword(str string) (Policy, string) {
	stores := strings.Fields(str)
	policy := Policy{}
	var password string
	for index, s := range stores {
		switch index {
		case 0:
			min, max := getMinMax(s)
			policy.min = min
			policy.max = max
		case 1:
			policy.term = s[0:1]
		case 2:
			password = s
		}
	}

	return policy, password
}

func FetchData(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %d", url, res.StatusCode)
	}

	return res, nil
}

func Day2() {
	res, err := FetchData(os.Getenv("FILE_SERVER") + "/day2.txt")
	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	passwordPolicies := make([]PasswordPolicy, 0, 10)

	for scanner.Scan() {
		str := scanner.Text()
		policy, password := parsePassword(str)
		passwordPolicies = append(passwordPolicies, PasswordPolicy{policy: policy, password: password})
	}

	validsTask1 := validPasswords(passwordPolicies, validPasswordTask1)
	fmt.Println(len(validsTask1))

	validsTask2 := validPasswords(passwordPolicies, validPasswordTask2)
	fmt.Println(len(validsTask2))
}

func validPasswordTask1(passwordPolicy PasswordPolicy) bool {
	frequency := strings.Count(passwordPolicy.password, passwordPolicy.policy.term)
	return frequency >= passwordPolicy.policy.min && frequency <= passwordPolicy.policy.max
}

func validPasswordTask2(passwordPolicy PasswordPolicy) bool {
	min, max, term, password := passwordPolicy.policy.min, passwordPolicy.policy.max, passwordPolicy.policy.term, passwordPolicy.password

	pos1 := password[min-1:min] == term
	pos2 := password[max-1:max] == term

	return (pos1 && !pos2) || (pos2 && !pos1)

}

func validPasswords(passwordPolicies []PasswordPolicy, validate func(PasswordPolicy) bool) []PasswordPolicy {
	valid := make([]PasswordPolicy, 0, 10)
	for _, passwordPolicy := range passwordPolicies {
		if validate(passwordPolicy) {
			valid = append(valid, passwordPolicy)
		}
	}
	return valid
}
