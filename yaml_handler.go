package main

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}

	paths := make(map[string]string)
	for _, pu := range pathUrls {
		paths[pu.Path] = pu.URL
	}

	return MapHandler(paths, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
