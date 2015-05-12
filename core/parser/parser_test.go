package parser

import (
	"fmt"
	"testing"
)

func TestThatValidSyntaxIsParsedAsExpected(t *testing.T) {
	tests := map[string]string{
		``:                                             `[]`,
		`min`:                                          `[{ name: 'min', params: (none) }]`,
		`min()`:                                        `[{ name: 'min', params: (none) }]`,
		`min(1)`:                                       `[{ name: 'min', params: '1' }]`,
		`min(1, 2)`:                                    `[{ name: 'min', params: '1', '2' }]`,
		`min(1, 2,3)`:                                  `[{ name: 'min', params: '1', '2', '3' }]`,
		`min,min(1)`:                                   `[{ name: 'min', params: (none) }, { name: 'min', params: '1' }]`,
		`min|min(1)`:                                   `[{ name: 'min', params: (none) } { name: 'min', params: '1' }]`,
		`min(1),min(1,2),min|min|min(1),min(3, 4),min`: `[{ name: 'min', params: '1' }, { name: 'min', params: '1', '2' }, { name: 'min', params: (none) } { name: 'min', params: (none) } { name: 'min', params: '1' }, { name: 'min', params: '3', '4' }, { name: 'min', params: (none) }]`,
	}

	for test, expected := range tests {
		methodGroups, err := Parse(test)

		if err != nil {
			t.Fatalf("Tag '%s'. Didn't expect error, but got %s.", test, err)
		}

		if expected != fmt.Sprint(methodGroups) {
			t.Fatalf("Tag '%s'. Expected '%s' but got '%s'.", test, expected, fmt.Sprint(methodGroups))
		}
	}
}

func TestThatInvalidSyntaxFailsWithError(t *testing.T) {
	tests := []string{
		`,`,
		`min,`,
		`min(`,
		`min)`,
		`min|`,
		`|`,
		`)`,
		`(`,
	}

	for _, test := range tests {
		methodGroups, err := Parse(test)

		if len(methodGroups) != 0 {
			t.Fatalf("Tag: '%s'. Didn't expect any method groups, but got %d.", test, len(methodGroups))
		}

		if err == nil {
			t.Fatalf("Tag '%s'. Expected error, but didn't get any.", test)
		}
	}
}
