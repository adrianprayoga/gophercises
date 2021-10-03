package urlshort

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirectPath, ok := pathsToUrls[r.URL.Path]
		if ok {
			w.Header().Add("Location", redirectPath)
			w.WriteHeader(301)
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// m := make([]map[string]string, 10)
	var urls []Redirector
	err := yaml.Unmarshal(yml, &urls)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return MapHandler(convertToMap(urls), fallback), nil
	// return MapHandler(ymlToMap(urls), fallback), err
}

func JSONHandler(jsonByte []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var urls []Redirector
	err := json.Unmarshal(jsonByte, &urls)
	if err != nil {
		return nil, err
	}
	return MapHandler(convertToMap(urls), fallback), err
}

type Redirector struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func convertToMap(yml []Redirector) map[string]string {
	redirectMap := make(map[string]string)

	for _, s := range yml {
		redirectMap[s.Path] = s.Url
	}

	fmt.Println(redirectMap)
	return redirectMap
}
