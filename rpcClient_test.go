package main

import (
	"testing"
)

type MockSomething struct {
	MockTest func() bool
}
func (ms MockSomething) Test () bool {
	if ms.MockTest != nil {
		return ms.MockTest()
	}
	return false
}

func TestGetInfo(t *testing.T) {
	//check actual amounts before
	expected1 := "109"
	expected2 := "957"

	actual1, actual2 := GetInfo()
	
	if actual1 != expected1 {
		t.Errorf("getInfo returned unexpected NEO value: got %v want %v", actual1, expected1)
	}

	
	if actual2 != expected2 {
		t.Errorf("getInfo returned unexpected GAS value: got %v want %v", actual2, expected2)
	}
	
}