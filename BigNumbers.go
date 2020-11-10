package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/*
 1   222   333     4   55555  666  77777  888   999    000
11  2   2 3   3   4 4  5     6   6     7 8   8 9   9  0   0
 1     2      3  4  4  5     6        7  8   8 9   9 0     0
 1    2     33  4   4   555  6666    7    888   999  0     0
 1   2        3 444444     5 6   6  7    8   8     9 0     0
 1  2     3   3     4  5   5 6   6 7     8   8     9  0   0
111 22222  333      4   555   666  7      888      9   000
*/

var bigDigits = [][]string {
	{"  000  ", " 0   0 ", "0     0", "0     0", "0     0", " 0   0 ", "  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  4 4 ", " 4  4 ", "4   4 ", "444444", "    4 ", "    4 "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6   6", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 999 ", "9   9", "9   9", " 999 ", "    9", "    9", "    9"},
}

func main()  {
	if len(os.Args) == 1 { //If no arguments
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1) //code 1 is failure, 0 is success.
	}

	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <=9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("Invalid whole number.")
			}
		}
		fmt.Println(line)
	}
}

