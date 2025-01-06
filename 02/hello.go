package main

func main() {
	Hello("Golang", "")
}

const defaultLanguage = "Hello "

const defaultGreeting = "Hello World"

var greetings = map[string]string{
	"portuguese": "Oi ",
	"english":    "Hello ",
	"spanish":    "Ola ",
}

var defaultGreetings = map[string]string{
	"portuguese": "Oi Mundo",
	"english":    "Hello World",
	"spanish":    "Ola Mundo",
}

func Hello(name, language string) string {
	if name == "" && language == "" {
		return defaultGreeting
	}
	if name == "" {
		return defaultGreetings[language]
	}
	if language == "" {
		return defaultLanguage + name
	}
	return greetings[language] + name
}
