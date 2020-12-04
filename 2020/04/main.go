package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var requiredFields = []string{
	"byr", // Birth Year
	"iyr", // Issue Year
	"eyr", // Expiration Year
	"hgt", // Height
	"hcl", // Hair Color
	"ecl", // Eye Color
	"pid", // Passport ID
	"cid", // Country ID
}

type passport struct {
	Byr string // Birth Year
	Iyr string // Issue Year
	Eyr string // Expiration Year
	Hgt string // Height
	Hcl string // Hair Color
	Ecl string // Eye Color
	Pid string // Passport ID
	Cid string // Country ID
}

func main() {
	status := make(chan bool)

	//////////////
	//  TESTING //
	//////////////
	testLines, testErr := readLines("input.test.txt")

	if testErr != nil {
		panic(testErr)
	}

	testPassports := parsePassports(testLines)
	go partOne(status, testPassports)
	<-status

	testLines2Valid, testErr2Valid := readLines("input.test.2.valid.txt")

	if testErr2Valid != nil {
		panic(testErr2Valid)
	}

	testPassports = parsePassports(testLines2Valid)
	go partTwo(status, testPassports)
	<-status

	testLines2Invalid, testErr2Invalid := readLines("input.test.2.invalid.txt")

	if testErr2Invalid != nil {
		panic(testErr2Invalid)
	}

	testPassports = parsePassports(testLines2Invalid)
	go partTwo(status, testPassports)
	<-status

	fmt.Printf("\n/////////////////////\n// NO MORE TESTING //\n/////////////////////\n")

	/////////////////////
	// NO MORE TESTING //
	/////////////////////

	lines, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}

	passports := parsePassports(lines)
	go partOne(status, passports)
	<-status
	go partTwo(status, passports)
	<-status
}

func partOne(status chan bool, passports []passport) {
	validCount := validatePassportsPartOne(passports)

	fmt.Printf("\n====The answer for part 1====\n%d passports are valid out of %d\n=============================\n", validCount, len(passports))
	status <- true
}

func partTwo(status chan bool, passports []passport) {
	validCount := validatePassportsPartTwo(passports)

	fmt.Printf("\n====The answer for part 2====\n%d passports are valid out of %d\n=============================\n", validCount, len(passports))
	status <- true
}

func validatePassportsPartOne(passports []passport) int {
	counter := 0
	for _, psPort := range passports {
		if psPort.Byr == "" ||
			psPort.Iyr == "" ||
			psPort.Eyr == "" ||
			psPort.Hgt == "" ||
			psPort.Hcl == "" ||
			psPort.Ecl == "" ||
			psPort.Pid == "" {
			continue
		}

		counter++
	}
	return counter
}

func validatePassportsPartTwo(passports []passport) int {
	counter := 0
	for _, psPort := range passports {
		if psPort.Byr == "" ||
			psPort.Iyr == "" ||
			psPort.Eyr == "" ||
			psPort.Hgt == "" ||
			psPort.Hcl == "" ||
			psPort.Ecl == "" ||
			psPort.Pid == "" {
			continue
		}

		byr, _ := strconv.Atoi(psPort.Byr)
		iyr, _ := strconv.Atoi(psPort.Iyr)
		eyr, _ := strconv.Atoi(psPort.Eyr)

		if byr < 1920 || byr > 2002 {
			continue
		}

		if iyr < 2010 || iyr > 2020 {
			continue
		}

		if eyr < 2020 || eyr > 2030 {
			continue
		}

		if len(psPort.Hcl) != 7 || strings.Index(psPort.Hcl, "#") != 0 {
			continue
		}

		if psPort.Ecl != "amb" &&
			psPort.Ecl != "blu" &&
			psPort.Ecl != "brn" &&
			psPort.Ecl != "gry" &&
			psPort.Ecl != "grn" &&
			psPort.Ecl != "hzl" &&
			psPort.Ecl != "oth" {
			continue
		}

		if len(psPort.Pid) != 9 {
			continue
		}

		if strings.Index(psPort.Hgt, "cm") == -1 && strings.Index(psPort.Hgt, "in") == -1 {
			continue
		}

		if strings.Index(psPort.Hgt, "cm") != -1 {
			heights := strings.Split(psPort.Hgt, "cm")
			height, err := strconv.Atoi(heights[0])
			if err != nil {
				continue
			}

			if height < 150 || height > 193 {
				continue
			}
		} else if strings.Index(psPort.Hgt, "in") != -1 {
			heights := strings.Split(psPort.Hgt, "in")
			height, err := strconv.Atoi(heights[0])
			if err != nil {
				continue
			}

			if height < 59 || height > 76 {
				continue
			}
		}

		counter++
	}
	return counter
}

func parsePassports(lines []string) []passport {
	passports := []passport{}
	pssPort := passport{}
	saved := false
	for _, line := range lines {
		if line == "" {
			passports = append(passports, pssPort)
			pssPort = passport{}
			saved = true
			continue
		}

		saved = false

		fields := strings.Split(line, " ")
		for _, field := range fields {
			fieldValue := strings.Split(field, ":")
			switch fieldValue[0] {
			case "byr":
				{
					pssPort.Byr = fieldValue[1]
					break
				}
			case "iyr":
				{
					pssPort.Iyr = fieldValue[1]
					break
				}
			case "eyr":
				{
					pssPort.Eyr = fieldValue[1]
					break
				}
			case "hgt":
				{
					pssPort.Hgt = fieldValue[1]
					break
				}
			case "hcl":
				{
					pssPort.Hcl = fieldValue[1]
					break
				}
			case "ecl":
				{
					pssPort.Ecl = fieldValue[1]
					break
				}
			case "pid":
				{
					pssPort.Pid = fieldValue[1]
					break
				}
			case "cid":
				{
					pssPort.Cid = fieldValue[1]
					break
				}
			default:
				{
					panic(fmt.Sprintf("%s not defined as a field\n", fieldValue[0]))
				}
			}
		}
	}

	if !saved {
		passports = append(passports, pssPort)
	}

	return passports
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		lines = append(lines, lineStr)
	}

	return lines, scanner.Err()
}
