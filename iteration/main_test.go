package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeats := Repeat("a", 3)
	expected := "aaa"

	if repeats != expected {
		t.Errorf("Expected '%s', got '%s'", expected, repeats)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeats := Repeat("a", 5)
	fmt.Println(repeats)
	// Output: aaaaa
}
