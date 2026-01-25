package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name  string
		input http.Header
		want  string
		err   error
	}

	tests := []test{
		{
			name: "simple",
			input: http.Header{
				"Authorization": []string{"ApiKey 777"},
			},
			want: "777",
			err:  nil,
		},
		{
			name:  "no authorization header",
			input: http.Header{},
			want:  "",
			err:   ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed authorization header",
			input: http.Header{
				"Authorization": []string{"Bearer 777"},
			},
			want: "",
			err:  ErrMalformedAuthHeader,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)

			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error %v, got %v", tc.err, err)
			}

			if tc.want != got {
				t.Fatalf("expected %q, got %q", tc.want, got)
			}
		})
	}
}
