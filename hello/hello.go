package main

import (
	"fmt"
)

type Lang int

const (
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
	english       = Lang(iota)
	spanish
	french
)

func Hello(name string, l Lang) string {
	var prefix string
	switch l {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}

	if name == "" {
		name = "World"
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", english))
}
