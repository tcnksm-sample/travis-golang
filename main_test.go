package main

import (
	"testing"
)

func HelloTest(t *testing.T) {
	result := hello()
	if result != "hello" {
		t.Errorf("hello() = %s, want %s", result, "hello")
	}
}
