package pkg

import "testing"

func TestProcessIntergalacticStatement(t *testing.T) {
	var result []string
	var err error
	var stringArgument = make([]string, 12)
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
	stringArgument[9] = "how many Credits is glob prok Gold ?"
	stringArgument[10] = "how many Credits is glob prok Iron ?"
	stringArgument[11] = "how much wood could a woodchuck chuck if a woodchuck could chuck wood ?"
	result, err = ProcessIntergalacticStatement(stringArgument)
	resultLength = len(result)
	if resultLength != 5 {
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
	if err != nil {
		t.Logf("ProcessIntergalacticStatement() success, got %v", err)
	} else {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", err)
	}

	//Negative case for how many Credits is xxx ?
	stringArgument = make([]string, 4)
	stringArgument[0] = "glob is I"
	stringArgument[1] = "glob glob Silver is 34 Credits"
	stringArgument[2] = "how much is pish tegj glob glob ?"
	stringArgument[3] = "how many Credits is glob prok Iron ?"
	resultLength = len(result)
	result, err = ProcessIntergalacticStatement(stringArgument)
	if err != nil {
		t.Logf("ProcessIntergalacticStatement() success, got %v", err)
	} else {
		t.Errorf("ProcessIntergalacticStatement() failed, got %v", err)
	}
}
