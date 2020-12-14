package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/l33m4n123/adventOfCodeGo/utils"
)

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

// Solve runs the puzzle
func Solve(lines []string) {
	passports := parsePassports(lines)
	partOne(passports)
	partTwo(passports)
}

func partOne(passports []passport) {
	validCount := validatePassportsPartOne(passports)

	utils.PostSolution(4, 1, validCount)
}

func partTwo(passports []passport) {
	validCount := validatePassportsPartTwo(passports)

	utils.PostSolution(4, 2, validCount)
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
