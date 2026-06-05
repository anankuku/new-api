package model

import (
	"reflect"
	"testing"
)

func TestParseTokenGroupPriority(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want []string
	}{
		{
			name: "json array",
			raw:  `["default","gemini-cli","default"," backup "]`,
			want: []string{"default", "gemini-cli", "backup"},
		},
		{
			name: "comma separated fallback",
			raw:  "default,gemini-cli\nbackup",
			want: []string{"default", "gemini-cli", "backup"},
		},
		{
			name: "invalid json fallback to empty",
			raw:  `["default"`,
			want: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseTokenGroupPriority(test.raw)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("ParseTokenGroupPriority() = %#v, want %#v", got, test.want)
			}
		})
	}
}

func TestTokenNormalizeGroupPriority(t *testing.T) {
	token := &Token{Group: "old", GroupPriority: `["default","backup"]`}

	token.NormalizeGroupPriority()

	if token.Group != "default" {
		t.Fatalf("Group = %q, want default", token.Group)
	}
	if token.GroupPriority != `["default","backup"]` {
		t.Fatalf("GroupPriority = %q, want normalized JSON", token.GroupPriority)
	}
}
