package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var mapRomanToArabicNumeral = make(map[string]int)

func init() {
	mapRomanToArabicNumeral = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
}

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

		//Check whether the sentence contains 3 words, word "is", and the value of Roman Numeral exists
		if len(splitItem) == 3 && strings.Contains(trimString, "is") {
			//Check the third string is valid roman numerals
			_, doesStringExistInRomanNumerals := mapRomanToArabicNumeral[splitItem[2]]

			if doesStringExistInRomanNumerals {
				/*fmt.Printf("Sentence [%s] is intergalactic numeral declaration.\n", trimString)*/
				intergalacticToRomanNumeral[splitItem[0]] = splitItem[2]
			}
		}

		//Check whether the sentence contains is and Credits
		if strings.Contains(sentenceValue, "is") && strings.Contains(sentenceValue, "Credits") && !strings.Contains(sentenceValue, "?") {
			if len(mapRomanToArabicNumeral) != 0 {
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
							fmt.Printf("System could not recognize the character.\n")
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
					arabicNumeral := converter.convertRomanToArabic(romanNumeral, mapRomanToArabicNumeral)
					materialWorthValue := creditsValue / float64(arabicNumeral)
					materialWorth[materialName] = materialWorthValue

					/*fmt.Printf("1 %s worth %.2f\n", materialName, materialWorth[materialName])*/
				}
			} else {
				fmt.Printf("Intergalactic numeral declaration must be done at least once.\n")
			}
		}

		//Check whether the sentence contains convertion question
		if strings.Contains(sentenceValue, "how much") {
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
						fmt.Printf("One of intergalactic numerals is unidentified.\n", trimString)
						stringContainer = ""
						romanNumeral = ""
						break
					}
				}

				//Form the string
				if stringContainer != "" && romanNumeral != "" {
					convertedRomanToArabicNumeral := converter.convertRomanToArabic(romanNumeral, mapRomanToArabicNumeral)
					fmt.Printf("%s is %d.\n", strings.TrimSpace(stringContainer), convertedRomanToArabicNumeral)
				}

			} else {
				fmt.Printf("I have no idea what you are talking about.\n")
			}
		}

		//Check whether the sentence contains material worth question
		if strings.Contains(sentenceValue, "how many") {
			if splitItem[0] == "how" && splitItem[1] == "many" && splitItem[2] == "Credits" && splitItem[3] == "is" && splitItem[lnSplitItem-1] == "?" {
				//Convert the intergalactic numerals
				stringContainer := ""
				romanNumeral := ""

				for i := 4; i < lnSplitItem-2; i++ {
					//Check the string is valid intergalactic numerals
					_, doesStringExistInIntergalacticNumerals := intergalacticToRomanNumeral[splitItem[i]]

					if doesStringExistInIntergalacticNumerals {
						stringContainer += splitItem[i] + " "
						romanNumeral += intergalacticToRomanNumeral[splitItem[i]]
					} else {
						fmt.Printf("One of intergalactic numerals is unidentified.\n", trimString)
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
					convertedRomanToArabicNumeral := converter.convertRomanToArabic(romanNumeral)
					countNumeralAndCredits := materialWorthValue * float64(convertedRomanToArabicNumeral)
					fmt.Printf("%s %s is %.0f Credits.\n", strings.TrimSpace(stringContainer), materialName, countNumeralAndCredits)
				}
			}
		}
	}
}
