package main

func main() {
	Hello("Golang")
}

const greetings = "Ola "
const defaultGreeting = "Ola Mundo"

func Hello(name string) string {
	if name == "" {
		return defaultGreeting
	}
	message := greetings + name
	return message
}
