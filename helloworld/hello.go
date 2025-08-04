package main

import "fmt"

const (
	russian = "Russian"
	french = "French"
	englishHelloPrefix = "Hello, "
	russianHelloPrefix = "Привет, "
	frenchHelloPrefix = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case russian: 
		prefix = russianHelloPrefix
	case french: 
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
