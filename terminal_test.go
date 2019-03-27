package main

import "testing"

func TestGetTerminalSize(t *testing.T) {
	col, row, err := GetTerminalSize()
	if err != nil || col == 0 || row == 0 {
		t.Fatal(err)
	}
}
