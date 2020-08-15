package pkg

import (
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

	//Save the value of intergalactic numeral declaration
	for _, sentenceValue := range splitStringifiedContent {
		//Trim space the sentence & split it into words
		trimString := strings.TrimSpace(string(sentenceValue))
		splitItem := strings.Split(trimString, " ")
		lnSplitItem := len(splitItem)

		// xxx is xxx case
		if len(splitItem) == 3 && strings.Contains(trimString, "is") {
			//Check whether the sentence contains 3 words, word "is", and the value of Roman Numeral exists
			//Check the third string is valid roman numerals
			_, doesStringExistInRomanNumerals := MapRomanToArabicNumeral[splitItem[2]]

			if doesStringExistInRomanNumerals {
				/*fmt.Printf("Sentence [%s] is intergalactic numeral declaration.\n", trimString)*/
				intergalacticToRomanNumeral[splitItem[0]] = splitItem[2]
			} else {
				err = errors.New("System could not recognize the character " + splitItem[2])
				return nil, err
			}
		}

		// xxx is xxx Credits case
		if strings.Contains(sentenceValue, "is") && strings.Contains(sentenceValue, "Credits") && !strings.Contains(sentenceValue, "?") {
			//Check whether the sentence contains is and Credits
			if len(intergalacticToRomanNumeral) != 0 {
				/*fmt.Printf("Sentence [%s] is material worth declaration.\n", trimString)*/

				// Find is & credit index
				isIndex := 0
				creditsIndex := 0
				for splitSentenceKey, splitSentenceValue := range splitItem {
					if splitSentenceValue == "is" {
						isIndex = splitSentenceKey
					} else if splitSentenceValue == "Credits" {
						creditsIndex = splitSentenceKey
					}
				}

				if isIndex != 0 && creditsIndex != 0 {
					//Form a roman numeral
					romanNumeral := ""
					for i := 0; i < isIndex-1; i++ {
						//Check the string is valid intergalactic numerals
						_, doesStringExistInIntergalacticNumerals := intergalacticToRomanNumeral[splitItem[i]]

						if doesStringExistInIntergalacticNumerals {
							romanNumeral += intergalacticToRomanNumeral[splitItem[i]]
						} else {
							err = errors.New("System could not recognize the character" + splitItem[i])
							return nil, err
						}
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
					materialWorthValue := creditsValue / float64(arabicNumeral)
					materialWorth[materialName] = materialWorthValue

					/*fmt.Printf("1 %s worth %.2f\n", materialName, materialWorth[materialName])*/
				}
			} else {
				err = errors.New("Intergalactic numeral declaration must be done at least once")
				return nil, err
			}
		}

		// how much is xxx ? case
		if strings.Contains(sentenceValue, "how much") {
			//Check whether the sentence contains convertion question
			if splitItem[0] == "how" && splitItem[1] == "much" && splitItem[2] == "is" && splitItem[lnSplitItem-1] == "?" {
				stringContainer, romanNumeral, _ := ConvertInterGalacticToRoman(3, lnSplitItem-1, splitItem, intergalacticToRomanNumeral)

				//Form the string
				if stringContainer != "" && romanNumeral != "" {
					convertedRomanToArabicNumeral, _ := ConvertRomanToArabic(romanNumeral)
					returnedString = append(returnedString, strings.TrimSpace(stringContainer)+" is "+strconv.Itoa(convertedRomanToArabicNumeral))
				}

			} else {
				returnedString = append(returnedString, "I have no idea what you are talking about")
			}
		}

		// how many Credits is xxx ? case
		if strings.Contains(sentenceValue, "how many") {
			//Check whether the sentence contains material worth question
			if splitItem[0] == "how" && splitItem[1] == "many" && splitItem[2] == "Credits" && splitItem[3] == "is" && splitItem[lnSplitItem-1] == "?" {
				//Convert the intergalactic numerals
				stringContainer, romanNumeral, _ := ConvertInterGalacticToRoman(4, lnSplitItem-2, splitItem, intergalacticToRomanNumeral)

				//Get material worth
				materialName := splitItem[lnSplitItem-2]
				var materialWorthValue float64
				_, doesMaterialExist := materialWorth[materialName]
				if doesMaterialExist {
					materialWorthValue = materialWorth[materialName]
				} else {
					err = errors.New("Material " + materialName + " is unidentified")
					return nil, err
				}

				//Calculate both
				if stringContainer != "" && romanNumeral != "" && doesMaterialExist {
					convertedRomanToArabicNumeral, _ := ConvertRomanToArabic(romanNumeral)
					countNumeralAndCredits := materialWorthValue * float64(convertedRomanToArabicNumeral)
					convertedCountNumeralAndCredits := strconv.FormatFloat(countNumeralAndCredits, 'f', -1, 64)
					returnedString = append(returnedString, strings.TrimSpace(stringContainer)+" "+materialName+" is "+convertedCountNumeralAndCredits+" Credits")
				}
			}
		}
	}

	return returnedString, nil
}
