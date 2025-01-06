package main

func main() {
	Hello("Golang")
}

const greetings = "Hello "
const defaultGreeting = "Hello World"

func Hello(name string) string {
	if name == "" {
		return defaultGreeting
	}
	message := greetings + name
	return message
}
