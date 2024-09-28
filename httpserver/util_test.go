package httpserver

import (
	"testing"
)

func TestSplitUriQuery(t *testing.T) {
	type test struct {
		Name     string
		uri      string
		expected [2]string
	}

	tests := []test{
		{
			Name:     "Path only",
			uri:      "/test",
			expected: [2]string{"/test", ""},
		},
		{
			Name:     "Query only",
			uri:      "?test=\"QUERY PARAM\"",
			expected: [2]string{"/", "test=\"QUERY PARAM\""},
		},
		{
			Name:     "Path and query",
			uri:      "/test?testparam=\"QUERY PARAM\"",
			expected: [2]string{"/test", "testparam=\"QUERY PARAM\""},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(tt *testing.T) {
			path, query := splitUriQuery(test.uri)
			if path != test.expected[0] || query != test.expected[1] {
				t.Errorf(
					"expected [%s, %s] but got [%s, %s]",
					test.expected[0],
					test.expected[1],
					path,
					query,
				)
			}
		})
	}
}
