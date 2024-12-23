package main

import (
	"advent1.go/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs/input2.txt")
	utils.Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeCount := 0
	for scanner.Scan() {
		reportText := scanner.Text()
		report := strings.Split(reportText, " ")
		var reportArr []int
		for _, num := range report {
			intNum, err := strconv.Atoi(num)
			utils.Check(err)
			reportArr = append(reportArr, intNum)
		}

		if incOrDec(reportArr) {
			safeCount += 1
		}

	}
	fmt.Println(safeCount)

}

func incOrDec(reportArr []int) bool {

	// The levels are either all increasing or all decreasing.
	// Any two adjacent levels differ by at least one and at most three.
	//increasing := incOrDec(reportArr[0], reportArr[len(reportArr)-1])
	var initalState *bool

	prev := reportArr[0]
	for i, curr := range reportArr {
		if i < 1 {
			continue
		}
		currentState := false

		if curr < prev { // decreasing
			currentState = false
		} else if curr > prev { // increaseing
			currentState = true
		} else {
			return false
		}

		checkedDiff := checkdiff(prev, curr)
		if !checkedDiff {
			return false
		}
		prev = curr

		if initalState == nil {
			initalState = &currentState
		}

		if initalState != nil && *initalState != currentState {
			return false
		}
	}
	return true
}

func checkdiff(a, b int) bool {

	absDiff := utils.AbsDiffInt(a, b)
	if absDiff < 1 || absDiff > 3 {
		return false
	}
	return true
}
