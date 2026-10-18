package service

import (
	"testing"
)

func TestModifyWords(t *testing.T) {
	tr := "   Hello   "
	r := "    Car    "

	wantT := "Hello"
	wantR := "Car"

	tranclate, responce := modifyWords(tr, r)

	if tranclate != wantT {
		t.Errorf("expected word %v, got %v", wantT, tranclate)
	}

	if responce != wantR {
		t.Errorf("expected word %v, got %v", wantR, responce)
	}
}
