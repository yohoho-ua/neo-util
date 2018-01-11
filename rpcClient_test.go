package main

import (
	"testing"
)

func TestGetInfo(t *testing.T) {
	//check actual amounts before
	expected1 := "109"
	expected2 := "957"

	actual1, actual2 := getInfo()
	
	if actual1 != expected1 {
		t.Errorf("getInfo returned unexpected NEO value: got %v want %v", actual1, expected1)
	}

	
	if actual2 != expected2 {
		t.Errorf("getInfo returned unexpected GAS value: got %v want %v", actual2, expected2)
	}
	
}