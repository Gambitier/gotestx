package gotestx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Example 1: Simple string to int function
func TestStringLength(t *testing.T) {
	tests := []TableTestCase[string, int]{
		{
			Name:  "empty string",
			Input: "",
			Assert: func(t *testing.T, input string, actual int) {
				assert.Equal(t, 0, actual, "empty string should have length 0")
			},
		},
		{
			Name:  "single character",
			Input: "a",
			Assert: func(t *testing.T, input string, actual int) {
				assert.Equal(t, 1, actual, "single character should have length 1")
			},
		},
		{
			Name:  "multiple characters",
			Input: "hello world",
			Assert: func(t *testing.T, input string, actual int) {
				assert.Equal(t, 11, actual, "should count all characters including space")
			},
		},
	}

	RunTableTests(t, tests, func(input string) int {
		return len(input)
	})
}

// Example 2: String to bool function
func TestIsPalindrome(t *testing.T) {
	tests := []TableTestCase[string, bool]{
		{
			Name:  "empty string is palindrome",
			Input: "",
			Assert: func(t *testing.T, input string, actual bool) {
				assert.True(t, actual, "empty string should be considered palindrome")
			},
		},
		{
			Name:  "single character is palindrome",
			Input: "a",
			Assert: func(t *testing.T, input string, actual bool) {
				assert.True(t, actual, "single character should be palindrome")
			},
		},
		{
			Name:  "palindrome word",
			Input: "racecar",
			Assert: func(t *testing.T, input string, actual bool) {
				assert.True(t, actual, "racecar should be palindrome")
			},
		},
		{
			Name:  "non-palindrome word",
			Input: "hello",
			Assert: func(t *testing.T, input string, actual bool) {
				assert.False(t, actual, "hello should not be palindrome")
			},
		},
	}

	isPalindrome := func(input string) bool {
		if len(input) <= 1 {
			return true
		}
		for i := 0; i < len(input)/2; i++ {
			if input[i] != input[len(input)-1-i] {
				return false
			}
		}
		return true
	}

	RunTableTests(t, tests, isPalindrome)
}

func TestPersonToTags(t *testing.T) {
	// Example 3: Struct input and slice output
	type Person struct {
		Name string
		Age  int
	}

	tests := []TableTestCase[Person, []string]{
		{
			Name:  "young person",
			Input: Person{Name: "Alice", Age: 20},
			Assert: func(t *testing.T, input Person, actual []string) {
				assert.Contains(t, actual, "young")
				assert.Contains(t, actual, "Alice")
				assert.Len(t, actual, 2)
			},
		},
		{
			Name:  "adult person",
			Input: Person{Name: "Bob", Age: 35},
			Assert: func(t *testing.T, input Person, actual []string) {
				assert.Contains(t, actual, "adult")
				assert.Contains(t, actual, "Bob")
				assert.Len(t, actual, 2)
			},
		},
		{
			Name:  "senior person",
			Input: Person{Name: "Charlie", Age: 70},
			Assert: func(t *testing.T, input Person, actual []string) {
				assert.Contains(t, actual, "senior")
				assert.Contains(t, actual, "Charlie")
				assert.Len(t, actual, 2)
			},
		},
	}

	personToTags := func(p Person) []string {
		tags := []string{p.Name}
		switch {
		case p.Age < 25:
			tags = append(tags, "young")
		case p.Age < 65:
			tags = append(tags, "adult")
		default:
			tags = append(tags, "senior")
		}
		return tags
	}

	RunTableTests(t, tests, personToTags)
}

