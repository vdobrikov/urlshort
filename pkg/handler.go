package urlshort

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		url, ok := pathsToUrls[request.URL.Path]
		if ok {
			http.Redirect(writer, request, url, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(writer, request)
		}
	}
}

type Url struct {
	Path string `json:"path" yaml:"path"` // Supporting both JSON and YAML.
	Url string `json:"url" yaml:"url"`
}

func YAMLHandler(filename string, fallback http.Handler) (http.HandlerFunc, error) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Failed to read file")
		return nil, err
	}

	var urls []Url
	err = yaml.Unmarshal(buffer, &urls)
	if err != nil {
		fmt.Println("Failed to read file")
		return nil, err
	}

	pathsToUrls := make(map[string]string)
	for _, url := range urls {
		pathsToUrls[url.Path] = url.Url
	}

	return MapHandler(pathsToUrls, fallback), nil
}
