package main

import (
	"bufio"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jdockeray/adventofcode/pkg/day2"
	"strconv"
	"strings"
)

type Passport struct {
	Byr int    `validate:"gte=1920,lte=2002"`
	Iyr int    `validate:"gte=2010,lte=2020"`
	Eyr int    `validate:"gte=2020,lte=2030"`
	Hgt string `validate:"hgt"`
	Hcl string `validate:"hexcolor"`
	Ecl string `validate:"ecl"`
	Pid string `validate:"len=9,numeric"`
}

func (receiver Passport) WithByr(byr int) Passport {
	receiver.Byr = byr
	return receiver
}
func (receiver Passport) WithIyr(iyr int) Passport {
	receiver.Iyr = iyr
	return receiver
}
func (receiver Passport) WithEyr(eyr int) Passport {
	receiver.Eyr = eyr
	return receiver
}
func (receiver Passport) WithHgt(hgt string) Passport {
	receiver.Hgt = hgt
	return receiver
}
func (receiver Passport) WithHcl(hcl string) Passport {
	receiver.Hcl = hcl
	return receiver
}
func (receiver Passport) WithEcl(ecl string) Passport {
	receiver.Ecl = ecl
	return receiver
}
func (receiver Passport) WithPid(pid string) Passport {
	receiver.Pid = pid
	return receiver
}

type PassportManager struct {
	passports []Passport
}

func (receiver *PassportManager) add(passport Passport) {
	if receiver.passports == nil {
		receiver.passports = []Passport{}
	}
	receiver.passports = append(receiver.passports, passport)
}

func printErrors(err validator.ValidationErrors) {

	for _, err := range err {

		fmt.Println(err.Namespace())
		fmt.Println(err.Field())
		fmt.Println(err.StructNamespace())
		fmt.Println(err.StructField())
		fmt.Println(err.Tag())
		fmt.Println(err.ActualTag())
		fmt.Println(err.Kind())
		fmt.Println(err.Type())
		fmt.Println(err.Value())
		fmt.Println(err.Param())
		fmt.Println()
	}
}

var validate *validator.Validate

func (receiver *PassportManager) validate() int {
	valid := 0
	for _, p := range receiver.passports {

		err := validate.Struct(p)

		if err != nil {
			// printErrors(err.(validator.ValidationErrors))
		} else {
			valid = valid + 1
		}
	}

	return valid
}

func ValidateEyeColor(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, color := range colors {
		if val == color {
			return true
		}
	}
	return false
}

func ValidateHeight(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	if len(str) <= 2 {
		return false
	}
	unit := str[0 : len(str)-2]
	metric := str[len(str)-2:]

	numeric, err := strconv.Atoi(unit)
	if err != nil {
		return false
	}

	if metric == "cm" {
		return numeric >= 150 && numeric <= 193
	}
	if metric == "in" {
		return numeric >= 59 && numeric <= 76
	}

	return false

}

func ValidateSetUp() {
	validate = validator.New()
	err := validate.RegisterValidation("ecl", ValidateEyeColor)
	err = validate.RegisterValidation("hgt", ValidateHeight)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ValidateSetUp()

	res, _ := day2.FetchData("http://127.0.0.1:8080/data/day4.txt")
	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)
	manager := PassportManager{
		passports: nil,
	}

	passport := Passport{}
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) == 0 {
			manager.add(passport)
			passport = Passport{}
		} else {
			for _, field := range strings.Fields(str) {
				split := strings.Split(field, ":")
				key, val := split[0], split[1]
				//passport[key] = val
				switch key {
				case "byr":
					passport.Byr, _ = strconv.Atoi(val)
				case "iyr":
					passport.Iyr, _ = strconv.Atoi(val)
				case "eyr":
					passport.Eyr, _ = strconv.Atoi(val)
				case "hgt":
					passport.Hgt = val
				case "hcl":
					passport.Hcl = val
				case "ecl":
					passport.Ecl = val
				case "pid":
					passport.Pid = val
				}
			}
		}
	}
	manager.add(passport)

	fmt.Println(manager.validate())
}
