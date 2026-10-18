package dictionary

import (
	"errors"
	"maps"
	"testing"
)

func TestCreate(t *testing.T) {
	wantD := make(Dictionary)
	dictionary := Create()
	checkMap := *dictionary

	if !maps.Equal(wantD, checkMap) {
		t.Errorf("expected as map")
	}
}

func TestAdd(t *testing.T) {
	wantMap := Dictionary{
		"Bird": "Птица",
		"Car":  "Машина",
	}
	var wantErr error = nil

	dictionary := Dictionary{
		"Bird": "Птица",
	}
	err := dictionary.Add("Car", "Машина")

	if !errors.Is(err, wantErr) {
		t.Fatalf("expected error %v, got %v", wantErr, err)
	}

	if !maps.Equal(dictionary, wantMap) {
		t.Errorf("expected map %v, got %v", wantMap, dictionary)
	}
}

func TestDelete(t *testing.T) {
	wantMap := Dictionary{
		"Bird": "Птица",
	}
	var wantErr error = nil

	dictionary := Dictionary{
		"Bird": "Птица",
		"Car":  "Машина",
	}
	err := dictionary.Delete("Car")

	if !errors.Is(err, wantErr) {
		t.Fatalf("expected error %v, got %v", wantErr, err)
	}

	if !maps.Equal(dictionary, wantMap) {
		t.Errorf("expected map %v, got %v", wantMap, dictionary)
	}
}

func TestReturnInfo(t *testing.T) {
	wantMap := Dictionary{
		"Bird": "Птица",
	}
	dictionary := Dictionary{
		"Bird": "Птица",
	}

	result := dictionary.ReturnInfo()

	if !maps.Equal(wantMap, result) {
		t.Errorf("expected map %v, got %v", wantMap, result)
	}
}
