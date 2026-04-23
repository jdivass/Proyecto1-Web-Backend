package utils

import "net/http"

func BuildImageURL(r *http.Request, path string) string {
	if path == "" {
		return ""
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	return scheme + "://" + r.Host + "/" + path
}