package dictionary

import (
	"errors"
	"maps"
	"testing"
)

func TestValidateExist(t *testing.T) {
	tests := []struct {
		name           string
		dictionary     Dictionary
		word           string
		wantDictionary Dictionary
		wantErr        error
	}{
		{
			name: "success case",
			dictionary: Dictionary{
				"Bird": "Птица",
				"Car":  "Машина",
			},
			word: "Home",
			wantDictionary: Dictionary{
				"Bird": "Птица",
				"Car":  "Машина",
			},
			wantErr: nil,
		},

		{
			name: "error case",
			dictionary: Dictionary{
				"Bird": "Птица",
				"Car":  "Машина",
			},
			word: "Car",
			wantDictionary: Dictionary{
				"Bird": "Птица",
				"Car":  "Машина",
			},
			wantErr: ErrWordExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dictionary.ValidateExist(tt.word)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if !maps.Equal(tt.dictionary, tt.wantDictionary) {
				t.Errorf("expected dictionary %v, got %v", tt.wantDictionary, tt.dictionary)
			}
		})
	}
}

func TestValidateUnexist(t *testing.T) {
	tests := []struct {
		name           string
		dictionary     Dictionary
		word           string
		wantDictionary Dictionary
		wantErr        error
	}{
		{
			name: "success case",
			dictionary: Dictionary{
				"Bird": "Птица",
				"Car":  "Машина",
			},
			word: "Car",
			wantDictionary: Dictionary{
				"Bird": "Птица",
				"Car":  "Машина",
			},
			wantErr: nil,
		},

		{
			name: "error case",
			dictionary: Dictionary{
				"Bird": "Птица",
				"Car":  "Машина",
			},
			word: "Home",
			wantDictionary: Dictionary{
				"Bird": "Птица",
				"Car":  "Машина",
			},
			wantErr: ErrWordUnexist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dictionary.ValidateUnexist(tt.word)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if !maps.Equal(tt.dictionary, tt.wantDictionary) {
				t.Errorf("expected dictionary %v, got %v", tt.wantDictionary, tt.dictionary)
			}
		})
	}
}

func TestValidateWordsAmount(t *testing.T) {
	tests := []struct {
		name    string
		amount  int
		d       Dictionary
		wantD   Dictionary
		wantErr error
	}{
		{
			name:   "success case",
			amount: 2,
			d: Dictionary{
				"Word": "Слово",
				"Lion": "Лев",
			},
			wantD: Dictionary{
				"Word": "Слово",
				"Lion": "Лев",
			},
			wantErr: nil,
		},
		{
			name:   "error case",
			amount: 3,
			d: Dictionary{
				"Word": "Слово",
				"Lion": "Лев",
			},
			wantD: Dictionary{
				"Word": "Слово",
				"Lion": "Лев",
			},
			wantErr: ErrWordsNotEnough,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.d.ValidateWordsAmount(tt.amount)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if !maps.Equal(tt.d, tt.wantD) {
				t.Errorf("expected dictionary %v, got %v", tt.wantD, tt.d)
			}
		})
	}
}
