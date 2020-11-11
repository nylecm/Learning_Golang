package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type pair struct {
	lowerIndex, upperIndex int
}

func main() {
	s1 := "Trump"
	s2 := "Donald Trump said that Donald Trump is the Greatest President that God has ever created! Everyone agrees with Trump!"

	resp, err, indexes := problem1(s1, s2)
	if err != nil {
		log.Printf("ERROR! %v\n", err)
		os.Exit(1)
	}

	fmt.Println(resp)
	fmt.Println(indexes)
}

func problem1(s1 string, s2 string) (int, error, []pair) {
	if len(s1) >= len(s2) {
		return -1, errors.New("s1 must be longer, or the same length as s2"), nil
	}

	s1Count, s1Cur := 0, 0
	var indexes []pair

	for i := 0; i < len(s2); i++ {
		if s2[i] == s1[s1Cur] && s1Cur == len(s1)-1 {
			indexes = append(indexes, pair{i - len(s1), i})
			s1Count++
			s1Cur = 0
		} else if s2[i] == s1[s1Cur] {
			s1Cur++
		}
	}
	return s1Count, nil, indexes
}