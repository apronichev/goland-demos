package main

type MyInterface interface {
	IsNumber(c rune) bool
	IsUpper(c rune) bool
	IsPunct(c rune) bool
	IsSymbol(c rune) bool
	IsLetter(c rune) bool
}
