package modules

import (
	"reflect"
	"testing"
)

func TestExtractString(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name      string
		arguments args
		expected  string
	}{
		{name: "First Summary", arguments: args{lines: []string{"Title: ", "Test", "summary: TestSummary"}}, expected: "TestSummary"},
		{name: "No Summary", arguments: args{lines: []string{"Title: ", "Test", "Empty lines"}}, expected: ""},
		{name: "Summary in between", arguments: args{lines: []string{"Title: ", "summary: This is the summary.", "Empty lines"}}, expected: "This is the summary."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := ExtractSummary(tt.arguments.lines); !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ExtractSummary(%v) current:'%v' result:'%v', expected: '%v'", tt.name, tt.arguments.lines, result, tt.expected)
			}

		})
	}
}
