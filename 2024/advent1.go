package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
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

	file, err := os.Open("input.txt")
	check(err)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var locationIDs []int
	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)
		var sb strings.Builder

		for i := 0; i < len(line); i++ {

			if unicode.IsNumber(rune(line[i])) {
				sb.WriteByte(line[i])
				check(err)
				// some if statement to append the byte if its a the last index
				if i == len(line)-1 {
					num1, _ := strconv.Atoi(sb.String())
					locationIDs = append(locationIDs, num1)
					sb.Reset()
				}
			} else {
				if sb.Len() == 0 {
					continue
				}

				num1, _ := strconv.Atoi(sb.String())
				locationIDs = append(locationIDs, num1)
				sb.Reset()

			}
		}
	}
	fmt.Print(locationIDs)

	var leftCol []int
	var rightCol []int

	for i, num := range locationIDs {
		if i%2 == 0 {
			leftCol = append(leftCol, num)
		} else {
			rightCol = append(rightCol, num)
		}
	}
	slices.Sort(leftCol)
	fmt.Printf("left col: %v\n", leftCol)
	slices.Sort(rightCol)
	fmt.Printf("right col: %v\n", rightCol)

	//distanceSum := 0

	//for i := range leftCol {
	//	num := absDiffInt(leftCol[i], rightCol[i])
	//	distanceSum += num
	//}
	//
	//fmt.Print(distanceSum)

	score := calcSimilarity(leftCol, rightCol)
	fmt.Print(score)

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
