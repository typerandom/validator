package parser_test

import (
	"fmt"
	. "github.com/typerandom/validator/core/parser"
	"testing"
)

func testThatValidSyntaxIsParsedAsExpected(t *testing.T, test string, expected string) {
	methodGroups, err := Parse(test)

	if err != nil {
		t.Fatalf("Tested '%s'. Didn't expect error, but got %s.", test, err)
	}

	if expected != fmt.Sprint(methodGroups) {
		t.Fatalf("Tested '%s'. Expected '%s' but got '%s'.", test, expected, fmt.Sprint(methodGroups))
	}
}

func TestThatWhenParsingEmptyStringItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "", "[]")
}

func TestThatWhenParsingSingleMethodWithoutParentesesItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc", "[{ name: 'abc', args: (none) }]")
}

func TestThatWhenParsingSingleMethodWithParentesesButNoArgumentsItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc()", "[{ name: 'abc', args: (none) }]")
}

func TestThatWhenParsingSingleMethodWithSingleIntArgumentItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc(1)", "[{ name: 'abc', args: 1 }]")
}

func TestThatWhenParsingSingleMethodWithSingleFloatArgumentItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc(1.234)", "[{ name: 'abc', args: 1.234 }]")
}

func TestThatWhenParsingSingleMethodWithSingleUnboundedStrArgumentItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc(def)", "[{ name: 'abc', args: 'def' }]")
}

func TestThatWhenParsingSingleMethodWithSingleBoundedStrArgumentItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc(´def´)", "[{ name: 'abc', args: 'def' }]")
}

func TestThatWhenParsingSingleMethodWithSingleBoolArgumentItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc(true)", "[{ name: 'abc', args: true }]")
	testThatValidSyntaxIsParsedAsExpected(t, "abc(false)", "[{ name: 'abc', args: false }]")
}

func TestThatWhenParsingSingleMethodWithSingleNilArgumentItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc(nil)", "[{ name: 'abc', args: <nil> }]")
}

func TestThatWhenParsingSingleMethodWithMultipleArgumentsItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc(1, 1.1, def, ´ghi´, true, false, nil)", "[{ name: 'abc', args: 1, 1.1, 'def', 'ghi', true, false, <nil> }]")
}

func TestThatWhenParsingSingleMethodWithMultipleArgumentsWithoutWhiteSpaceItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc(1,1.1,def,´ghi´,true,false,nil)", "[{ name: 'abc', args: 1, 1.1, 'def', 'ghi', true, false, <nil> }]")
}

func TestThatWhenParsingSingleMethodWithMultipleArgumentsWithExessiveWhiteSpaceItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc(    		1,   		1.1,		def, ´ghi´,		true,		false,	   		nil 	   )", "[{ name: 'abc', args: 1, 1.1, 'def', 'ghi', true, false, <nil> }]")
}

func TestThatWhenParsingMultipleMethodsInSingleGroupItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc,def", "[{ name: 'abc', args: (none) }, { name: 'def', args: (none) }]")
}

func TestThatWhenParsingMultipleMethodsWithArgumentsInSingleGroupItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc(123),def(456)", "[{ name: 'abc', args: 123 }, { name: 'def', args: 456 }]")
}

func TestThatWhenParsingMethodsInMultipleGroupsItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc|def", "[{ name: 'abc', args: (none) } { name: 'def', args: (none) }]")
}

func TestThatWhenParsingMultipleMethodsInMultipleGroupsItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc,def()|jkl(123),ghi|xyz", "[{ name: 'abc', args: (none) }, { name: 'def', args: (none) } { name: 'jkl', args: 123 }, { name: 'ghi', args: (none) } { name: 'xyz', args: (none) }]")
}

func TestThatWhenPasingMethodWithBoundedTextArgItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "test(´abc´)", "[{ name: 'test', args: 'abc' }]")
}

func TestThatWhenPasingMethodWithBoundedTextArgAndEscapedValueItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "test(´te\\´st´)", "[{ name: 'test', args: 'te´st' }]")
	testThatValidSyntaxIsParsedAsExpected(t, "test(´te\\st´)", "[{ name: 'test', args: 'test' }]")
	testThatValidSyntaxIsParsedAsExpected(t, "test(´te\\\\st´)", "[{ name: 'test', args: 'te\\st' }]")
	testThatValidSyntaxIsParsedAsExpected(t, "test(´test\\´´)", "[{ name: 'test', args: 'test´' }]")
}

