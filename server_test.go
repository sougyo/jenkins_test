package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Add(3, 5) returned %d, expected %d", result, expected)
	}

	result = Add(-1, 1)
	expected = 0
	if result != expected {
		t.Errorf("Add(-1, 1) returned %d, expected %d", result, expected)
	}
}
