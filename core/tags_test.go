package core_test

import (
	"fmt"
	. "github.com/typerandom/validator/core"
	"testing"
)

func TestThatValidTagsAreParsedAsExpected(t *testing.T) {
	tags := map[string]string{
		``:                                             `[]`,
		`min`:                                          `[{ name: 'min', options: (none) }]`,
		`min()`:                                        `[{ name: 'min', options: (none) }]`,
		`min(1)`:                                       `[{ name: 'min', options: '1' }]`,
		`min(1, 2)`:                                    `[{ name: 'min', options: '1', '2' }]`,
		`min(1, 2,3)`:                                  `[{ name: 'min', options: '1', '2', '3' }]`,
		`min,min(1)`:                                   `[{ name: 'min', options: (none) }, { name: 'min', options: '1' }]`,
		`min|min(1)`:                                   `[{ name: 'min', options: (none) } { name: 'min', options: '1' }]`,
		`min(1),min(1,2),min|min|min(1),min(3, 4),min`: `[{ name: 'min', options: '1' }, { name: 'min', options: '1', '2' }, { name: 'min', options: (none) } { name: 'min', options: (none) } { name: 'min', options: '1' }, { name: 'min', options: '3', '4' }, { name: 'min', options: (none) }]`,
	}

	for tag, expected := range tags {
		tagGroups, err := ParseTag(tag)

		if err != nil {
			t.Fatalf("Didn't expect error, but got %s.", err)
		}

		if expected != fmt.Sprint(tagGroups) {
			t.Fatalf("Expected '%s' but got '%s'.", expected, fmt.Sprint(tagGroups))
		}
	}
}
