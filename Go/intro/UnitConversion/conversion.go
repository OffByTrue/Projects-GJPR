package UnitConversion

import (
	"fmt"
	"strconv"
	"strings"
)

// what we have to do is a small Parser for input [Args]
// and then some methods to convert the units
// I should run the Program from this file only as it includes Args;

func Parse(action []string) {
	fmt.Println(action)
	var act [4]float64

	if strings.ToLower(action[0]) == "convert" {
		act[0] = 1
	} else {
		act[0] = 2
	}

	val, err := strconv.ParseFloat(action[1], 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	act[1] = val

	// I will do if elseif check but i will also see how to use enums here Later

	var from = strings.ToLower(action[2])
	if from == "km" {
		act[2] = 1
	} else if from == "m" {
		act[2] = 2
	} else if from == "cm" {
		act[2] = 3
	} else if from == "mm" {
		act[2] = 4
	} else if from == "miles" {
		act[2] = 5
	} else {
		act[2] = 0
	}

	var to = strings.ToLower(action[3])
	if to == "km" {
		act[3] = 1
	} else if to == "m" {
		act[3] = 2
	} else if to == "cm" {
		act[3] = 3
	} else if to == "mm" {
		act[3] = 4
	} else if from == "miles" {
		act[2] = 5
	} else {
		act[3] = 0
	}

	fmt.Println(act)
	var newVal float64
	if act[2] == act[3] {
		fmt.Println("cant convert to itself.")
		return
	} else if act[2] == 1 && act[3] == 2 {
		newVal = val * 1000
	} else if act[2] == 2 && act[3] == 3 {
		newVal = val * 100
	} else if act[2] == 3 && act[3] == 4 {
		newVal = val * 10
	} else if act[2] == 1 && act[3] == 5 {
		newVal = val * 0.621371
	} else if act[2] == 2 && act[3] == 5 {
		newVal = val * 0.000621371
	} else if act[2] == 2 && act[3] == 3 {
		fmt.Println("Preferred to do to km first")
	}

	fmt.Println("from : ", val, " ", action[2])
	fmt.Println("to : ", newVal, " ", action[3])

}
