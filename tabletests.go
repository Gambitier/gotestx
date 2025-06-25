package gotestx

import (
	"encoding/json"
	"testing"
)

type TableTestCase[I any, O any] struct {
	Name   string     `json:"name,omitempty"`
	Input  I          `json:"input"`
	Assert func(I, O) `json:"-"`
}

func RunTableTests[I any, O any](t *testing.T, tests []TableTestCase[I, O], fn func(I) O) {
	for _, tc := range tests {
		var name string
		if tc.Name != "" {
			name = tc.Name
		} else {
			nameBytes, err := json.Marshal(tc)
			if err != nil {
				t.Fatalf("failed to marshal test case: %v", err)
			}
			name = string(nameBytes)
		}

		tc := tc

		t.Run(name, func(t *testing.T) {
			actual := fn(tc.Input)
			tc.Assert(tc.Input, actual)
		})
	}
}
