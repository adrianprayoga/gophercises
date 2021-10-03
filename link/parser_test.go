package parser

import (
	"os"
	"reflect"
	"testing"
)

func TestParseHtml(t *testing.T) {
	var tests = []struct {
		html_filename string
		expected      []Link
	}{
		{"htmls/ex1.html", []Link{Link{"/other-page", "paragrah 1 A link to another page"}}},
		{"htmls/ex2.html", []Link{
			Link{"https://www.twitter.com/joncalhoun", "Check me out on twitter"},
			Link{"https://github.com/gophercises", "Gophercises is on Github!"},
		}},
		{"htmls/ex3.html", []Link{
			Link{"#", "Login"},
			Link{"/lost", "Lost? Need help?"},
			Link{"https://twitter.com/marcusolsson", "@marcusolsson"},
		}},
		{"htmls/ex4.html", []Link{Link{"/dog-cat", "dog cat"}}},
	}

	for _, tt := range tests {
		t.Run(tt.html_filename, func(t *testing.T) {
			r, _ := os.Open(tt.html_filename)
			defer r.Close()

			ans, _ := GetLinks(r)
			if !reflect.DeepEqual(ans, tt.expected) {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}
