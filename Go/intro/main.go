package main

import (
	"intro/UnitConversion"
	"os"
)

func main() {
	//GuessingGame.Menu()
	UnitConversion.Parse(os.Args[1:])
}
