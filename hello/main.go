package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	languages := map[string]string{
		"en-US": "Hello, ",
		"pt-BR": "Olá, ",
		"fr":    "Bonjour, ",
	}

	return languages[language] + name + "!"
}

func main() {
	fmt.Println(Hello("world", "en-US"))
	fmt.Println(Hello("mundo", "pt-BR"))
	fmt.Println(Hello("monde", "fr"))
}
