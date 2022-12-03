package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	list := getList("input.txt")
	passports := getPassports(list)
	valids := []passport{}
	for _, pp := range passports {
		valid := validatePassport(pp)
		if valid {
			valids = append(valids, pp)
		}
	}
	fmt.Println("count: ", len(valids))

	count := 0
	for _, pp := range valids {
		valid := validatePassport2(pp)
		if valid {
			count++
		}
	}
	fmt.Println("count 2: ", count)
}

func getList(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func getPassports(list []string) []passport {
	passports := []passport{}

	pp := passport{}
	for _, line := range list {
		if line == "" {
			passports = append(passports, pp)
			pp = passport{}
			continue
		}
		fields := strings.Split(line, " ")
		for _, field := range fields {
			parts := strings.Split(field, ":")
			key := strings.Trim(parts[0], " ")
			value := strings.Trim(parts[1], " ")
			switch key {
			case "byr":
				pp.byr = value
			case "iyr":
				pp.iyr = value
			case "eyr":
				pp.eyr = value
			case "hgt":
				pp.hgt = value
			case "hcl":
				pp.hcl = value
			case "ecl":
				pp.ecl = value
			case "pid":
				pp.pid = value
			case "cid":
				pp.cid = value
			}
		}
	}
	return passports
}

func validatePassport(pp passport) bool {
	if pp.byr == "" ||
		pp.iyr == "" ||
		pp.eyr == "" ||
		pp.hgt == "" ||
		pp.hcl == "" ||
		pp.ecl == "" ||
		pp.pid == "" {
		return false
	}
	return true

}

func validatePassport2(pp passport) bool {
	if !validateByr(pp) {
		return false
	}
	if !validateIyr(pp) {
		return false
	}
	if !validateEyr(pp) {
		return false
	}
	if !validateHgt(pp) {
		return false
	}
	if !validateHcl(pp) {
		return false
	}
	if !validateEcl(pp) {
		return false
	}
	if !validatePid(pp) {
		return false
	}
	return true
}

func validateByr(pp passport) bool {
	if len(pp.byr) != 4 {
		return false
	}
	if pp.byr < "1920" || pp.byr > "2002" {
		return false
	}
	return true
}

func validateIyr(pp passport) bool {
	if len(pp.iyr) != 4 {
		return false
	}
	if pp.iyr < "2010" || pp.iyr > "2020" {
		return false
	}
	return true
}

func validateEyr(pp passport) bool {
	if len(pp.eyr) != 4 {
		return false
	}
	if pp.eyr < "2020" || pp.eyr > "2030" {
		return false
	}
	return true
}

func validateHgt(pp passport) bool {
	var num int
	var unit string
	fmt.Sscanf(pp.hgt, "%d%s", &num, &unit)
	if unit == "cm" {
		if num < 150 || num > 193 {
			return false
		}
	} else if unit == "in" {
		if num < 59 || num > 76 {
			return false
		}
	} else {
		return false
	}

	return true
}

func validateHcl(pp passport) bool {
	var col string
	fmt.Sscanf(pp.hcl, "#%s", &col)
	if len(col) != 6 {
		return false
	}
	for _, c := range col {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return false
		}
	}
	return true
}

func validateEcl(pp passport) bool {
	col := pp.ecl

	if col == "amb" {
		return true
	}
	if col == "blu" {
		return true
	}
	if col == "brn" {
		return true
	}
	if col == "gry" {
		return true
	}
	if col == "grn" {
		return true
	}
	if col == "hzl" {
		return true
	}
	if col == "oth" {
		return true
	}
	return false
}

func validatePid(pp passport) bool {
	if len(pp.pid) != 9 {
		return false
	}
	for _, c := range pp.pid {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
