package airportrobot

import (
	"fmt"
)

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, any Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", any.LanguageName(), any.Greet(name))
}

type Italian struct {
	languageName string
	name         string
}

func (i Italian) LanguageName() string {
	i.languageName = "Italian"
	return i.languageName
}

func (i Italian) Greet(name string) string {
	i.name = name
	return fmt.Sprintf("Ciao %v!", i.name)
}

type Portuguese struct {
	languageName string
	name         string
}

func (i Portuguese) LanguageName() string {
	i.languageName = "Portuguese"
	return i.languageName
}

func (i Portuguese) Greet(name string) string {
	i.name = name
	return fmt.Sprintf("Olá %v!", i.name)
}
