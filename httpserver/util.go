package httpserver

import "strings"

func splitUriQuery(uri string) (path string, query string) {
	uriQuery := strings.Split(uri, "?")

	if uriQuery[0] == "" {
		path = "/"
	} else {
		path = uriQuery[0]
	}

	if len(uriQuery) > 1 {
		query = strings.Join(uriQuery[1:], "?")
	}

	return path, query
}
