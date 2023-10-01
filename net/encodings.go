// Copyright 2023 Christopher Briscoe.  All rights reserved.
package net

import (
	"net/http"
	"strings"
)

// SetPreferredEncoding will default to using BR compression if the client accepts
// that encoding.  Otherwise, use GZIP.
func SetPreferredEncoding(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Accept-Encoding")
	encodings := strings.Split(header, ", ")
	br := false
	for _, s := range encodings {
		if s == "br" {
			br = true
			break
		}
	}
	if br {
		w.Header().Add("Content-Encoding", "br")
		return
	}
	w.Header().Add("Content-Encoding", "gzip")
}

// GetRequestParams splits off the preferrend encoding from the request keys
func GetRequestParams(key string) ([]string, string) {
	var encoding string
	keys := strings.Split(key, "|")

	if len(keys) == 0 {
		return []string{""}, ""
	}

	last := keys[len(keys)-1]
	if last == "br" || last == "gz" {
		encoding = last
		keys = keys[:len(keys)-1]
	}

	if len(keys) == 0 {
		keys = []string{""}
	}

	return keys, encoding
}
