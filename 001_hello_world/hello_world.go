package main

import "fmt"

func greetingPrefix(lang string) string {
	prefixes := map[string]string{
		"en": "Hello, ",
		"es": "Hola, ",
		"fr": "Alo, ",
	}

	value, ok := prefixes[lang]

	if !ok {
		return prefixes["en"]
	}

	return value
}

func Hello(name string, lang string) string {
	prefix := greetingPrefix(lang)

	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("%s%s", prefix, name)
}

func main() {
	fmt.Println(Hello("Fikri", "en"))
}
