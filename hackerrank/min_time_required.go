package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/minimum-time-required/problem?h_l=interview&playlist_slugs%5B%5D%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D%5B%5D=search

func minTime(machines []int64, goal int64) int64 {
	min, max := machines[0], machines[0]
	for _, el := range machines {
		if el > max {
			max = el
		}

		if el < min {
			min = el
		}
	}

	var lenM int64 = int64(len(machines))
	fastest := int64(float64(goal)/ ((float64(lenM) / float64(min))))
	slowest := int64(float64(goal)/ ((float64(lenM)/ float64(max))))

	for fastest < slowest {
		numDays := (fastest + slowest) / 2

		var total int64 = 0

		for _, el := range machines {
			total += numDays / el
		}

		if total >= goal {
			slowest = numDays
		} else {
			fastest = numDays + 1
		}

	}

	return fastest
}

func MinTimeApp() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nGoal := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nGoal[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	goal, err := strconv.ParseInt(nGoal[1], 10, 64)
	checkError(err)

	machinesTemp := strings.Split(readLine(reader), " ")

	var machines []int64

	for i := 0; i < int(n); i++ {
		machinesItem, err := strconv.ParseInt(machinesTemp[i], 10, 64)
		checkError(err)
		machines = append(machines, machinesItem)
	}

	ans := minTime(machines, goal)

	fmt.Fprintf(writer, "%d\n", ans)

	writer.Flush()
}
