package utils

import (
	"fmt"
	"net/url"
)

type UrlBuilder struct {
	basePath string
	values   url.Values
}

func NewUrlBuilder(basePath string) *UrlBuilder {
	if basePath == "" {
		return nil
	}
	return &UrlBuilder{
		basePath: basePath,
		values:   url.Values{},
	}
}

func (builder *UrlBuilder) Add(key string, value string) *UrlBuilder {
	if builder.values == nil {
		builder.values = url.Values{}
	}
	// if key is empty, will ignore it
	if key == "" {
		return builder
	}
	builder.values.Add(key, value)
	return builder
}

func (builder *UrlBuilder) Build() string {
	if builder.values == nil || len(builder.values) <= 0 {
		if builder.basePath != "" {
			return builder.basePath
		}
		return ""
	}
	if builder.basePath == "" {
		return builder.values.Encode()
	}
	return fmt.Sprintf("%s?%s", builder.basePath, builder.values.Encode())
}
