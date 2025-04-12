package atenea

import "errors"

var (
	errCourseNotFound       = errors.New("course not found")
	errDataExtractionFailed = errors.New("failed to parse course data")
	errInternal             = errors.New("an internal error occurred")
	errUnknown              = errors.New("unknown error occurred")
)
