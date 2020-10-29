package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

// InterfacesMain func
func InterfacesMain() {
	e := englishBot{}
	s := spanishBot{}
	printGreeting(e)
	printGreeting(s)
}

func printGreeting(eb bot) {
	fmt.Println(eb.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}
func (spanishBot) getGreeting() string {
	return "Hola"
}
