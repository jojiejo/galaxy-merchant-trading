package pkg

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

//ProcessIntergalacticStatement => Process given string to be readable by intergalactic merchant
func ProcessIntergalacticStatement(splitStringifiedContent []string) ([]string, error) {
	var intergalacticToRomanNumeral = make(map[string]string)
	var materialWorth = make(map[string]float64)
	var returnedString []string
	var err error
	var phrase = make(map[string]bool)

	for _, sentenceValue := range splitStringifiedContent {
		//Trim space the sentence & split it into words
		trimSentence := strings.TrimSpace(string(sentenceValue))
		splitItem := strings.Split(trimSentence, " ")
		lnSplitItem := len(splitItem)

		// xxx is xxx case
		phrase = map[string]bool{
			"is": true,
		}
		if len(splitItem) == 3 && CheckPhraseInSentence(trimSentence, phrase) {
			//Check the third string is valid roman numerals
			if _, doesStringExistInRomanNumerals := MapRomanToArabicNumeral[splitItem[2]]; doesStringExistInRomanNumerals == false {
				err = errors.New(splitItem[2] + " could not be recognized as roman numeral")
				return nil, err
			}

			intergalacticToRomanNumeral[splitItem[0]] = splitItem[2]
		}

		// xxx is xxx Credits case
		phrase = map[string]bool{
			"is":      true,
			"Credits": true,
			"?":       false,
		}
		if CheckPhraseInSentence(trimSentence, phrase) {
			//Check whether the sentence contains is and Credits
			if len(intergalacticToRomanNumeral) == 0 {
				err = errors.New("Intergalactic numeral declaration must be done at least once")
				return nil, err
			}

			// Find is & credit index
			isIndex, _ := CheckWordIndexInSentence(splitItem, "is")
			creditsIndex, _ := CheckWordIndexInSentence(splitItem, "Credits")

			if isIndex != 0 && creditsIndex != 0 {
				//Form a roman numeral
				romanNumeral := ""
				for i := 0; i < isIndex-1; i++ {
					//Check the string is valid intergalactic numerals
					if _, doesStringExistInIntergalacticNumerals := intergalacticToRomanNumeral[splitItem[i]]; doesStringExistInIntergalacticNumerals == false {
						returnedString = append(returnedString, "Requested number is in invalid format")
						continue
					}

					romanNumeral += intergalacticToRomanNumeral[splitItem[i]]
				}

				//Form the name of material
				materialName := splitItem[isIndex-1]

				//Form the value of credits
				creditsValue, err := strconv.ParseFloat(splitItem[creditsIndex-1], 64)
				if err != nil {
					return nil, err
				}

				//Calculation of material worth
				arabicNumeral, err := ConvertRomanToArabic(romanNumeral)
				if err != nil {
					returnedString = append(returnedString, "Requested number is in invalid format")
					continue
				}

				materialWorthValue := creditsValue / float64(arabicNumeral)
				materialWorth[materialName] = materialWorthValue
			}
		}

		// how much is xxx ? case
		phrase = map[string]bool{
			"how much": true,
		}
		if CheckPhraseInSentence(trimSentence, phrase) {
			//Check whether the sentence contains convertion question
			if splitItem[0] == "how" && splitItem[1] == "much" && splitItem[2] == "is" && splitItem[lnSplitItem-1] == "?" {
				romanNumeral, arabicNumeral, err := ConvertInterGalacticToArabic(3, lnSplitItem-1, splitItem, intergalacticToRomanNumeral)
				if err != nil {
					returnedString = append(returnedString, err.Error())
					continue
				}

				returnedString = append(returnedString, strings.TrimSpace(romanNumeral)+" is "+strconv.Itoa(arabicNumeral))

			} else {
				returnedString = append(returnedString, "I have no idea what you are talking about")
			}
		}

		// how many Credits is xxx ? case
		phrase = map[string]bool{
			"how many": true,
		}
		if CheckPhraseInSentence(trimSentence, phrase) {
			//Check whether the sentence contains material worth question
			if splitItem[0] == "how" && splitItem[1] == "many" && splitItem[2] == "Credits" && splitItem[3] == "is" && splitItem[lnSplitItem-1] == "?" {
				materialName := splitItem[lnSplitItem-2]
				args := CountMaterialWorthArgs{
					firstChar:                   4,
					lastChar:                    lnSplitItem - 2,
					words:                       splitItem,
					intergalacticToRomanNumeral: intergalacticToRomanNumeral,
					materialName:                materialName,
					materialWorth:               materialWorth,
				}

				intergalacticNumeral, materialWorthValue, err := MultiplyInterGalacticNumeralWithMaterialWorth(args)
				if err != nil {
					returnedString = append(returnedString, err.Error())
					continue
				}

				convertedMaterialWorthValue := strconv.FormatFloat(materialWorthValue, 'f', -1, 64)
				returnedString = append(returnedString, strings.TrimSpace(intergalacticNumeral)+" "+materialName+" is "+convertedMaterialWorthValue+" Credits")
			}
		}

		//Does xxx xxx has more Credits than xxx xxx ? case
		phrase = map[string]bool{
			"Does":                  true,
			"has more Credits than": true,
			"?":                     true,
		}
		if CheckPhraseInSentence(trimSentence, phrase) {
			if splitItem[0] == "Does" && splitItem[lnSplitItem-1] == "?" {
				hasIndex, _ := CheckWordIndexInSentence(splitItem, "has")
				thanIndex, _ := CheckWordIndexInSentence(splitItem, "than")

				if hasIndex != 0 && thanIndex != 0 {
					//First numeral
					firstMaterialName := splitItem[hasIndex-1]
					firstArgs := CountMaterialWorthArgs{
						firstChar:                   1,
						lastChar:                    hasIndex - 1,
						words:                       splitItem,
						intergalacticToRomanNumeral: intergalacticToRomanNumeral,
						materialName:                firstMaterialName,
						materialWorth:               materialWorth,
					}

					firstIntergalacticNumeral, firstMaterialWorthValue, err := MultiplyInterGalacticNumeralWithMaterialWorth(firstArgs)
					if err != nil {
						returnedString = append(returnedString, err.Error())
						continue
					}

					//Second numeral
					secondMaterialName := splitItem[lnSplitItem-2]
					secondArgs := CountMaterialWorthArgs{
						firstChar:                   thanIndex + 1,
						lastChar:                    lnSplitItem - 2,
						words:                       splitItem,
						intergalacticToRomanNumeral: intergalacticToRomanNumeral,
						materialName:                secondMaterialName,
						materialWorth:               materialWorth,
					}

					secondIntergalacticNumeral, secondMaterialWorthValue, err := MultiplyInterGalacticNumeralWithMaterialWorth(secondArgs)
					if err != nil {
						returnedString = append(returnedString, err.Error())
						continue
					}

					//Compare both numerals
					if firstMaterialWorthValue != 0 && secondMaterialWorthValue != 0 {
						var comparativeWord, prepWord string

						if firstMaterialWorthValue > secondMaterialWorthValue {
							comparativeWord = "more"
							prepWord = "than"
						} else if firstMaterialWorthValue < secondMaterialWorthValue {
							comparativeWord = "less"
							prepWord = "than"
						} else {
							comparativeWord = "same"
							prepWord = "as"
						}

						returnedString = append(returnedString, strings.TrimSpace(firstIntergalacticNumeral)+" "+
							firstMaterialName+" has "+comparativeWord+" Credits "+prepWord+" "+strings.TrimSpace(secondIntergalacticNumeral)+" "+secondMaterialName)
					}
				}
			}
		}

		//Is xxx larger than xxx ? case
		phrase = map[string]bool{
			"Is":          true,
			"larger than": true,
			"?":           true,
		}
		if CheckPhraseInSentence(trimSentence, phrase) {
			if splitItem[0] == "Is" && splitItem[lnSplitItem-1] == "?" {
				// Find larger than index
				largerIndex, _ := CheckWordIndexInSentence(splitItem, "larger")
				thanIndex, _ := CheckWordIndexInSentence(splitItem, "than")

				if largerIndex != 0 && thanIndex != 0 {
					//First numeral
					firstRomanNumeral, firstArabicNumeral, err := ConvertInterGalacticToArabic(1, largerIndex, splitItem, intergalacticToRomanNumeral)
					if err != nil {
						returnedString = append(returnedString, err.Error())
						continue
					}

					//Second numeral
					secondRomanNumeral, secondArabicNumeral, err := ConvertInterGalacticToArabic(thanIndex+1, lnSplitItem-1, splitItem, intergalacticToRomanNumeral)
					if err != nil {
						returnedString = append(returnedString, err.Error())
						continue
					}

					//Compare both numerals
					if firstArabicNumeral != 0 && secondArabicNumeral != 0 {
						var comparativeWord string
						if firstArabicNumeral > secondArabicNumeral {
							comparativeWord = "larger than"
						} else if firstArabicNumeral < secondArabicNumeral {
							comparativeWord = "smaller than"
						} else {
							comparativeWord = "equal to"
						}

						returnedString = append(returnedString, strings.TrimSpace(firstRomanNumeral)+" is "+comparativeWord+" "+strings.TrimSpace(secondRomanNumeral))
					}
				}
			}
		}
	}

	return returnedString, nil
}

//SplitLines => Split string by line
func SplitLines(stringifiedContent string) []string {
	var lines []string
	stringScanner := bufio.NewScanner(strings.NewReader(stringifiedContent))

	for stringScanner.Scan() {
		lines = append(lines, stringScanner.Text())
	}

	return lines
}

//CheckPhraseInSentence => Check whether phrase given exists in sentence
func CheckPhraseInSentence(sentence string, phrase map[string]bool) bool {
	for word := range phrase {
		if phrase[word] == true {
			if !strings.Contains(sentence, word) {
				return false
			}
		}

		if phrase[word] == false {
			if strings.Contains(sentence, word) {
				return false
			}
		}
	}

	return true
}

//CheckWordIndexInSentence => Check word index in sentence
func CheckWordIndexInSentence(sentence []string, word string) (int, error) {
	for splitSentenceKey, splitSentenceValue := range sentence {
		if splitSentenceValue == word {
			return splitSentenceKey, nil
		}
	}

	err := errors.New(word + "could not be found")
	return 0, err
}
