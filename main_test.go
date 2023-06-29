package main_test

import (
	"testing"

	"github.com/apfohl/go-pipe"
)

func TestHello(t *testing.T) {
	input := "world"

	expected := "Hello " + input + "!"

	result := main.Hello(input)

	if result != expected {
		t.Errorf("Expected %s, but got %s!\n", expected, result)
	}
}
