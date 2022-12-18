package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var delimiter byte

	r := bufio.NewReader(os.Stdin)
	m := map[byte]interface{}{'+': nil, '-': nil, '*': nil, '/': nil}
	s, _, _ := r.ReadLine()
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' {
			s = append(s[:i], s[i+1:]...)
			i--
		} else if _, ok := m[s[i]]; ok == true {
			delimiter = s[i]
		}
	}
	if s[0] >= '0' && s[0] <= '9' {
		doItArabicWay(s, delimiter)
	} else {
		doItRomeWay(s, delimiter)
	}
}

func doItRomeWay(s []byte, delimiter byte) {
	l := strings.Split(string(s), string(delimiter))
	if len(l) != 2 {
		fmt.Println("Incorrect parameters")
		return
	}
	first, err := romeToArabic(l[0])
	if err != nil {
		fmt.Println("Cant parse number")
		return
	}
	second, err := romeToArabic(l[1])
	if err != nil {
		fmt.Println("Cant parse number")
		return
	}
	if first < 1 || first > 10 || second < 1 || second > 10 {
		fmt.Println("Please enter number between 1 and 10")
		return
	}

	result := 0
	switch delimiter {
	case '+':
		result = first + second
	case '-':
		result = first - second
	case '*':
		result = first * second
	case '/':
		result = first / second
	}
	if result < 0 {
		fmt.Println("result of operation with rome numbers cant be negative")
	} else {
		fmt.Println(result)
	}
}

func romeToArabic(s string) (int, error) {
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10}
	pre := 0
	sameNumberInRowCount := 0
	result := 0
	for c, n := range s {
		if v, ok := m[byte(n)]; ok == false {
			err := errors.New("Dont known symbol " + string(n))
			return 0, err
		} else {
			if pre == v {
				result += v
				sameNumberInRowCount++
				if (pre == 5 && sameNumberInRowCount >= 1) || (pre == 1 && sameNumberInRowCount > 2) {
					return 0, errors.New("there is no such number")
				}
				pre = v
			} else if pre < v {
				if sameNumberInRowCount > 0 || pre > 1 {
					return 0, errors.New("there is no such number")
				}
				result = v - result
				pre = v
			} else {
				b, err := romeToArabic(s[c:])
				if err != nil {
					return 0, err
				}
				result += b
				return result, nil
			}
		}
	}
	return result, nil
}

func doItArabicWay(s []byte, delimiter byte) {
	l := strings.Split(string(s), string(delimiter))
	if len(l) != 2 {
		fmt.Println("Incorrect parameters")
		return
	}
	first, err := strconv.Atoi(l[0])
	if err != nil {
		fmt.Println("Cant parse number")
		return
	}
	second, err := strconv.Atoi(l[1])
	if err != nil {
		fmt.Println("Cant parse number")
		return
	}
	if first < 1 || first > 10 || second < 1 || second > 10 {
		fmt.Println("Please enter number between 1 and 10")
		return
	}
	result := 0
	switch delimiter {
	case '+':
		result = first + second
	case '-':
		result = first - second
	case '*':
		result = first * second
	case '/':
		result = first / second
	}
	fmt.Println(result)
}
