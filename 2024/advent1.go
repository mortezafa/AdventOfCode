package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 1.read from file
	// 2. spilt into left and right list
	// 3. sort each list
	// 4.take abs of their difference
	var leftCol []int
	var rightCol []int
	leftCol, rightCol = buildlists(leftCol, rightCol)

	slices.Sort(leftCol)
	fmt.Printf("left col: %v\n", leftCol)
	slices.Sort(rightCol)
	fmt.Printf("right col: %v\n", rightCol)

	distanceSum := 0

	for i := range leftCol {
		num := absDiffInt(leftCol[i], rightCol[i])
		distanceSum += num
	}
	fmt.Printf("Distance Sum: %d\n", distanceSum) // Final Sum of Distance

	score := calcSimilarity(leftCol, rightCol)
	fmt.Printf("Similarity Score: %d\n", score) // Final score of Calc Similarity

}

func absDiffInt(x, y int) int {
	if y > x {
		return y - x
	} else {
		return x - y
	}
}

func calcSimilarity(leftCol, rightCol []int) int {
	counter := 0
	res := 0

	for _, num := range leftCol {
		leftNum := num
		for _, jnum := range rightCol {
			if leftNum == jnum {
				counter++
			}
		}
		res += counter * leftNum
		counter = 0
	}
	return res
}

func buildlists(leftCol, rightCol []int) ([]int, []int) {

	file, err := os.Open("input.txt")
	check(err)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		lineArr := strings.Split(line, "   ")
		leftNum, _ := strconv.Atoi(lineArr[0])
		rightNum, _ := strconv.Atoi(lineArr[1])

		leftCol = append(leftCol, leftNum)
		rightCol = append(rightCol, rightNum)
	}
	return leftCol, rightCol
}
