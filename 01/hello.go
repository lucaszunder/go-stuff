package main

func main() {
	Hello("Golang")
}

func Hello(name string) string {
	message := "Ola " + name
	return message
}
