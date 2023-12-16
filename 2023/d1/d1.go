package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func extract1(text string) int {
	var num int

	// forward
	for _, c := range text {
		if digit, err := strconv.Atoi(string(c)); err == nil {
			num = digit * 10
			break
		}
	}

	// backward
	for i := len(text) - 1; i >= 0; i-- {
		if digit, err := strconv.Atoi(string(text[i])); err == nil {
			num += digit
			break
		}
	}

	return num
}

func findFirstDigit(text string) (int, int) {
	for i, c := range text {
		if d, err := strconv.Atoi(string(c)); err == nil {
			return i, d
		}
	}
	return -1, -1
}

func findLastDigit(text string) (int, int) {
	for i := len(text) - 1; i >= 0; i-- {
		if d, err := strconv.Atoi(string(text[i])); err == nil {
			return i, d
		}
	}
	return -1, -1
}

func findFirstWord(text string) (int, int) {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i := range text {
		for j, word := range words {
			if strings.HasPrefix(text[i:len(text)-1], word) {
				return i, j + 1
			}
		}
	}
	return -1, -1
}

func findLastWord(text string) (int, int) {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i := len(text)-1; i >= 0; i-- {
		for j, word := range words {
			if strings.HasSuffix(text[0:i], word) {
				return i - len(word), j + 1
			}
		}
	}
	return -1, -1
}

func extract2(text string) int {
	var num int

	fdi, fd := findFirstDigit(text)
	ldi, ld := findLastDigit(text)
	fwi, fw := findFirstWord(text)
	lwi, lw := findLastWord(text)

	if fdi == -1 {
		num = fw * 10
	} else if fwi == -1 {
		num = fd * 10
	} else if fdi < fwi {
		num = fd * 10
	} else {
		num = fw * 10
	}

	if ldi == -1 {
		num += lw 
	} else if lwi == -1 {
		num += ld 
	} else if ldi > lwi {
		num += ld 
	} else {
		num += lw 
	}

	return num
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sum := 0
	for {
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		sum += extract2(text)
	}
	fmt.Printf("Sum is %d\n", sum)
}
