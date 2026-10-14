package service

import (
	"errors"
	"exerciser/internal/storage/dictionary"
	"maps"
	"testing"
)

var (
	errFake = errors.New("test error")
)

type FakeStorage struct{}

func (fs FakeStorage) Add(word, translate string) error {
	return errFake
}

func (fs FakeStorage) Delete(word string) error {
	return errFake
}

func (fs FakeStorage) ValidateExist(word string) error {
	return nil
}

func (fs FakeStorage) ValidateUnexist(word string) error {
	return nil
}

func (fs FakeStorage) ReturnInfo() map[string]string {
	newMap := make(map[string]string)
	return newMap
}

func TestNewPair(t *testing.T) {
	tests := []struct {
		name        string
		storage     dictionary.Dictionary
		word        string
		translate   string
		wantStorage dictionary.Dictionary
		wantErr     error
	}{
		{
			name: "success case",
			storage: dictionary.Dictionary{
				"Car": "Машина",
			},
			word:      "bird",
			translate: "птица",
			wantStorage: dictionary.Dictionary{
				"Car":  "Машина",
				"bird": "птица",
			},
			wantErr: nil,
		},

		{
			name: "word already exist",
			storage: dictionary.Dictionary{
				"Car": "Машина",
			},
			word:      "Car",
			translate: "Машина",
			wantStorage: dictionary.Dictionary{
				"Car": "Машина",
			},
			wantErr: dictionary.ErrWordExist,
		},

		{
			name: "word already exist",
			storage: dictionary.Dictionary{
				"Car": "Машина",
			},
			word:      "Car",
			translate: "Машина",
			wantStorage: dictionary.Dictionary{
				"Car": "Машина",
			},
			wantErr: dictionary.ErrWordExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := New(tt.storage)
			err := service.NewPair(tt.word, tt.translate)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if !maps.Equal(tt.storage, tt.wantStorage) {
				t.Errorf("expected dictionary %v, got %v", tt.wantStorage, tt.storage)
			}
		})
	}
}

func TestDeletePair(t *testing.T) {
	tests := []struct {
		name        string
		storage     dictionary.Dictionary
		word        string
		wantStorage dictionary.Dictionary
		wantErr     error
	}{
		{
			name: "success case",
			storage: dictionary.Dictionary{
				"Car":  "Машина",
				"bird": "птица",
			},
			word: "bird",
			wantStorage: dictionary.Dictionary{
				"Car": "Машина",
			},
			wantErr: nil,
		},

		{
			name: "word unexist",
			storage: dictionary.Dictionary{
				"Car":  "Машина",
				"bird": "птица",
			},
			word: "Dinamo",
			wantStorage: dictionary.Dictionary{
				"Car":  "Машина",
				"bird": "птица",
			},
			wantErr: dictionary.ErrWordUnexist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := New(tt.storage)
			err := service.DeletePair(tt.word)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if !maps.Equal(tt.storage, tt.wantStorage) {
				t.Errorf("expected dictionary %v, got %v", tt.wantStorage, tt.storage)
			}
		})
	}
}

func TestNewPairFakeStorage(t *testing.T) {
	storage := FakeStorage{}
	word := "Dinamo"
	translate := "Fake"
	wantErr := errFake

	service := New(storage)
	err := service.NewPair(word, translate)

	if !errors.Is(err, wantErr) {
		t.Fatalf("expected error %v, got %v", wantErr, err)
	}
}

func TestDeletePairFakeStorage(t *testing.T) {
	storage := FakeStorage{}
	word := "Dinamo"
	wantErr := errFake

	service := New(storage)
	err := service.DeletePair(word)

	if !errors.Is(err, wantErr) {
		t.Fatalf("expected error %v, got %v", wantErr, err)
	}
}
