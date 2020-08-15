package pkg

import (
	"errors"
)

//RomanToArabicalNumeral => Struct Containing Roman Numeral to Arabic Numeral
type RomanToArabicNumeral struct {
	Roman                string
	Arabic               int
	IsAllowedConsecutive bool
}

//MapRomanToArabicNumeral => Map Containing Roman Numeral to Arabic Numeral
var MapRomanToArabicNumeral = make(map[string]RomanToArabicNumeral)

func init() {
	MapRomanToArabicNumeral = map[string]RomanToArabicNumeral{
		"I": {"I", 1, true},
		"V": {"V", 5, false},
		"X": {"X", 10, true},
		"L": {"L", 50, false},
		"C": {"C", 100, true},
		"D": {"D", 500, false},
		"M": {"M", 1000, true},
	}
}

//ConvertRomanToArabic => Convert Roman Numeral to Arabic Numeral
func ConvertRomanToArabic(romanNumeral string) (int, error) {
	arabicNumeral := 0
	lnRomanNumeral := len(romanNumeral)
	consecutiveCounter := 0
	var err error

	//Read each character
	for i := 0; i < lnRomanNumeral; i++ {
		currentChar := string(romanNumeral[i])
		currentCharArabicValue := MapRomanToArabicNumeral[currentChar].Arabic

		if i < lnRomanNumeral-1 {
			nextChar := string(romanNumeral[i+1])
			nextCharArabicValue := MapRomanToArabicNumeral[nextChar].Arabic

			//Count same characters are displayed maximum 3 consecutively
			if MapRomanToArabicNumeral[currentChar].IsAllowedConsecutive == true && currentChar == nextChar {
				consecutiveCounter++
			} else if MapRomanToArabicNumeral[currentChar].IsAllowedConsecutive == false && currentChar == nextChar {
				err = errors.New(currentChar + " cannot appear 3 times in a row")
				return 0, err
			}

			//Check whether current value is smaller than next value
			if currentCharArabicValue < nextCharArabicValue && consecutiveCounter == 0 {
				arabicNumeral += nextCharArabicValue - currentCharArabicValue
				i++
			} else if currentCharArabicValue >= nextCharArabicValue && consecutiveCounter <= 2 {
				arabicNumeral += currentCharArabicValue
			} else {
				err = errors.New("Lower values cannot appear 3 times in a row before higher values")
				return 0, err
			}

			//Reset counter
			if currentCharArabicValue != nextCharArabicValue {
				consecutiveCounter = 0
			}

		} else {
			arabicNumeral += currentCharArabicValue
		}
	}

	return arabicNumeral, nil
}

//ConvertInterGalacticToRoman => Convert Inter Galactic Numeral to Roman Numeral
func ConvertInterGalacticToRoman(firstChar int, lastChar int, splitItem []string, intergalacticToRomanNumeral map[string]string) (string, string, error) {
	stringContainer := ""
	romanNumeral := ""
	var err error

	for i := firstChar; i < lastChar; i++ {
		//Check the string is valid intergalactic numerals
		_, doesStringExistInIntergalacticNumerals := intergalacticToRomanNumeral[splitItem[i]]

		if doesStringExistInIntergalacticNumerals {
			stringContainer += splitItem[i] + " "
			romanNumeral += intergalacticToRomanNumeral[splitItem[i]]
		} else {
			err = errors.New("One of intergalactic numerals is unidentified")
			stringContainer = ""
			romanNumeral = ""
			break
		}
	}

	return stringContainer, romanNumeral, err
}
