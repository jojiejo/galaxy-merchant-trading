package processor

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jojiejo/galaxy-merchant-trading/converter"
)

func ProcessStringForeachLine(sentenceValue string) {
	var intergalacticToRomanNumeral = make(map[string]string)
	var materialWorth = make(map[string]float64)

	//Trim space the sentence & split it into words
	trimString := strings.TrimSpace(string(sentenceValue))
	splitItem := strings.Split(trimString, " ")
	lnSplitItem := len(splitItem)

	if len(splitItem) == 3 && strings.Contains(trimString, "is") {
		//Check whether the sentence contains 3 words, word "is", and the value of Roman Numeral exists
		//Check the third string is valid roman numerals
		_, doesStringExistInRomanNumerals := converter.MapRomanToArabicNumeral[splitItem[2]]

		if doesStringExistInRomanNumerals {
			fmt.Printf("Sentence [%s] is intergalactic numeral declaration.\n", trimString)
			intergalacticToRomanNumeral[splitItem[0]] = splitItem[2]
		}
	} else if strings.Contains(sentenceValue, "is") && strings.Contains(sentenceValue, "Credits") && !strings.Contains(sentenceValue, "?") {
		//Check whether the sentence contains is and Credits
		if len(converter.MapRomanToArabicNumeral) != 0 {
			fmt.Printf("Sentence [%s] is material worth declaration.\n", trimString)

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
						fmt.Println("System could not recognize the character.")
					}
				}

				//Form the name of material
				materialName := splitItem[isIndex-1]

				//Form the value of credits
				creditsValue, err := strconv.ParseFloat(splitItem[creditsIndex-1], 64)
				if err != nil {
					fmt.Printf(err.Error())
				}

				//Calculation of material worth
				arabicNumeral := converter.ConvertRomanToArabic(romanNumeral)
				materialWorthValue := creditsValue / float64(arabicNumeral)
				materialWorth[materialName] = materialWorthValue

				/*fmt.Printf("1 %s worth %.2f\n", materialName, materialWorth[materialName])*/
			}
		} else {
			fmt.Println("Intergalactic numeral declaration must be done at least once.")
		}
	} else if strings.Contains(sentenceValue, "how much") {
		//Check whether the sentence contains convertion question
		if splitItem[0] == "how" && splitItem[1] == "much" && splitItem[2] == "is" && splitItem[lnSplitItem-1] == "?" {
			stringContainer := ""
			romanNumeral := ""

			for i := 3; i < lnSplitItem-1; i++ {
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

			//Form the string
			if stringContainer != "" && romanNumeral != "" {
				convertedRomanToArabicNumeral := converter.ConvertRomanToArabic(romanNumeral)
				fmt.Printf("%s is %d.\n", strings.TrimSpace(stringContainer), convertedRomanToArabicNumeral)
			}

		} else {
			fmt.Println("I have no idea what you are talking about.")
		}
	} else if strings.Contains(sentenceValue, "how many") {
		//Check whether the sentence contains material worth question
		if splitItem[0] == "how" && splitItem[1] == "many" && splitItem[2] == "Credits" && splitItem[3] == "is" && splitItem[lnSplitItem-1] == "?" {
			//Convert the intergalactic numerals
			stringContainer := ""
			romanNumeral := ""

			for i := 4; i < lnSplitItem-2; i++ {
				//Check the string is a valid intergalactic numerals
				_, doesStringExistInIntergalacticNumerals := intergalacticToRomanNumeral[splitItem[i]]

				if doesStringExistInIntergalacticNumerals {
					stringContainer += splitItem[i] + " "
					romanNumeral += intergalacticToRomanNumeral[splitItem[i]]
				} else {
					fmt.Printf("One of intergalactic numerals is unidentified.")
					stringContainer = ""
					romanNumeral = ""
					break
				}
			}

			//Get material worth
			materialName := splitItem[lnSplitItem-2]
			var materialWorthValue float64
			_, doesMaterialExist := materialWorth[materialName]
			if doesMaterialExist {
				materialWorthValue = materialWorth[materialName]
			} else {
				fmt.Printf("Material [%s] is unidentified.\n", materialName)
			}

			//Calculate both
			if stringContainer != "" && romanNumeral != "" && doesMaterialExist {
				convertedRomanToArabicNumeral := converter.ConvertRomanToArabic(romanNumeral)
				countNumeralAndCredits := materialWorthValue * float64(convertedRomanToArabicNumeral)
				fmt.Printf("%s %s is %.0f Credits.\n", strings.TrimSpace(stringContainer), materialName, countNumeralAndCredits)
			}
		}
	} else {
		fmt.Println("Unidentified sentence.")
	}
}
