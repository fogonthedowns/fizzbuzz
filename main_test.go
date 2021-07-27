package main

import (
	"reflect"
	"testing"
)

func TestFun(t *testing.T) {
	in := 3
	out := fizzBuzz(in)
	want := []string{"1", "2", "Fizz"}

	// test two slices are equal
	if !reflect.DeepEqual(out, want) {
		t.Errorf("got %v, want %v", out, want)
	}
}