func TestStringContains(t *testing.T) {
	// Example 4: Multiple input parameters using struct
	type StringOperation struct {
		Text          string
		Pattern       string
		CaseSensitive bool
	}
	tests := []TableTestCase[StringOperation, bool]{
		{
			Name:  "case sensitive match",
			Input: StringOperation{Text: "Hello World", Pattern: "World", CaseSensitive: true},
			Assert: func(t *testing.T, input StringOperation, actual bool) {
				assert.True(t, actual, "should find 'World' in 'Hello World'")
			},
		},
		{
			Name:  "case sensitive no match",
			Input: StringOperation{Text: "Hello World", Pattern: "world", CaseSensitive: true},
			Assert: func(t *testing.T, input StringOperation, actual bool) {
				assert.False(t, actual, "should not find 'world' in 'Hello World' (case sensitive)")
			},
		},
		{
			Name:  "case insensitive match",
			Input: StringOperation{Text: "Hello World", Pattern: "world", CaseSensitive: false},
			Assert: func(t *testing.T, input StringOperation, actual bool) {
				assert.True(t, actual, "should find 'world' in 'Hello World' (case insensitive)")
			},
		},
	}

	RunTableTests(t, tests, func(input StringOperation) bool {
		if input.CaseSensitive {
			return strings.Contains(input.Text, input.Pattern)
		}
		return strings.Contains(strings.ToLower(input.Text), strings.ToLower(input.Pattern))
	})
}

// Example 5: Complex assertion with multiple checks
func TestStringProcessing(t *testing.T) {
	tests := []TableTestCase[string, map[string]interface{}]{
		{
			Name:  "mixed case string",
			Input: "Hello World 123!",
			Assert: func(t *testing.T, input string, actual map[string]interface{}) {
				assert.Equal(t, 15, actual["length"])
				assert.Equal(t, 2, actual["words"])
				assert.Equal(t, 2, actual["uppercase"])
				assert.Equal(t, 8, actual["lowercase"])
				assert.Equal(t, 3, actual["digits"])
				assert.Equal(t, 1, actual["special"])
			},
		},
		{
			Name:  "empty string",
			Input: "",
			Assert: func(t *testing.T, input string, actual map[string]interface{}) {
				assert.Equal(t, 0, actual["length"])
				assert.Equal(t, 0, actual["words"])
				assert.Equal(t, 0, actual["uppercase"])
				assert.Equal(t, 0, actual["lowercase"])
				assert.Equal(t, 0, actual["digits"])
				assert.Equal(t, 0, actual["special"])
			},
		},
	}

	RunTableTests(t, tests, func(input string) map[string]interface{} {
		result := map[string]interface{}{
			"length":    len(input),
			"words":     0,
			"uppercase": 0,
			"lowercase": 0,
			"digits":    0,
			"special":   0,
		}

		if len(input) == 0 {
			return result
		}

		words := strings.Fields(input)
		result["words"] = len(words)

		for _, char := range input {
			switch {
			case char >= 'A' && char <= 'Z':
				result["uppercase"] = result["uppercase"].(int) + 1
			case char >= 'a' && char <= 'z':
				result["lowercase"] = result["lowercase"].(int) + 1
			case char >= '0' && char <= '9':
				result["digits"] = result["digits"].(int) + 1
			default:
				if char != ' ' {
					result["special"] = result["special"].(int) + 1
				}
			}
		}

		return result
	})
}

// Example 6: Using custom assertion helpers
func TestWithCustomAssertions(t *testing.T) {
	// Custom assertion helper
	assertEven := func(t *testing.T, input int, actual bool) {
		if input%2 == 0 {
			assert.True(t, actual, "even number %d should return true", input)
		} else {
			assert.False(t, actual, "odd number %d should return false", input)
		}
	}

	tests := []TableTestCase[int, bool]{
		{Name: "zero", Input: 0, Assert: assertEven},
		{Name: "one", Input: 1, Assert: assertEven},
		{Name: "two", Input: 2, Assert: assertEven},
		{Name: "three", Input: 3, Assert: assertEven},
		{Name: "four", Input: 4, Assert: assertEven},
	}

	RunTableTests(t, tests, func(input int) bool {
		return input%2 == 0
	})
}

func TestLeapYear(t *testing.T) {
	type input struct {
		Year int
	}

	tests := []TableTestCase[input, bool]{
		{
			Name:  "leap year",
			Input: input{2000},
			Assert: func(t *testing.T, ip input, op bool) {
				assert.True(t, op)
			},
		},
		{
			Name:  "not leap year",
			Input: input{2003},
			Assert: func(t *testing.T, ip input, op bool) {
				assert.False(t, op)
			},
		},
		{
			Input: input{2004},
			Assert: func(t *testing.T, ip input, op bool) {
				assert.False(t, op)
			},
		},
	}

	leapYear := func(input input) bool {
		return input.Year%4 == 0
	}

	RunTableTests(t, tests, leapYear)
}
