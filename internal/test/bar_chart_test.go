package test

import "testing"

func TestBarChartFixtureCounts(t *testing.T) {
	changes := []struct {
		name     string
		added    int
		removed  int
		expected int
	}{
		{name: "documentation", added: 24, removed: 3, expected: 21},
		{name: "go-tests", added: 18, removed: 2, expected: 16},
		{name: "fixtures", added: 12, removed: 0, expected: 12},
	}

	for _, change := range changes {
		got := change.added - change.removed
		if got != change.expected {
			t.Fatalf("%s: got net change %d, want %d", change.name, got, change.expected)
		}
	}
}
