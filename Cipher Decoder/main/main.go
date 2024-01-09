package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getA() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	parts := strings.Fields(input)
	lastPart := parts[len(parts)-1]

	scanA, err := strconv.Atoi(lastPart)
	if err != nil {
		fmt.Println("Error converting input to integer: ", err)
	}
	return scanA
}

func diffieConst(userInput string) ([]string, []string, []string, []int, []int) {
	words := strings.Fields(userInput)
	var nums, p, g []string
	var intP, intG []int

	for _, word := range words {
		number, _ := strconv.ParseInt(word, 10, 64)
		if number > 0 {
			nums = append(nums, strconv.Itoa(int(number)))
		}
	}
	for i, value := range nums[:2] {
		if i == 0 {
			p = append(p, value)
		} else if i == 1 {
			g = append(g, value)
		}
	}
	for _, val := range p {
		intValP, err := strconv.Atoi(val)
		if err != nil {
			fmt.Printf("Error converting string to integer: %v/n", err)
		}
		intP = append(intP, intValP)
	}
	for _, val := range g {
		intValG, err := strconv.Atoi(val)
		if err != nil {
			fmt.Printf("Error converting string to integer: %v/n", err)
		}
		intG = append(intG, intValG)
	}
	return nums, p, g, intP, intG
}

func calcB(intG []int, intP []int, b int) int {
	g := intG[0]
	p := intP[0]
	result := 1
	base := g % p

	for b > 0 {
		if b%2 == 1 {
			result = (result * base) % p
		}
		base = (base * base) % p
		b /= 2
	}
	return result
}

func calcSharedSecret(A int, b int, intP []int) int {
	p := intP[0]
	result := 1
	A = A % p

	for b > 0 {
		if b%2 == 1 {
			result = (result * A) % p
		}
		A = (A * A) % p
		b /= 2
	}
	return result
}

func main() {
	var b int
	b = 15
	scanInput := getUserInput()
	fmt.Println("OK")
	AResult := getA()
	_, _, _, intP, intG := diffieConst(scanInput)
	BResult := calcB(intG, intP, b)
	SResult := calcSharedSecret(AResult, b, intP)
	fmt.Println("B is ", BResult)
	fmt.Println("A is ", AResult)
	fmt.Println("S is ", SResult)
}
