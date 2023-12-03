package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var scanInput string
	scanInput = scanner.Text()
	return scanInput
}

func diffieConst(userInput string) []string {
	words := strings.Fields(userInput)
	var nums []string

	for _, word := range words {
		number, _ := strconv.ParseInt(word, 10, 64)
		if number > 0 {
			nums = append(nums, strconv.Itoa(int(number)))
		}
	}
	return nums
}

func main() {
	scanInput := getUserInput()
	nums := diffieConst(scanInput)
	fmt.Printf("g=%v and p=%v", nums[0], nums[1])
}
