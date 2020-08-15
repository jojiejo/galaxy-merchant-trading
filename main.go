package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jojiejo/galaxy-merchant-trading/pkg"
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

	//Process input string
	processedString, err := pkg.ProcessIntergalacticStatement(splitStringifiedContent)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	//Start printing the output
	fmt.Println("-----------------")
	fmt.Println("Output")
	fmt.Println("-----------------")
	for i := 0; i < len(processedString); i++ {
		fmt.Println(processedString[i])
	}
}
