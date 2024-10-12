package main

import "fmt"

var prefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		prefix = "Hola, "
	}
	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("Marcelo", "Spanish"))
}