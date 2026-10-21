package result

import (
	"errors"
	"maps"
	"testing"
)

func TestAddSuccess(t *testing.T) {
	rs := ResultStorage{
		{
			"Lion": "correct",
		},
	}

	result := map[string]string{
		"Car": "uncorrect",
	}

	var wantErr error = nil
	wantRs := ResultStorage{
		{
			"Lion": "correct",
		},
		{
			"Car": "uncorrect",
		},
	}
	err := rs.Add(result)

	if len(rs) != len(wantRs) {
		t.Fatalf("expected lenth %v, got %v", len(wantRs), len(rs))
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected error %v, got %v", wantErr, err)
	}

	for i := range wantRs {
		if !maps.Equal(rs[i], wantRs[i]) {
			t.Errorf("expected result %v, got %v", rs[i], wantRs[i])
		}
	}

	result["Car"] = "changed"

	for i := range wantRs {
		if !maps.Equal(rs[i], wantRs[i]) {
			t.Errorf("expected result %v, got %v", rs[i], wantRs[i])
		}
	}
}

func TestReturnResults(t *testing.T) {
	rs := ResultStorage{
		{
			"car": "cat",
		},
	}

	wantRs := ResultStorage{
		{
			"car": "cat",
		},
	}

	result := rs.ReturnResults()

	if len(result) != len(wantRs) {
		t.Fatalf("expected lenth %v, got %v", len(wantRs), len(result))
	}

	for i := range rs {
		if !maps.Equal(result[i], wantRs[i]) {
			t.Errorf("expected map %v, got %v", wantRs[i], result[i])
		}
	}

	rs[0]["car"] = "changed"

	if len(result) != len(wantRs) {
		t.Fatalf("expected lenth %v, got %v", len(wantRs), len(result))
	}

	for i := range rs {
		if !maps.Equal(result[i], wantRs[i]) {
			t.Errorf("expected map %v, got %v", wantRs[i], result[i])
		}
	}

	result[0] = map[string]string{"car": "caf"}
	for i := range result {
		if maps.Equal(result[i], rs[i]) {
			t.Errorf("expected map %v, got %v", wantRs[i], result[i])
		}
	}
}
