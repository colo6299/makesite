package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your age: ")
	age := int(parseStringToFloat(reader, "Your age must be a number."))

	fmt.Println("Please enter birth year: ")
	yearBirthf := parseStringToFloat(reader, "Your birth year must be a number.")

	//I just realized there's no way to use this info without already knowing the date
	//fmt.Println("Please enter your birth month: ")
	//monthBirthf := parseStringToFloat("Your birth month must be a number.")

	//fmt.Println("Please enter the day of your birth: ")
	//dayBirthf := parseStringToFloat("Your day of birth must be a number. (as in 'born on the 11th')")

	probableYear := age + int(yearBirthf)
	fmt.Printf("The year is either %d or %d.", probableYear, probableYear+1)
}

func parseStringToFloat(reader *bufio.Reader, errMessage string) float64 {
	inString, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	outputFloat, err := strconv.ParseFloat(inString[0:len(inString)-2], 64)
	if err != nil {
		panic(err)
	}
	return outputFloat
}