package main

import (
	"reflect"
	"testing"
)

func TestNonEmptyLines(t *testing.T) {
	got := nonEmptyLines("abc\n\n  def  \n")
	want := []string{"abc", "def"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestNonEmptyLinesEmptyInput(t *testing.T) {
	if got := nonEmptyLines(" \n\t\n"); len(got) != 0 {
		t.Fatalf("got %v, want no rows", got)
	}
}
