package main

import (
	"fmt"
	"regexp"
)

type Comparison struct {
	String1 string
	String2 string
	diffs   []int
}

func main() {

	fmt.Println(replaceCharacters("13901 Via Rimin, Novato, CA 94945"))

	fmt.Printf("Comparison: %#v\n", fullCompare("Bob Johnson", "Bob johnson"))
}

func fullCompare(s1 string, s2 string) Comparison {
	return Comparison{
		replaceCharacters(s1),
		replaceCharacters(s2),
		indexCompare(s1, s2),
	}
}

func replaceCharacters(s string) string {
	re := regexp.MustCompile("[a-z]")
	re1 := regexp.MustCompile("[A-Z]")
	re2 := regexp.MustCompile("[0-9]")

	return re2.ReplaceAllString(
		re1.ReplaceAllString(
			re.ReplaceAllString(s, "n"),
			"N"),
		"#")
}

func indexCompare(s1 string, s2 string) []int {
	shorter, longer := findShorter(s1, s2)
	res := make([]int, 0, len(shorter))
	fmt.Println(shorter, longer)

	for i, c := range shorter {
		if c != rune(longer[i]) {
			res = append(res, i)
		}
	}

	return res
}

func findShorter(a string, b string) (string, string) {
	if len(a) <= len(b) {
		return a, b
	} else {
		return b, a
	}
}
