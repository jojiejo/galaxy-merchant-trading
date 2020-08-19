package pkg

import "testing"

func TestProcessIntergalacticStatement(t *testing.T) {
	var result []string
	var err error
	var stringArgument = make([]string, 15)
	var resultLength int

	//Test for Empty Result
	result, _ = ProcessIntergalacticStatement(stringArgument)
	resultLength = len(result)
	if resultLength != 0 {
		t.Errorf("ProcessIntergalacticStatement() failed, expected %v, got %v", 0, resultLength)
	} else {
		t.Logf("ProcessIntergalacticStatement() success, expected %v, got %v", 0, resultLength)
	}

	//Overall cases
	stringArgument[0] = "glob is I"
	stringArgument[1] = "prok is V"
	stringArgument[2] = "pish is X"
	stringArgument[3] = "tegj is L"
	stringArgument[4] = "glob glob Silver is 34 Credits"
	stringArgument[5] = "glob prok Gold is 57800 Credits"
	stringArgument[6] = "pish pish Iron is 3910 Credits"
	stringArgument[7] = "how much is pish tegj glob glob ?"
	stringArgument[8] = "how many Credits is glob prok Silver ?"
	stringArgument[9] = "how many Credits is glob glob Gold ?"
	stringArgument[10] = "how many Credits is glob glob glob glob glob glob Gold ?"
	stringArgument[11] = "how many Credits is pish tegj glob Iron ?"
	stringArgument[12] = "Does pish tegj glob glob Iron has more Credits than glob glob Gold ?"
	stringArgument[13] = "Is glob prok larger than pish pish ?"
	stringArgument[14] = "how much wood could a woodchuck chuck if a woodchuck could chuck wood ?"
	result, err = ProcessIntergalacticStatement(stringArgument)
	resultLength = len(result)
	if resultLength != 8 {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", err)
	} else {
		t.Logf("ProcessIntergalacticStatement() success, expected %v, got %v", 5, resultLength)
	}

	//Negative case for xxx is xxx case
	stringArgument = make([]string, 1)
	stringArgument[0] = "glob is S"
	resultLength = len(result)
	result, err = ProcessIntergalacticStatement(stringArgument)
	if err != nil {
		t.Logf("ProcessIntergalacticStatement() success, got %v", err)
	} else {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", err)
	}

	//Negative case for xxx is xxx Credits
	stringArgument = make([]string, 1)
	stringArgument[0] = "prok Silver is 34 Credits"
	resultLength = len(result)
	result, err = ProcessIntergalacticStatement(stringArgument)
	if err != nil {
		t.Logf("ProcessIntergalacticStatement() success, got %v", err)
	} else {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", err)
	}

	//Negative case for xxx is xxx Credits
	stringArgument = make([]string, 2)
	stringArgument[0] = "glob is I"
	stringArgument[1] = "prok Silver is 34 Credits"
	resultLength = len(result)
	result, err = ProcessIntergalacticStatement(stringArgument)
	if result[0] != "Requested number is in invalid format" {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", result[0])
	} else {
		t.Logf("ProcessIntergalacticStatement() success, got %v", result[0])
	}

	//Negative case for how many Credits is xxx ?
	stringArgument = make([]string, 4)
	stringArgument[0] = "glob is I"
	stringArgument[1] = "glob glob Silver is 34 Credits"
	stringArgument[2] = "how much is pish tegj glob glob ?"
	stringArgument[3] = "how many Credits is glob prok Iron ?"
	resultLength = len(result)
	result, err = ProcessIntergalacticStatement(stringArgument)
	if result[0] != "Requested number is in invalid format" {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", result[0])
	} else {
		t.Logf("ProcessIntergalacticStatement() success, got %v", result[0])
	}

	//Is xxx larger than xxx ? [Larger Case]
	stringArgument = make([]string, 8)
	stringArgument[0] = "glob is I"
	stringArgument[1] = "prok is V"
	stringArgument[2] = "pish is X"
	stringArgument[3] = "tegj is L"
	stringArgument[4] = "glob glob Silver is 34 Credits"
	stringArgument[5] = "glob prok Gold is 57800 Credits"
	stringArgument[6] = "pish pish Iron is 3910 Credits"
	stringArgument[7] = "Is pish pish larger than glob prok ?"
	result, err = ProcessIntergalacticStatement(stringArgument)
	if result[0] != "pish pish is larger than glob prok" {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", result[0])
	} else {
		t.Logf("ProcessIntergalacticStatement() success, got %v", result[0])
	}

	//Is xxx larger than xxx ? [EqualCase]
	stringArgument = make([]string, 8)
	stringArgument[0] = "glob is I"
	stringArgument[1] = "prok is V"
	stringArgument[2] = "pish is X"
	stringArgument[3] = "tegj is L"
	stringArgument[4] = "glob glob Silver is 34 Credits"
	stringArgument[5] = "glob prok Gold is 57800 Credits"
	stringArgument[6] = "pish pish Iron is 3910 Credits"
	stringArgument[7] = "Is pish pish larger than pish pish ?"
	result, err = ProcessIntergalacticStatement(stringArgument)
	if result[0] != "pish pish is equal to pish pish" {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", result[0])
	} else {
		t.Logf("ProcessIntergalacticStatement() success, got %v", result[0])
	}

	//Is xxx larger than xxx ? [Negative Case]
	stringArgument = make([]string, 8)
	stringArgument[0] = "glob is I"
	stringArgument[1] = "prok is V"
	stringArgument[2] = "pish is X"
	stringArgument[3] = "tegj is L"
	stringArgument[4] = "glob glob Silver is 34 Credits"
	stringArgument[5] = "glob prok Gold is 57800 Credits"
	stringArgument[6] = "pish pish Iron is 3910 Credits"
	stringArgument[7] = "Is posh posh larger than pash pash ?"
	result, err = ProcessIntergalacticStatement(stringArgument)
	if result[0] != "Requested number is in invalid format" {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", result[0])
	} else {
		t.Logf("ProcessIntergalacticStatement() success, got %v", result[0])
	}

	//Does xxx xxx has more Credits than xxx xxx ? [Larger Case]
	stringArgument = make([]string, 8)
	stringArgument[0] = "glob is I"
	stringArgument[1] = "prok is V"
	stringArgument[2] = "pish is X"
	stringArgument[3] = "tegj is L"
	stringArgument[4] = "glob glob Silver is 34 Credits"
	stringArgument[5] = "glob prok Gold is 57800 Credits"
	stringArgument[6] = "pish pish Iron is 3910 Credits"
	stringArgument[7] = "Does glob glob Gold has more Credits than pish tegj glob glob Iron ?"
	result, err = ProcessIntergalacticStatement(stringArgument)
	if result[0] != "glob glob Gold has more Credits than pish tegj glob glob Iron" {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", result[0])
	} else {
		t.Logf("ProcessIntergalacticStatement() success, got %v", result[0])
	}

	//Does xxx xxx has more Credits than xxx xxx ? [Equal Case]
	stringArgument = make([]string, 8)
	stringArgument[0] = "glob is I"
	stringArgument[1] = "prok is V"
	stringArgument[2] = "pish is X"
	stringArgument[3] = "tegj is L"
	stringArgument[4] = "glob glob Silver is 34 Credits"
	stringArgument[5] = "glob prok Gold is 57800 Credits"
	stringArgument[6] = "pish pish Iron is 3910 Credits"
	stringArgument[7] = "Does pish tegj glob glob Iron has more Credits than pish tegj glob glob Iron ?"
	result, err = ProcessIntergalacticStatement(stringArgument)
	if result[0] != "pish tegj glob glob Iron has same Credits as pish tegj glob glob Iron" {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", result[0])
	} else {
		t.Logf("ProcessIntergalacticStatement() success, got %v", result[0])
	}

	//Does xxx xxx has more Credits than xxx xxx ? [Negative Case]
	stringArgument = make([]string, 8)
	stringArgument[0] = "glob is I"
	stringArgument[1] = "prok is V"
	stringArgument[2] = "pish is X"
	stringArgument[3] = "tegj is L"
	stringArgument[4] = "glob glob Silver is 34 Credits"
	stringArgument[5] = "glob prok Gold is 57800 Credits"
	stringArgument[6] = "pish pish Iron is 3910 Credits"
	stringArgument[7] = "Does glab glab Gold has more Credits than pish tegj glab glab Iron ?"
	result, err = ProcessIntergalacticStatement(stringArgument)
	if result[0] != "Requested number is in invalid format" {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", result[0])
	} else {
		t.Logf("ProcessIntergalacticStatement() success, got %v", result[0])
	}
}

func TestCheckWordInSentence(t *testing.T) {
	//var result int
	var err error
	var sentence = make([]string, 15)
	var word string

	//Negative Case
	sentence = make([]string, 6)
	sentence[0] = "pish"
	sentence[1] = "pish"
	sentence[2] = "Iron"
	sentence[3] = "is"
	sentence[4] = "3910"
	sentence[5] = "Credits"
	word = "are"
	_, err = CheckWordIndexInSentence(sentence, word)
	if err != nil {
		t.Logf("CheckWordInSentence() success, got %v", err)
	} else {
		t.Errorf("CheckWordInSentence() failed, got %v", err)
	}
}
