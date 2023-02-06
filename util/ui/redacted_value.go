package ui

import (
	"net/http"
)

// RedactedValue is the text that is displayed for redacted content. (eg
// authorization tokens, passwords, etc.)

func RedactHeaders(header http.Header) http.Header {
	redactedHeaders := make(http.Header)

	return redactedHeaders
}
