package service

import (
	"errors"
	"testing"
)

func TestValidatePairSuccessCase(t *testing.T) {
	word := "Hello"
	translate := "Привет"
	var wantErr error = nil

	if err := validatePair(word, translate); !errors.Is(err, wantErr) {
		t.Errorf("expected error %v, got %v", wantErr, err)
	}
}

func TestValidatePairErrorCases(t *testing.T) {
	tests := []struct {
		name      string
		word      string
		translate string
		wantErr   error
	}{
		{
			name:      "both empty",
			word:      " ",
			translate: "",
			wantErr:   errWordEmpty,
		},
		{
			name:      "word empty",
			word:      "",
			translate: "тест",
			wantErr:   errWordEmpty,
		},
		{
			name:      "translate empty",
			word:      "Test",
			translate: "",
			wantErr:   errTranslateEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePair(tt.word, tt.translate); !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestValidateAmountSuccessCase(t *testing.T) {
	amount := 1
	var wantErr error = nil

	if err := ValidateAmount(amount); !errors.Is(err, wantErr) {
		t.Errorf("expected error %v, got %v", wantErr, err)
	}
}

func TestValidateAmountErrorCase(t *testing.T) {
	amount := -1
	wantErr := errAmountInvalid

	if err := ValidateAmount(amount); !errors.Is(err, wantErr) {
		t.Errorf("expected error %v, got %v", wantErr, err)
	}
}

func TestValidateAmountEdgeCase(t *testing.T) {
	amount := 0
	wantErr := errAmountInvalid

	if err := ValidateAmount(amount); !errors.Is(err, wantErr) {
		t.Errorf("expected error %v, got %v", wantErr, err)
	}
}
