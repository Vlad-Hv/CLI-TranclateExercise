package service

import (
	"errors"
	"testing"
)

var (
	errTest = errors.New("tests")
)

type FakeResult []map[string]string

func (fs FakeResult) Add(result map[string]string) error {
	return errTest
}

func (fr FakeResult) ReturnResults() []map[string]string {
	return fr
}

type FakeSuccessResult []map[string]string

func (fs FakeSuccessResult) Add(result map[string]string) error {
	return nil
}

func (fr FakeSuccessResult) ReturnResults() []map[string]string {
	return fr
}

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

func TestCheckAnswer(t *testing.T) {
	tests := []struct {
		name   string
		t      string
		ur     string
		wantOk bool
	}{
		{
			name:   "success case",
			t:      "cAr",
			ur:     "car",
			wantOk: true,
		},

		{
			name:   "error case",
			t:      "Cat ",
			ur:     "car",
			wantOk: false,
		},

		{
			name:   "edge case",
			t:      "    ХромОй гипопоТам  ",
			ur:     "  хрОмОй гИпопоТам  ",
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := CheckAnswer(tt.t, tt.ur)
			if ok != tt.wantOk {
				t.Errorf("Expected %v, got %v", tt.wantOk, ok)
			}
		})
	}
}

func TestSaveResultFake(t *testing.T) {
	tests := []struct {
		name    string
		result  map[string]string
		s       Service
		wantErr error
	}{
		{
			name: "error case",
			s: Service{
				Results: FakeResult{},
			},
			result: map[string]string{
				"car": "correct",
				"cat": "correct",
			},

			wantErr: errTest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.SaveResult(tt.result)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestSaveResult(t *testing.T) {
	tests := []struct {
		name    string
		result  map[string]string
		s       Service
		wantErr error
	}{
		{
			name: "error case",
			s: Service{
				Results: FakeSuccessResult{},
			},
			result: map[string]string{
				"car": "correct",
				"cat": "correct",
			},

			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.SaveResult(tt.result)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}
