package main

func main() {
}

const interactTimes = 5

func Interact(char string) string {
	var output string
	for i := 0; i < interactTimes; i += 1 {
		output += char
	}

	return output
}
