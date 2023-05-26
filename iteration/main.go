package main

func Repeat(character string, quantity int) string {
	var repeats string

	for i := 0; i < quantity; i++ {
		repeats += character
	}

	return repeats
}
