package main

import "testing"

func TestDuplicateExists(t *testing.T) {
	f := containsDuplicate([]int{1, 2, 1})
	if f == false {
		t.Fail()
	}
}

func TestDuplicateDoesNotExists(t *testing.T) {
	f := containsDuplicate([]int{1, 2})
	if f == true {
		t.Fail()
	}
}
