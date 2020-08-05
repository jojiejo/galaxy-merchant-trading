package pkg

import (
	"strings"
	"testing"
)

func TestConvertRomanToArabic(t *testing.T) {
	var result int

	//Test for Empty Result
	result = ConvertRomanToArabic("")
	if result != 0 {
		t.Errorf("ConvertRomanToArabic() failed, expected %v, got %v", 0, result)
	} else {
		t.Logf("ConvertRomanToArabic() success, expected %v, got %v", 0, result)
	}

	//Test for Valid Argument [Small Number]
	result = ConvertRomanToArabic("III")
	if result != 3 {
		t.Errorf("ConvertRomanToArabic() failed, expected %v, got %v", 0, result)
	} else {
		t.Logf("ConvertRomanToArabic() success, expected %v, got %v", 0, result)
	}

	//Test for Valid Argument [Large Number]
	result = ConvertRomanToArabic("MCMXLIV")
	if result != 1944 {
		t.Errorf("ConvertRomanToArabic() failed, expected %v, got %v", 0, result)
	} else {
		t.Logf("ConvertRomanToArabic() success, expected %v, got %v", 0, result)
	}
}

func TestConvertInterGalacticToArabic(t *testing.T) {
	var resultStringContainer, resultRomanNumeral string
	var intergalacticToRomanNumeral = make(map[string]string)
	var splitItem []string
	var lnSplitItem int

	//Test for Empty Result
	resultStringContainer, resultRomanNumeral = ConvertInterGalacticToRoman(0, 0, splitItem, intergalacticToRomanNumeral)
	if resultStringContainer != "" || resultRomanNumeral != "" {
		t.Errorf("ConvertInterGalacticToRoman() failed, expected %v and %v, got %v and %v", "", "", resultStringContainer, resultRomanNumeral)
	} else {
		t.Logf("ConvertInterGalacticToRoman() success, expected %v and %v, got %v and %v", "", "", resultStringContainer, resultRomanNumeral)
	}

	//Test for Valid Argument
	splitItem = strings.Split("how many Credits is glob prok Silver ?", " ")
	lnSplitItem = len(splitItem)
	intergalacticToRomanNumeral["glob"] = "I"
	intergalacticToRomanNumeral["prok"] = "V"

	resultStringContainer, resultRomanNumeral = ConvertInterGalacticToRoman(4, lnSplitItem-2, splitItem, intergalacticToRomanNumeral)
	if resultStringContainer != "glob prok " || resultRomanNumeral != "IV" {
		t.Errorf("ConvertInterGalacticToRoman() failed, expected %v and %v, got %v and %v", "glob prok", "IV", resultStringContainer, resultRomanNumeral)
	} else {
		t.Logf("ConvertInterGalacticToRoman() success, expected %v and %v, got %v and %v", "glob prok", "IV", resultStringContainer, resultRomanNumeral)
	}

	//Negative Case
	splitItem = strings.Split("how many Credits is glob prik Silver ?", " ")
	lnSplitItem = len(splitItem)
	intergalacticToRomanNumeral["glob"] = "I"
	intergalacticToRomanNumeral["prok"] = "V"

	resultStringContainer, resultRomanNumeral = ConvertInterGalacticToRoman(4, lnSplitItem-2, splitItem, intergalacticToRomanNumeral)
	if resultStringContainer != "" || resultRomanNumeral != "" {
		t.Errorf("ConvertInterGalacticToRoman() failed, expected %v and %v, got %v and %v", "", "", resultStringContainer, resultRomanNumeral)
	} else {
		t.Logf("ConvertInterGalacticToRoman() success, expected %v and %v, got %v and %v", "", "", resultStringContainer, resultRomanNumeral)
	}
}
