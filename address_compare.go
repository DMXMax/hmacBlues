package main

import (
	"fmt"
	"regexp"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

type AddressTable []AddressPair

func main() {

	for _, comparison := range ProcessAddressTable(addressTable) {
		fmt.Printf("Address ID: %s\n", comparison.id)
		fmt.Printf("Address from Source 1: \"%s\"\n", comparison.Address1)
		fmt.Printf("Address from Source 1: \"%s\"\n", comparison.Address2)
		fmt.Printf("Indexed Difference: %#v\n", comparison.diffs)
		fmt.Printf("Edit Distance: %d\n", comparison.distance)
		fmt.Println()
	}

}

func ProcessAddressTable(table AddressTable) []AddressComparison {
	res := make([]AddressComparison, 0, len(table))

	for _, a := range table {
		res = append(res, CompareAddress(a))
	}

	return res
}

func CompareAddress(ap AddressPair) AddressComparison {
	return AddressComparison{
		AddressPair{
			ap.id,
			replaceCharacters(ap.Address1),
			replaceCharacters(ap.Address2),
		},
		indexCompare(ap.Address1, ap.Address2),
		levenshtein.DistanceForStrings([]rune(ap.Address1), []rune(ap.Address2), levenshtein.DefaultOptions),
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

type AddressPair struct {
	id       string
	Address1 string
	Address2 string
}

type AddressComparison struct {
	AddressPair //A promoted Structure
	diffs       []int
	distance    int
}

type AdressTest struct {
	AddressPair
}

var addressTable []AddressPair = []AddressPair{
	AddressPair{
		id:       "000001", //The Dursleys 4 Privet Drive, Little Whinging, Surrey, UK
		Address1: "4 Privet Drive",
		Address2: "Four Privet Dr.",
	},
	AddressPair{
		id:       "000002", //Nightmare on Elm Street
		Address1: "1428 Elm Street",
		Address2: "1428 Ellm Street",
	},
	AddressPair{
		id:       "000003", //Doc Brown 1640 Riverside Drive, Hill Valley, California
		Address1: "1640 Riverside Drive",
		Address2: "164 Riverside Drive",
	},
	AddressPair{
		id:       "000004", //Buffy the Vampire Slayer: 1630 Revello Drive, Sunnydale, CA
		Address1: "1630 Revello Drive",
		Address2: "1640 Riverside Drive",
	},
	AddressPair{
		id:       "000005", //Fox Mulder, 2630 Hegal Place, Apt. 42, Alexandria, Virginia, 23242
		Address1: "2630 Hegal Place, Apt. 42",
		Address2: "2630 Hegal Place, Apartment 42",
	},
	AddressPair{
		id:       "000006", //Dana Scully, 3170 W. 53 Rd. #35, Annapolis, Maryland
		Address1: "3170 W. 53 Rd. #35",
		Address2: "3170 West 53 Road #35",
	},
	AddressPair{
		id:       "000007", //Spongebob Squarepants 124 Conch Street, Bikini Bottom, Pacific Ocean
		Address1: "124 Conch Street",
		Address2: "124 Couch Street",
	},
	AddressPair{
		id:       "000008", //Sirius Black, 12 Grimmauld Place, London, UK
		Address1: "12 Grimmauld Place",
		Address2: "640 Riverside Drive",
	},
		AddressPair{
		id:       "000009", //Sherlock Holmes, 221B Baker Street, London, UK
		Address1: "221B Baker Street",
		Address2: "221 Baker Street",
	},
		AddressPair{
		id:       "000010", //The Munsters, 1313 Mockingbird Lane, Mockingbird Heights, USA
		Address1: "1313 Mockingbird Lane",
		Address2: "1313 Mockingbird Ln",
	},
}
