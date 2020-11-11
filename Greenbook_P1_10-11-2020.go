package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type indexPair struct {
	lowerIndex, upperIndex int
}

func main() {
	s1 := "Trump"
	s2 := "Donald Trump said that Donald Trump is the Greatest President that God has ever created! Everyone agrees with Trump!"
	runProblem1(s1, s2)

	fmt.Println()

	s1 = "abCd"
	s2 = "ffffabCdCOVID"
	runProblem1(s1, s2)
}

func runProblem1(s1 string, s2 string) {
	resp, err, indexes := problem1(s1, s2)
	if err != nil {
		log.Printf("ERROR! %v\n", err)
		os.Exit(1)
	}

	fmt.Println(resp)
	fmt.Println(indexes)
}

func problem1(s1 string, s2 string) (int, error, []indexPair) {
	if len(s1) >= len(s2) {
		return -1, errors.New("s1 must be longer, or the same length as s2"), nil
	}

	s1Count, s1Cur := 0, 0
	var indexes []indexPair

	for i := 0; i < len(s2); i++ {
		if s2[i] == s1[s1Cur] && s1Cur == len(s1)-1 {
			indexes = append(indexes, indexPair{i - (len(s1) - 1), i})
			s1Count++
			s1Cur = 0
		} else if s2[i] == s1[s1Cur] {
			s1Cur++
		} else if i == len(s2)-len(s1) { // No chance of s1 being in s2 as there are not enough chars in s2 to fit s1.
			break
		}
	}
	return s1Count, nil, indexes
}
