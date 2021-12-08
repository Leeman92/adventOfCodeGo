package day8

import (
	"github.com/l33m4n123/adventOfCodeGo/2021/utils"
	"math"
	"sort"
	"strings"
)

type ByLen []string

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Less(i, j int) bool { return len(a[i]) < len(a[j]) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Solve(lines []string) {
	var result []string
	for _, line := range lines {
		output := strings.Split(line, " | ")
		splittedOutput := strings.Split(output[1], " ")
		for _, s := range splittedOutput {
			// Digit 1
			switch len(s) {
			case 2:
				fallthrough
			case 3:
				fallthrough
			case 4:
				fallthrough
			case 7:
				result = append(result, s)
			}
		}
	}

	utils.PostSolution(8, 1, len(result))

	val := 0
	for _, line := range lines {
		mapping := make(map[int]string)
		split := strings.Split(line, " | ")
		signals := split[0]
		output := split[1]
		splitSignal := strings.Split(signals, " ")
		sort.Sort(ByLen(splitSignal))

		for _, signal := range splitSignal {
			if signal == "" {
				continue
			}
			switch len(signal) {
			case 2:
				mapping[1] = signal
			case 3:
				mapping[7] = signal
			case 4:
				mapping[4] = signal
			case 7:
				mapping[8] = signal
			case 5:
				if signalContained(mapping[1], signal) {
					mapping[3] = signal
				} else if offByOne(mapping[4], signal) {
					mapping[5] = signal
				} else {
					mapping[2] = signal
				}
			case 6:
				if signalContained(mapping[7], signal) && signalContained(mapping[4], signal) {
					mapping[9] = signal
				} else if signalContained(mapping[7], signal) && !signalContained(mapping[4], signal) {
					mapping[0] = signal
				} else {
					mapping[6] = signal
				}
			}
		}

		val += parseOutput(output, mapping)
	}

	utils.PostSolution(8, 2, val)
}

func parseOutput(output string, mapping map[int]string) int {
	splitOutput := strings.Split(output, " ")
	var res []int
	for _, digitSignal := range splitOutput {
		// Sort the signals for better comparison
		splitSignal := strings.Split(digitSignal, "")
		sort.Strings(splitSignal)
		digitSignal = strings.Join(splitSignal, "")
		for digit, val := range mapping {
			// Sort the signals for better comparison
			splitVal := strings.Split(val, "")
			sort.Strings(splitVal)
			val = strings.Join(splitVal, "")
			if val != digitSignal {
				continue
			}
			res = append(res, digit)
		}
	}
	result := 0
	for key, dig := range res {
		result += dig * int(math.Pow(10, float64(len(res)-key-1)))
	}
	return result
}

func signalContained(mapping string, signal string) bool {
	splitMapping := strings.Split(mapping, "")
	for _, pat := range splitMapping {
		if !strings.Contains(signal, pat) {
			return false
		}
	}

	return true
}

func offByOne(mapping string, signal string) bool {
	splitMapping := strings.Split(mapping, "")
	err := 0
	for _, pat := range splitMapping {
		if !strings.Contains(signal, pat) {
			err += 1
		}
	}

	return err == 1
}
