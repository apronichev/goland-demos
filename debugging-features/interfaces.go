package main

import "unicode"

type MyInterface interface {
	IsNumber(c rune) bool
	IsUpper(c rune) bool
	IsPunct(c rune) bool
	IsSymbol(c rune) bool
	IsLetter(c rune) bool
}

func IsNumber(c rune) bool {
	return unicode.IsNumber(c)
}

func IsUpper(c rune) bool {
	return unicode.IsUpper(c)
}

func IsPunct(c rune) bool {
	return unicode.IsPunct(c)
}

func IsSymbol(c rune) bool {
	return unicode.IsSymbol(c)
}

func IsLetter(c rune) bool {
	return unicode.IsLetter(c) || c == ' '
}
