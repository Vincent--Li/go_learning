package main

import "testing"

func TestAddOne(t *testing.T){
	res := addOne(10)
	if res != 11 {
		t.Fatalf("error")
	}
	t.Logf("correct")
}
