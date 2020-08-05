package pkg

import "fmt"

//MapRomanToArabicNumeral => Map Containing Roman Numeral to Arabic Numeral
var MapRomanToArabicNumeral = make(map[string]int)

func init() {
	MapRomanToArabicNumeral = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
}

//ConvertRomanToArabic => Convert Roman Numeral to Arabic Numeral
func ConvertRomanToArabic(romanNumeral string) int {
	arabicNumeral := 0
	lnRomanNumeral := len(romanNumeral)

	//Read each character
	for i := 0; i < lnRomanNumeral; i++ {
		currentChar := string(romanNumeral[i])
		currentCharArabicValue := MapRomanToArabicNumeral[currentChar]

		if i < lnRomanNumeral-1 {
			nextChar := string(romanNumeral[i+1])
			nextCharArabicValue := MapRomanToArabicNumeral[nextChar]

			//Check whether current value is smaller than next value
			if currentCharArabicValue < nextCharArabicValue {
				arabicNumeral += nextCharArabicValue - currentCharArabicValue
				i++
			} else {
				arabicNumeral += currentCharArabicValue
			}

		} else {
			arabicNumeral += currentCharArabicValue
		}
	}

	return arabicNumeral
}

//ConvertInterGalacticToRoman => Convert Inter Galactic Numeral to Roman Numeral
func ConvertInterGalacticToRoman(firstChar int, lastChar int, splitItem []string, intergalacticToRomanNumeral map[string]string) (string, string) {
	stringContainer := ""
	romanNumeral := ""

	for i := firstChar; i < lastChar; i++ {
		//Check the string is valid intergalactic numerals
		_, doesStringExistInIntergalacticNumerals := intergalacticToRomanNumeral[splitItem[i]]

		if doesStringExistInIntergalacticNumerals {
			stringContainer += splitItem[i] + " "
			romanNumeral += intergalacticToRomanNumeral[splitItem[i]]
		} else {
			fmt.Println("One of intergalactic numerals is unidentified.")
			stringContainer = ""
			romanNumeral = ""
			break
		}
	}

	return stringContainer, romanNumeral
}
