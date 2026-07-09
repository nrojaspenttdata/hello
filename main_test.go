package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fmt.Println("Hello, World!")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	got := buf.String()
	want := "Hello, World!\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
