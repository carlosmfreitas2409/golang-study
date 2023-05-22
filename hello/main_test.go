package main

import "testing"

func TestHello(t *testing.T) {
	testReturnedMessage := func(t *testing.T, result, expected string) {
		t.Helper()

		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	}

	t.Run("should be able to say hello to the given name", func(t *testing.T) {
		result := Hello("Carlos")
		expected := "Hello, Carlos!"

		testReturnedMessage(t, result, expected)
	})

	t.Run("should be able to say 'Hello, world!' when the name is not provided", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, world!"

		testReturnedMessage(t, result, expected)
	})
}
