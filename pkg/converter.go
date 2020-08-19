package pkg

import (
	"errors"
	"strconv"
)

//RomanToArabicNumeral => Struct Containing Roman Numeral to Arabic Numeral
type RomanToArabicNumeral struct {
	Roman                string
	Arabic               int
	IsAllowedConsecutive bool
}

//CountMaterialWorthArgs => Struct containing arguments for MultiplyInterGalacticNumeralWithMaterialWorth function
type CountMaterialWorthArgs struct {
	firstChar                   int
	lastChar                    int
	words                       []string
	intergalacticToRomanNumeral map[string]string
	materialName                string
	materialWorth               map[string]float64
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

			//Same characters are displayed maximum 3 consecutively
			if MapRomanToArabicNumeral[currentChar].IsAllowedConsecutive == true && currentChar == nextChar {
				consecutiveCounter++
			} else if MapRomanToArabicNumeral[currentChar].IsAllowedConsecutive == false && currentChar == nextChar {
				err = errors.New(currentChar + " cannot appear more than once in a row")
				return 0, err
			}

			//Check whether current value is smaller than next value
			if currentCharArabicValue < nextCharArabicValue && consecutiveCounter == 0 {
				arabicNumeral += nextCharArabicValue - currentCharArabicValue
				i++
			} else if currentCharArabicValue >= nextCharArabicValue && consecutiveCounter <= 2 {
				arabicNumeral += currentCharArabicValue
			} else {
				err = errors.New("Lower values (" + currentChar + ") cannot appear " + strconv.Itoa(consecutiveCounter+1) + " times in a row before higher values (" + nextChar + ")")
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

		if !doesStringExistInIntergalacticNumerals {
			err = errors.New("Requested number is in invalid format")
			stringContainer = ""
			romanNumeral = ""
			break
		}

		stringContainer += splitItem[i] + " "
		romanNumeral += intergalacticToRomanNumeral[splitItem[i]]
	}

	return stringContainer, romanNumeral, err
}

//ConvertInterGalacticToArabic => Convert Inter Galactic Numeral to Arabic Numeral
func ConvertInterGalacticToArabic(firstChar int, lastChar int, splitItem []string, intergalacticToRomanNumeral map[string]string) (string, int, error) {
	var err error

	intergalacticNumeral, romanNumeral, err := ConvertInterGalacticToRoman(firstChar, lastChar, splitItem, intergalacticToRomanNumeral)
	if err != nil {
		err = errors.New("Requested number is in invalid format")
		return "", 0, err
	}

	var arabicNumeral int
	if intergalacticNumeral != "" && romanNumeral != "" {
		arabicNumeral, err = ConvertRomanToArabic(romanNumeral)
		if err != nil {
			err = errors.New("Requested number is in invalid format")
			return "", 0, err
		}
	}

	return intergalacticNumeral, arabicNumeral, nil
}

//MultiplyInterGalacticNumeralWithMaterialWorth => Multiply Inter Galatic Numeral with Material Worth
func MultiplyInterGalacticNumeralWithMaterialWorth(args CountMaterialWorthArgs) (string, float64, error) {
	var err error
	var countNumeralAndCredits float64

	//Convert intergalactic to arabic numeral
	romanNumeral, arabicNumeral, err := ConvertInterGalacticToArabic(args.firstChar, args.lastChar, args.words, args.intergalacticToRomanNumeral)
	if err != nil {
		err = errors.New("Requested number is in invalid format")
		return "", 0, err
	}

	//Get material worth
	if _, doesMaterialExist := args.materialWorth[args.materialName]; doesMaterialExist == false {
		err = errors.New("Material " + args.materialName + " is unidentified")
		return "", 0, err
	}

	var materialWorthValue float64
	materialWorthValue = args.materialWorth[args.materialName]

	if romanNumeral != "" && arabicNumeral != 0 {
		countNumeralAndCredits = materialWorthValue * float64(arabicNumeral)
	}

	return romanNumeral, countNumeralAndCredits, nil
}
