package atenea

import "errors"

var (
	ErrCourseNotFound       = errors.New("course not found")
	ErrDataExtractionFailed = errors.New("failed to parse course data")
	ErrInternal             = errors.New("an internal error occurred")
	ErrUnknown              = errors.New("unknown error occurred")
)
