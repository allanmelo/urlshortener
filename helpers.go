package urlshortener

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func yamlParser(yamlBytes []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil

}

func pathMaps(paths []pathUrl) map[string]string {
	pathMap := make(map[string]string)
	for _, ref := range paths {
		pathMap[ref.Path] = ref.URL
	}
	return pathMap
}

func ReadFile(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
	return file, nil
}