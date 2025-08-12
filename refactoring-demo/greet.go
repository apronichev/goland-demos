package main

type Speaker interface {
	Present(name string) string
	Say(name string) string
	Inspire() string
}
