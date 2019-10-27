package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

// TODO: clean up code. Especially creating maps
var timeTable map[string]map[string]string
var guardSleepingAmount map[int]string
var guardSleepingTimes map[int]map[int]int

func main() {
	fmt.Println("=========== START ==========")
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	counter := 0
	for _, line := range lines {
		counter++
		processLine(line, counter)
	}

	orderInput()

	findLongestSleepingGuard()
	findMostFrequentSleepingGuard()
}

func findLongestSleepingGuard() {
	longestSleepingTime := -1
	sleepiestGuardNumber := -1
	commonMinute := -1
	commonMinuteAmount := -1

	for guardNumber := range guardSleepingAmount {
		currSleepingTime, _ := strconv.Atoi(guardSleepingAmount[guardNumber])
		if longestSleepingTime <= currSleepingTime {
			longestSleepingTime = currSleepingTime
			sleepiestGuardNumber = guardNumber
		}
	}

	for minute, amount := range guardSleepingTimes[sleepiestGuardNumber] {
		if commonMinuteAmount < amount {
			commonMinute = minute
			commonMinuteAmount = amount
		}
	}
	fmt.Println("========= SOLUTION FOR PART 1 =========")
	printGuardTime(sleepiestGuardNumber, commonMinute)
}

func findMostFrequentSleepingGuard() {
	commonMinuteAmount := -1
	commonMinute := -1
	mostFrequentGuardNumber := -1

	for guardNumber := range guardSleepingTimes {
		for minute, amount := range guardSleepingTimes[guardNumber] {
			if commonMinuteAmount < amount {
				commonMinute = minute
				commonMinuteAmount = amount
				mostFrequentGuardNumber = guardNumber
			}
		}
	}
	fmt.Println("========= SOLUTION FOR PART 2 =========")
	printGuardTime(mostFrequentGuardNumber, commonMinute)
}

func printGuardTime(guardNumber int, commonMinute int) {
	solution := guardNumber * commonMinute
	fmt.Println("==== Guard #" + strconv.Itoa(guardNumber) + " ====")
	fmt.Printf("Total minutes asleep: %s\n", guardSleepingAmount[guardNumber])
	fmt.Printf("Most asleep at minute: %d\n", commonMinute)
	fmt.Printf("Solution is: %d\n", solution)
}

func orderInput() {
	var lastGuardNumber, timeFellAsleep int

	// get dates in order
	sortedDates := make([]string, 0, len(timeTable))
	for key := range timeTable {
		sortedDates = append(sortedDates, key)
	}
	sort.Strings(sortedDates)

	for _, date := range sortedDates {
		sortedTimestamps := make([]string, 0, len(timeTable[date]))

		for key := range timeTable[date] {
			sortedTimestamps = append(sortedTimestamps, key)
		}
		sort.Strings(sortedTimestamps)

		for _, timeStamp := range sortedTimestamps {
			var action string
			var guardNumber int

			action = timeTable[date][timeStamp]
			action = strings.TrimSpace(action)

			if strings.Contains(action, "Guard") {
				guardNumber = getGuardNumber(action)
				lastGuardNumber = guardNumber
				setupGuard(guardNumber)
			} else if strings.Contains(action, "falls asleep") {
				timeFellAsleep, _ = strconv.Atoi(strings.Replace(timeStamp, "00:", "", -1))
			} else if strings.Contains(action, "wakes up") {
				timeWakeUp, _ := strconv.Atoi(strings.Replace(timeStamp, "00:", "", -1))
				wasAsleep := timeWakeUp - timeFellAsleep

				var currTimeAsleep int
				// Now we can log the data!
				currTimeAsleep, _ = strconv.Atoi(guardSleepingAmount[lastGuardNumber])
				for i := timeFellAsleep; i < timeWakeUp; i++ {
					if _, ok := guardSleepingTimes[lastGuardNumber]; !ok {
						guardSleepingTimes = map[int]map[int]int{
							lastGuardNumber: map[int]int{
								0:  0,
								1:  0,
								2:  0,
								3:  0,
								4:  0,
								5:  0,
								6:  0,
								7:  0,
								8:  0,
								9:  0,
								10: 0,
								11: 0,
								12: 0,
								13: 0,
								14: 0,
								15: 0,
								16: 0,
								17: 0,
								18: 0,
								19: 0,
								20: 0,
								21: 0,
								22: 0,
								23: 0,
								24: 0,
								25: 0,
								26: 0,
								27: 0,
								28: 0,
								29: 0,
								30: 0,
								31: 0,
								32: 0,
								33: 0,
								34: 0,
								35: 0,
								36: 0,
								37: 0,
								38: 0,
								39: 0,
								40: 0,
								41: 0,
								42: 0,
								43: 0,
								44: 0,
								45: 0,
								46: 0,
								47: 0,
								48: 0,
								49: 0,
								50: 0,
								51: 0,
								52: 0,
								53: 0,
								54: 0,
								55: 0,
								56: 0,
								57: 0,
								58: 0,
								59: 0,
							},
						}
					}
					guardSleepingTimes[lastGuardNumber][i] = guardSleepingTimes[lastGuardNumber][i] + 1

				}

				currTimeAsleep = currTimeAsleep + wasAsleep
				guardSleepingAmount[lastGuardNumber] = strconv.Itoa(currTimeAsleep)
			}
		}
	}
}

func setupGuard(guardNumber int) {
	if _, ok := guardSleepingAmount[guardNumber]; !ok {
		if guardSleepingAmount == nil {
			guardSleepingAmount = map[int]string{
				guardNumber: "0",
			}
		} else {
			guardSleepingAmount[guardNumber] = "0"
		}
	}

	if _, ok := guardSleepingTimes[guardNumber]; !ok {
		if guardSleepingTimes == nil {
			guardSleepingTimes = map[int]map[int]int{
				guardNumber: map[int]int{
					0: 0,
				},
			}
		} else {
			guardSleepingTimes[guardNumber] = map[int]int{
				0: 0,
			}
		}
	}
}

func getGuardNumber(action string) int {
	var guardTag, guardNumberString string
	fmt.Sscanf(action, "%s %s", &guardTag, &guardNumberString)
	guardNumber, err := strconv.Atoi(strings.Replace(guardNumberString, "#", "", -1))
	if err != nil {
		panic(err)
	}
	return guardNumber
}

func processLine(line string, counter int) {
	var date, time, action string
	_, err := fmt.Sscanf(line, "%s %s", &date, &time)
	action = strings.Replace(line, date, "", -1)
	action = strings.Replace(action, time, "", -1)
	date = strings.Replace(date, "[", "", -1)
	time = strings.Replace(time, "]", "", -1)

	if _, ok := timeTable[date]; !ok {
		if timeTable == nil {
			timeTable = map[string]map[string]string{
				date: map[string]string{
					time: action,
				},
			}
		} else {
			timeTable[date] = map[string]string{
				time: action,
			}
		}
	} else {
		timeTable[date][time] = action
	}

	if err != nil {
		log.Fatal(err)
	}

}
