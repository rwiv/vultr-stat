package cutil

func GetHeaders() map[string]string {
	return GetHeadersContentType("application/json")
}

func GetHeadersContentType(contentType string) map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = contentType

	return headers
}
