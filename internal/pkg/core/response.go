// Copyright (c) 2026 fishfinal
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package core

// PaginationBaseRequest defines the common pagination parameters
// used for listing endpoints that support page-based navigation.
type PaginationBaseRequest struct {
	// Number specifies the page number to retrieve.
	// Page numbering starts at 1. Defaults to 1 if not provided.
	Number int `json:"number" form:"number,default=1"`

	// Size specifies the number of items per page.
	// Defaults to 10 if not provided.
	Size int `json:"size" form:"size,default=10"`
}

// PaginationResponse wraps paginated data along with the pagination
// parameters used for the request. This allows clients to know
// which page and size were applied to the returned data.
type PaginationResponse struct {
	PaginationBaseRequest     // Embedded pagination request parameters
	Data                  any `json:"data"` // The actual paginated data payload
}

// Response defines the unified API response format used across all endpoints.
// It provides a consistent structure for success and error responses,
// including a business status code, a human-readable message, and optional data.
type Response struct {
	// Message provides a human-readable description of the response,
	// such as "success" for success or an error description.
	Message string `json:"message"`

	// Code is a business status code that indicates the result of the operation.
	// For example, 100100 represents success, while 100101 represents failure.
	Code int `json:"code"`

	// Data contains the response payload. This field is omitted when empty
	// (e.g., for error responses or operations that don't return data).
	Data any `json:"data,omitempty"`
}

// Business status code constants.
const (
	// CodeWithSuccess indicates that the operation completed successfully.
	CodeWithSuccess = 100100

	// CodeWithFailed indicates that the operation failed.
	CodeWithFailed = 100101
)

// Success creates a successful API response with the provided data.
// It sets the status code to CodeWithSuccess (100100) and the message to "success".
// Use this function for endpoints that return data upon successful completion.
func Success(data interface{}) *Response {
	return &Response{
		Code:    CodeWithSuccess,
		Message: "success",
		Data:    data,
	}
}

// Error creates an error API response with the given error message.
// It sets the status code to CodeWithFailed (100101) and leaves Data as nil.
// Use this function for endpoints that need to return error information.
func Error(message string) *Response {
	return &Response{
		Code:    CodeWithFailed,
		Message: message,
	}
}
