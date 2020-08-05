package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jojiejo/galaxy-merchant-trading/processor"
)

func main() {
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
		processor.ProcessStringForeachLine(sentenceValue)
	}
}
