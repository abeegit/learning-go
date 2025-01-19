package main

import (
	"fmt"
	"strings"
)

func wordCounter(s string) {
	stringsArray := strings.Split(s, " ")
	fmt.Println("The number of words in the string is ", len(stringsArray))
}