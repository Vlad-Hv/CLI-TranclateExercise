package service

import (
	"errors"
	"exerciser/internal/storage/dictionary"
	results "exerciser/internal/storage/results"
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

func (fs FakeStorage) ValidateWordsAmount(amount int) error {
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
		results     results.ResultStorage
		word        string
		translate   string
		wantStorage dictionary.Dictionary
		wantResults results.ResultStorage
		wantErr     error
	}{
		{
			name: "success case",
			storage: dictionary.Dictionary{
				"Car": "Машина",
			},
			results: []map[string]string{
				{
					"bird": "correct",
				},
			},
			word:      "bird",
			translate: "птица",
			wantStorage: dictionary.Dictionary{
				"Car":  "Машина",
				"bird": "птица",
			},
			wantResults: []map[string]string{
				{
					"bird": "correct",
				},
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
			service := NewWordStorage(tt.storage, &tt.results)
			err := service.NewPair(tt.word, tt.translate)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if !maps.Equal(tt.storage, tt.wantStorage) {
				t.Errorf("expected dictionary %v, got %v", tt.wantStorage, tt.storage)
			}

			if len(tt.results) != len(tt.wantResults) {
				t.Fatalf("expected lenth %v, got %v", len(tt.wantStorage), len(tt.storage))
			}

			for i := range tt.wantResults {
				if !maps.Equal(tt.results[i], tt.wantResults[i]) {
					t.Errorf("expected result %v, got %v", tt.results[i], tt.wantResults[i])
				}
			}
		})
	}
}

func TestDeletePair(t *testing.T) {
	tests := []struct {
		name        string
		storage     dictionary.Dictionary
		word        string
		results     results.ResultStorage
		wantStorage dictionary.Dictionary
		wantResults results.ResultStorage
		wantErr     error
	}{
		{
			name: "success case",
			storage: dictionary.Dictionary{
				"Car":  "Машина",
				"bird": "птица",
			},
			results: []map[string]string{
				{
					"bird": "correct",
				},
			},
			word: "bird",
			wantStorage: dictionary.Dictionary{
				"Car": "Машина",
			},
			wantResults: []map[string]string{
				{
					"bird": "correct",
				},
			},
			wantErr: nil,
		},

		{
			name: "word unexist",
			storage: dictionary.Dictionary{
				"Car":  "Машина",
				"bird": "птица",
			},
			results: []map[string]string{
				{
					"bird": "correct",
				},
			},
			word: "Dinamo",
			wantStorage: dictionary.Dictionary{
				"Car":  "Машина",
				"bird": "птица",
			},
			wantResults: []map[string]string{
				{
					"bird": "correct",
				},
			},
			wantErr: dictionary.ErrWordUnexist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewWordStorage(tt.storage, &tt.results)
			err := service.DeletePair(tt.word)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if len(tt.storage) != len(tt.wantStorage) {
				t.Fatalf("expected lenth %v, got %v", len(tt.wantStorage), len(tt.storage))
			}

			if len(tt.results) != len(tt.wantResults) {
				t.Fatalf("expected lenth %v, got %v", len(tt.wantResults), len(tt.results))
			}
			if !maps.Equal(tt.storage, tt.wantStorage) {
				t.Errorf("expected dictionary %v, got %v", tt.wantStorage, tt.storage)
			}

			for i := range tt.wantResults {
				if !maps.Equal(tt.results[i], tt.wantResults[i]) {
					t.Errorf("expected result %v, got %v", tt.results[i], tt.wantResults[i])
				}
			}
		})
	}
}

func TestNewPairFakeStorage(t *testing.T) {
	storage := FakeStorage{}
	word := "Dinamo"
	translate := "Fake"
	oneResult := results.ResultStorage{
		{
			"bird": "correct",
		},
	}
	wantResults := results.ResultStorage{
		{
			"bird": "correct",
		},
	}
	wantErr := errFake

	service := NewWordStorage(storage, &oneResult)
	err := service.NewPair(word, translate)

	if !errors.Is(err, wantErr) {
		t.Fatalf("expected error %v, got %v", wantErr, err)
	}

	if len(oneResult) != len(wantResults) {
		t.Fatalf("expected index %v, got %v", len(wantResults), len(oneResult))
	}

	for i := range wantResults {
		if !maps.Equal(oneResult[i], wantResults[i]) {
			t.Errorf("expected result %v, got %v", oneResult[i], wantResults[i])
		}
	}

}

func TestDeletePairFakeStorage(t *testing.T) {
	storage := FakeStorage{}
	word := "Dinamo"
	wantErr := errFake
	oneResult := results.ResultStorage{
		{
			"bird": "correct",
		},
	}
	wantResults := results.ResultStorage{
		{
			"bird": "correct",
		},
	}
	service := NewWordStorage(storage, &oneResult)
	err := service.DeletePair(word)

	if !errors.Is(err, wantErr) {
		t.Fatalf("expected error %v, got %v", wantErr, err)
	}

	if len(oneResult) != len(wantResults) {
		t.Fatalf("expected index %v, got %v", len(wantResults), len(oneResult))
	}
	for i := range wantResults {
		if !maps.Equal(oneResult[i], wantResults[i]) {
			t.Errorf("expected result %v, got %v", oneResult[i], wantResults[i])
		}
	}
}
