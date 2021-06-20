package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Duy")
	want := "Hello, Duy"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
