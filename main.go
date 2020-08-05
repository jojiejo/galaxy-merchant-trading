package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/jojiejo/galaxy-merchant-trading/pkg"
)

func main() {
	var intergalacticToRomanNumeral = make(map[string]string)
	var materialWorth = make(map[string]float64)

	//Read the input text file
	inputContent, err := ioutil.ReadFile("assets/input.txt")
	if err != nil {
		fmt.Printf(err.Error())
	}

	//Convert []byte to string
	stringifiedContent := string(inputContent)

	//Split converted string
	splitStringifiedContent := strings.Split(stringifiedContent, "\r")

	//Start printing the output
	fmt.Println("-----------------")
	fmt.Println("Expecting Output")
	fmt.Println("-----------------")

	//Save the value of intergalactic numeral declaration
	for _, sentenceValue := range splitStringifiedContent {
		//Trim space the sentence & split it into words
		trimString := strings.TrimSpace(string(sentenceValue))
		splitItem := strings.Split(trimString, " ")
		lnSplitItem := len(splitItem)

		if len(splitItem) == 3 && strings.Contains(trimString, "is") {
			//Check whether the sentence contains 3 words, word "is", and the value of Roman Numeral exists
			//Check the third string is valid roman numerals
			_, doesStringExistInRomanNumerals := pkg.MapRomanToArabicNumeral[splitItem[2]]

			if doesStringExistInRomanNumerals {
				/*fmt.Printf("Sentence [%s] is intergalactic numeral declaration.\n", trimString)*/
				intergalacticToRomanNumeral[splitItem[0]] = splitItem[2]
			}
		} else if strings.Contains(sentenceValue, "is") && strings.Contains(sentenceValue, "Credits") && !strings.Contains(sentenceValue, "?") {
			//Check whether the sentence contains is and Credits
			if len(pkg.MapRomanToArabicNumeral) != 0 {
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
					arabicNumeral := pkg.ConvertRomanToArabic(romanNumeral)
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
				stringContainer, romanNumeral := pkg.ConvertInterGalacticToRoman(3, lnSplitItem-1, splitItem, intergalacticToRomanNumeral)

				//Form the string
				if stringContainer != "" && romanNumeral != "" {
					convertedRomanToArabicNumeral := pkg.ConvertRomanToArabic(romanNumeral)
					fmt.Printf("%s is %d.\n", strings.TrimSpace(stringContainer), convertedRomanToArabicNumeral)
				}

			} else {
				fmt.Println("I have no idea what you are talking about.")
			}
		} else if strings.Contains(sentenceValue, "how many") {
			//Check whether the sentence contains material worth question
			if splitItem[0] == "how" && splitItem[1] == "many" && splitItem[2] == "Credits" && splitItem[3] == "is" && splitItem[lnSplitItem-1] == "?" {
				//Convert the intergalactic numerals
				stringContainer, romanNumeral := pkg.ConvertInterGalacticToRoman(4, lnSplitItem-2, splitItem, intergalacticToRomanNumeral)

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
					convertedRomanToArabicNumeral := pkg.ConvertRomanToArabic(romanNumeral)
					countNumeralAndCredits := materialWorthValue * float64(convertedRomanToArabicNumeral)
					fmt.Printf("%s %s is %.0f Credits.\n", strings.TrimSpace(stringContainer), materialName, countNumeralAndCredits)
				}
			}
		} else {
			fmt.Println("Unidentified sentence.")
		}
	}
}
