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

func TestFilterRowsIsCaseInsensitive(t *testing.T) {
	rows := []string{"abc\t2026-01-01\tWeb-App", "def\t2026-01-02\tworker"}
	got := filterRows(rows, "web")
	if !reflect.DeepEqual(got, []string{rows[0]}) {
		t.Fatalf("got %v, want %v", got, []string{rows[0]})
	}
}

func TestFilterRowsWithoutMatchKeepsAll(t *testing.T) {
	rows := []string{"one", "two"}
	if got := filterRows(rows, ""); !reflect.DeepEqual(got, rows) {
		t.Fatalf("got %v, want %v", got, rows)
	}
}
