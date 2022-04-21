package main

import (
	errorhandling "Dictionary/ErrorHandling"
	numberguessing "Dictionary/NumberGuessing"
	"Dictionary/dictionary"
)

func main() {
	dictionary.TheWords()
	numberguessing.Guessing()
	errorhandling.Error()
}
