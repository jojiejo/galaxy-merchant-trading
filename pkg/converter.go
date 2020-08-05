package converter

func convertRomanToArabic(romanNumeral string, mapRomanToArabicNumeral map[string]int) int {
	arabicNumeral := 0
	lnRomanNumeral := len(romanNumeral)

	//Read each character
	for i := 0; i < lnRomanNumeral; i++ {
		currentChar := string(romanNumeral[i])
		currentCharArabicValue := mapRomanToArabicNumeral[currentChar]

		if i < lnRomanNumeral-1 {
			nextChar := string(romanNumeral[i+1])
			nextCharArabicValue := mapRomanToArabicNumeral[nextChar]

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
