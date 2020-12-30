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

func Test_ParseGitUri(t *testing.T) {
	type args struct {
		gituri string
	}
	tests := []struct {
		name      string
		arguments args
		expected  URLParts
	}{
		{
			name:      "First",
			arguments: args{gituri: "git@192.168.0.110:/home/git/repos/gitea-install.git"},
			expected:  URLParts{Schema: "ssh", User: "git", Host: "192.168.0.110", Path: "/home/git/repos/gitea-install.git"},
		},
		{
			name:      "Second",
			arguments: args{gituri: "git@github.com:khmarbaise/githelper.git"},
			expected:  URLParts{Schema: "ssh", User: "git", Host: "github.com", Path: "khmarbaise/githelper.git"},
		},
		{
			name:      "Third",
			arguments: args{gituri: "ssh://git@github.com/khmarbaise/githelper.git"},
			expected:  URLParts{Schema: "ssh", User: "git", Host: "github.com", Path: "/khmarbaise/githelper.git"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result, err := ParseGitURI(tt.arguments.gituri); err == nil && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseGitUrl(%v) current:'%v' result:'%v', expected: '%v'", tt.name, tt.arguments.gituri, result, tt.expected)
			}
		})
	}
}
