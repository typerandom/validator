package parser

import (
	"fmt"
	"testing"
)

func TestThatValidSyntaxIsParsedAsExpected(t *testing.T) {
	tests := map[string]string{
		``:                                                           `[]`,
		`min`:                                                        `[{ name: 'min', args: (none) }]`,
		`min()`:                                                      `[{ name: 'min', args: (none) }]`,
		`min(1)`:                                                     `[{ name: 'min', args: '1' }]`,
		`min(1, 2)`:                                                  `[{ name: 'min', args: '1', '2' }]`,
		`min(1, 2,3)`:                                                `[{ name: 'min', args: '1', '2', '3' }]`,
		`min,min(1)`:                                                 `[{ name: 'min', args: (none) }, { name: 'min', args: '1' }]`,
		`min|min(1)`:                                                 `[{ name: 'min', args: (none) } { name: 'min', args: '1' }]`,
		`min(1),min(1,2),min|min|min(1),min(3, 4),min`:               `[{ name: 'min', args: '1' }, { name: 'min', args: '1', '2' }, { name: 'min', args: (none) } { name: 'min', args: (none) } { name: 'min', args: '1' }, { name: 'min', args: '3', '4' }, { name: 'min', args: (none) }]`,
		`m(123)|m(1.23)|m()|m(1),m(1, 2),m(1,2,3),m(1,´test123´, 3)`: `[{ name: 'm', args: '123' } { name: 'm', args: '1.23' } { name: 'm', args: (none) } { name: 'm', args: '1' }, { name: 'm', args: '1', '2' }, { name: 'm', args: '1', '2', '3' }, { name: 'm', args: '1', 'test123', '3' }]`,
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
		`min,`,
		`min(`,
		`min)`,
		`min|`,
		`min()|`,
		`|min()`,
		`min(),`,
		`,min()`,
		`min(,)`,
		`min(123,)`,
		`min(´test)`,
		`min(test´)`,
		`,`,
		`|`,
		`)`,
		`(`,
		`´`,
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