func TestThatWhenParsingValidMethodNameItSucceeds(t *testing.T) {
	testThatValidSyntaxIsParsedAsExpected(t, "abc", "[{ name: 'abc', args: (none) }]")
	testThatValidSyntaxIsParsedAsExpected(t, "Abc", "[{ name: 'Abc', args: (none) }]")
	testThatValidSyntaxIsParsedAsExpected(t, "abc123", "[{ name: 'abc123', args: (none) }]")
	testThatValidSyntaxIsParsedAsExpected(t, "a1B2c3", "[{ name: 'a1B2c3', args: (none) }]")
	testThatValidSyntaxIsParsedAsExpected(t, "test_123", "[{ name: 'test_123', args: (none) }]")
}

func testThatInvalidSyntaxFailsWithError(t *testing.T, test string, expectedErr string) {
	methodGroups, err := Parse(test)

	if len(methodGroups) != 0 {
		t.Fatalf("Tested '%s'. Didn't expect any method groups, but got %d.", test, len(methodGroups))
	}

	if err == nil {
		t.Fatalf("Tested '%s'. Expected error, but didn't get any.", test)
	}

	if err.Error() != expectedErr {
		t.Fatalf("Tested '%s'. Expected '%s' error, but got '%s'.", test, expectedErr, err)
	}
}

func TestThatWhenParsingInvalidMethodNameItFails(t *testing.T) {
	testThatInvalidSyntaxFailsWithError(t, ",", "Unexpected character U+002C ',' at position 1.")
	testThatInvalidSyntaxFailsWithError(t, ".", "Unexpected character U+002E '.' at position 1.")
	testThatInvalidSyntaxFailsWithError(t, "(", "Unexpected character U+0028 '(' at position 1.")
	testThatInvalidSyntaxFailsWithError(t, ")", "Unexpected character U+0029 ')' at position 1.")
	testThatInvalidSyntaxFailsWithError(t, "|", "Unexpected character U+007C '|' at position 1.")
	testThatInvalidSyntaxFailsWithError(t, "´", "Unexpected character U+00B4 '´' at position 2.")
	testThatInvalidSyntaxFailsWithError(t, "1", "Unexpected character U+0031 '1' at position 1.")
	testThatInvalidSyntaxFailsWithError(t, "_Test()", "Unexpected character U+005F '_' at position 1.")
}

func TestThatWhenParsingMethodNamesWithInvalidSeparatorsItFails(t *testing.T) {
	testThatInvalidSyntaxFailsWithError(t, ",test", "Unexpected character U+002C ',' at position 1.")
	testThatInvalidSyntaxFailsWithError(t, "test,", "Unexpected character U+002C ',' at position 5.")
	testThatInvalidSyntaxFailsWithError(t, ",test,", "Unexpected character U+002C ',' at position 1.")
	testThatInvalidSyntaxFailsWithError(t, "test(),", "Unexpected character U+002C ',' at position 7.")
	testThatInvalidSyntaxFailsWithError(t, ",test()", "Unexpected character U+002C ',' at position 1.")
	testThatInvalidSyntaxFailsWithError(t, ",test(),", "Unexpected character U+002C ',' at position 1.")
}

func TestThatWhenParsingMethodArgsWithInvalidSeparatorsItFails(t *testing.T) {
	testThatInvalidSyntaxFailsWithError(t, "test(,)", "Unexpected character U+002C ',' at position 6.")
	testThatInvalidSyntaxFailsWithError(t, "test(1,)", "Unexpected character U+0029 ')' at position 8.")
	testThatInvalidSyntaxFailsWithError(t, "test(,1)", "Unexpected character U+002C ',' at position 6.")
	testThatInvalidSyntaxFailsWithError(t, "test(,,)", "Unexpected character U+002C ',' at position 6.")
}

func TestThatWhenParsingInvalidGroupSeparatorsItFails(t *testing.T) {
	testThatInvalidSyntaxFailsWithError(t, "|", "Unexpected character U+007C '|' at position 1.")
	testThatInvalidSyntaxFailsWithError(t, "a|", "Unexpected character U+007C '|' at position 2.")
	testThatInvalidSyntaxFailsWithError(t, "|a", "Unexpected character U+007C '|' at position 1.")
	testThatInvalidSyntaxFailsWithError(t, "|a|", "Unexpected character U+007C '|' at position 1.")
	testThatInvalidSyntaxFailsWithError(t, "||", "Unexpected character U+007C '|' at position 1.")
}
