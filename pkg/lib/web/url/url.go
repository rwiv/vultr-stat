package url

import (
	"fmt"
)

func GetUrlQs(endpoint string, pathString string, queryString string) string {
	return fmt.Sprintf("%s%s%s", endpoint, getPathString(pathString), getQueryString(queryString))
}

func GetUrl(endpoint string, pathString string) string {
	return fmt.Sprintf("%s%s", endpoint, getPathString(pathString))
}

func getPathString(pathString string) string {
	if pathString[0] == '/' {
		return pathString
	} else {
		return "/" + pathString
	}
}

func getQueryString(queryString string) string {
	if queryString[0] == '?' {
		return queryString
	} else {
		return "?" + queryString
	}
}
