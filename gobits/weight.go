package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	total := 5
	fmt.Printf("Please enter %d weights: \n", total)
	var average float64
	for i := 0; i < total; i++ {
		weight, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		weightFloat, err := strconv.ParseFloat(weight[0:len(weight)-2], 64)
		if err != nil {
			panic(err)
		}
		average += weightFloat / float64(total)
	}
	fmt.Printf("The average weight is %f units.", average)
}
