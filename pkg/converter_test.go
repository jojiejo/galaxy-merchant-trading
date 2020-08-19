package pkg

import (
	"strings"
	"testing"
)

func TestConvertRomanToArabic(t *testing.T) {
	var result int
	var err error

	//Test for Empty Result
	result, _ = ConvertRomanToArabic("")
	if result != 0 {
		t.Errorf("ConvertRomanToArabic() failed, expected %v, got %v", 0, result)
	} else {
		t.Logf("ConvertRomanToArabic() success, expected %v, got %v", 0, result)
	}

	//Test for Valid Argument [Small Number]
	result, _ = ConvertRomanToArabic("III")
	if result != 3 {
		t.Errorf("ConvertRomanToArabic() failed, expected %v, got %v", 3, result)
	} else {
		t.Logf("ConvertRomanToArabic() success, expected %v, got %v", 3, result)
	}

	//Test for Valid Argument [Large Number]
	result, _ = ConvertRomanToArabic("MCMXLIV")
	if result != 1944 {
		t.Errorf("ConvertRomanToArabic() failed, expected %v, got %v", 1944, result)
	} else {
		t.Logf("ConvertRomanToArabic() success, expected %v, got %v", 1944, result)
	}

	//Test for Valid Argument Consecutive Chars
	result, _ = ConvertRomanToArabic("XXXIX")
	if result != 39 {
		t.Errorf("ConvertRomanToArabic() failed, expected %v, got %v", 39, result)
	} else {
		t.Logf("ConvertRomanToArabic() success, expected %v, got %v", 39, result)
	}

	//Negative Case for 3 Consecutive Chars
	result, err = ConvertRomanToArabic("IIIX")
	if err != nil {
		t.Logf("ConvertRomanToArabic() success, got %v", err)
	} else {
		t.Errorf("ConvertRomanToArabic() failed, got %v", err)
	}

	//Negative Case for 2 Consecutive Chars
	result, err = ConvertRomanToArabic("IIX")
	if err != nil {
		t.Logf("ConvertRomanToArabic() success, got %v", err)
	} else {
		t.Errorf("ConvertRomanToArabic() failed, got %v", err)
	}

	//Negative Case for Consecutive Chars
	result, err = ConvertRomanToArabic("DDD")
	if err != nil {
		t.Logf("ConvertRomanToArabic() success, got %v", err)
	} else {
		t.Errorf("ConvertRomanToArabic() failed, got %v", err)
	}
}

func TestConvertInterGalacticToArabic(t *testing.T) {
	var resultStringContainer, resultRomanNumeral string
	var intergalacticToRomanNumeral = make(map[string]string)
	var splitItem []string
	var lnSplitItem int
	var err error

	//Test for Empty Result
	resultStringContainer, resultRomanNumeral, _ = ConvertInterGalacticToRoman(0, 0, splitItem, intergalacticToRomanNumeral)
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

	resultStringContainer, resultRomanNumeral, _ = ConvertInterGalacticToRoman(4, lnSplitItem-2, splitItem, intergalacticToRomanNumeral)
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

	resultStringContainer, resultRomanNumeral, err = ConvertInterGalacticToRoman(4, lnSplitItem-2, splitItem, intergalacticToRomanNumeral)
	if err != nil {
		t.Logf("ConvertInterGalacticToRoman() success, got %v", err)
	} else {
		t.Errorf("ConvertInterGalacticToRoman() failed, got %v", err)
	}
}
