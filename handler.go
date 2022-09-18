package urlshortener

import (
	"net/http"
)


func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		} else {
			fallback.ServeHTTP(w, r)
			return
		}
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := yamlParser(yamlBytes)
	if err != nil {
		return nil, err
	}

	pathMap := pathMaps(pathUrls)
	return MapHandler(pathMap, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
