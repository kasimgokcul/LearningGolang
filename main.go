package main

import (
	errorhandling "Dictionary/ErrorHandling"
	numberguessing "Dictionary/NumberGuessing"
	"Dictionary/dictionary"
	"Dictionary/interfaces"
)

func main() {
	dictionary.TheWords()
	numberguessing.Guessing()
	errorhandling.Error()
	interfaces.Demo2()
}
