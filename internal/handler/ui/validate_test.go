package ui

import (
	"errors"
	"testing"
)

func TestValidateImport(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:    "success case",
			input:   "test",
			wantErr: nil,
		},

		{
			name:    "error case",
			input:   "",
			wantErr: errInputEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateImport(tt.input)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestValidateMenu(t *testing.T) {
	tests := []struct {
		name    string
		option  int
		err     error
		wantErr error
	}{
		{
			name:    "success case",
			option:  1,
			err:     nil,
			wantErr: nil,
		},

		{
			name:    "option is zero",
			option:  0,
			err:     nil,
			wantErr: errOptionInvalid,
		},

		{
			name:    "option invalid type",
			option:  2,
			err:     errors.New("test"),
			wantErr: errTypeInvalid,
		},

		{
			name:    "option less than one case",
			option:  -3,
			err:     nil,
			wantErr: errOptionInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateMenu(tt.option, tt.err)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}
