package converter

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
