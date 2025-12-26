package util

import "github.com/tidwall/gjson"

// ExtractTextContent extracts text from a content part, handling both flat and nested formats.
// Flat format: {"type": "text", "text": "content"}
// Nested format: {"type": "text", "text": {"text": "content"}} (from SDKs echoing reasoning_content)
func ExtractTextContent(textResult gjson.Result) string {
	if textResult.IsObject() {
		// Nested format - extract inner text
		return textResult.Get("text").String()
	}
	// Normal format - text is a string
	return textResult.String()
}
