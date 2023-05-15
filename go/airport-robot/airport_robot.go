package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(name string, greet Greeter) string {
	return "I can speak " + greet.LanguageName() + ": " + greet.Greet(name)
}

type Italian struct {
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

type Portuguese struct {
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}
