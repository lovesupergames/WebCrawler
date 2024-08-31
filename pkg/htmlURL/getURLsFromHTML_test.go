package htmlURL

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name        string
		inputBody   string
		expected    []string
		inputURL    string
		expectedErr bool
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute URLs",
			inputURL: "https://example.com",
			inputBody: `
<html>
	<body>
		<a href="https://example.com">
			<span>Example</span>
		</a>
	</body>
</html>
`,
			expected:    []string{"https://example.com"},
			expectedErr: false,
		},
		{
			name:     "relative URLs",
			inputURL: "https://example.com",
			inputBody: `
<html>
	<body>
		<a href="/path/to/page">
			<span>Relative Link</span>
		</a>
	</body>
</html>
`,
			expected:    []string{"https://example.com/path/to/page"},
			expectedErr: false,
		},
		{
			name:     "empty url",
			inputURL: "",
			inputBody: `<html>
	<body>
		<a href="/path/to/page">
			<span>Relative Link</span>
		</a>
	</body>
</html>
`,
			expected:    []string{},
			expectedErr: true,
		},
		{
			name:        "empty body",
			inputURL:    "https://example.com",
			inputBody:   ``,
			expected:    []string{},
			expectedErr: true,
		},
		{
			name:     "nested links",
			inputURL: "https://example.com",
			inputBody: `
<html>
	<body>
		<div>
			<a href="/deep/link">
				<span>Nested Link</span>
			</a>
		</div>
	</body>
</html>
`,
			expected:    []string{"https://example.com/deep/link"},
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tt.inputBody, tt.inputURL)
			if (err != nil) != tt.expectedErr {
				t.Errorf("Test %v FAIL: unexpected error: %v", tt.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Test %v FAIL: expected URLs: %v, actual: %v", tt.name, tt.expected, actual)
			}
		})
	}
}
