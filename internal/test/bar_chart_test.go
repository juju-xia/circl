package test

import "testing"

func TestBarChartFixtureCounts(t *testing.T) {
	changes := []struct {
		name     string
		files    int
		added    int
		removed  int
		expected int
	}{
		{name: "documentation", files: 3, added: 24, removed: 3, expected: 21},
		{name: "go-tests", files: 2, added: 18, removed: 2, expected: 16},
		{name: "fixtures", files: 5, added: 12, removed: 0, expected: 12},
	}

	for _, change := range changes {
		got := change.added - change.removed
		if got != change.expected {
			t.Fatalf("%s: got net change %d, want %d", change.name, got, change.expected)
		}
	}

	totalFiles, totalAdded, totalRemoved := 0, 0, 0
	for _, change := range changes {
		totalFiles += change.files
		totalAdded += change.added
		totalRemoved += change.removed
	}

	if totalFiles != 10 || totalAdded != 54 || totalRemoved != 5 {
		t.Fatalf(
			"totals: got files=%d added=%d removed=%d, want files=10 added=54 removed=5",
			totalFiles,
			totalAdded,
			totalRemoved,
		)
	}
}
