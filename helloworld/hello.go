package hello // switch to main

import "fmt"

const ( // can define multiple variable/constants using parentheses
	spanish = "Spanish"
	french = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// in go, functions starting with lowercase letters are PRIVATE, and uppercase are PUBLIC.
// this function is private because we don't want the internals of our algorithm exposed to the world.
func greetingPrefix(language string) (prefix string) { //! named return value (prefix string).
	// the value of prefix will be set to "" (str) or 0 (int) automatically by default
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return // just using return will return prefix
}

func main() {
	fmt.Println(Hello("Marcelo", "Spanish"))
}