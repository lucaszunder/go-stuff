package main

import (
	"fmt"
	"io"
)

func Greetings(writter io.Writer, name string) {
	fmt.Fprintf(writter, "Hello, %s", name)
}
