package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func getNumbers(command string, operand string) ([]int, error) {
	numbersAsStrings := strings.Split(command, operand)
	fmt.Println(reflect.TypeOf(numbersAsStrings))

	number1, error1 := strconv.Atoi(numbersAsStrings[0])
	if error1 != nil {
		return nil, fmt.Errorf("number 1 is not an integer")
	}
	fmt.Println(len(numbersAsStrings), numbersAsStrings[0], numbersAsStrings[1], number1, 2)

	number2, error2 := strconv.Atoi(numbersAsStrings[1])
	fmt.Println(number2)
	if error2 != nil {
		return nil, fmt.Errorf("number 2 is not an integer")
	}

	fmt.Println("The numbers are ", number1, number2)
	
	return []int{number1, number2}, nil
}

func calculate(command string) int {
	var isAddition, isSubtraction, isMultiplication, isDivision bool

	isAddition = strings.Contains(command, "+")
	isSubtraction = strings.Contains(command, "-")
	isMultiplication = strings.Contains(command, "*")
	isDivision = strings.Contains(command, "/")

	if isAddition {
		numbers, _ := getNumbers(command, "+")
		return numbers[0] + numbers[1]
	}

	if isSubtraction {
		numbers, _ := getNumbers(command, "-")
		return numbers[0] + numbers[1]
	}

	if isMultiplication {
		fmt.Println("This is a multiplication")
		numbers, _ := getNumbers(command, "*")
		
		return numbers[0] * numbers[1]
	}

	if isDivision {
		numbers, _ := getNumbers(command, "/")
		return numbers[0] + numbers[1]
	}

	return 0
}


// I had trouble writing this simple program because I kept getting an array out of bounds error. 
func getCalculatorInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	commandWithDelimiter, _ := reader.ReadString('\n')
	result := calculate(strings.Split(commandWithDelimiter, "\n")[0])
	fmt.Println(result)
}
