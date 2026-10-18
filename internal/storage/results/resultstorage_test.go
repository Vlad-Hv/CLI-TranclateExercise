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

	if !errors.Is(err, wantErr) {
		t.Errorf("expected error %v, got %v", wantErr, err)
	}

	for i := range wantRs {
		if !maps.Equal(rs[i], wantRs[i]) {
			t.Errorf("expected result %v, got %v", rs[i], wantRs[i])
		}
	}
}
