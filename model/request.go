package model

import "io"

// PostRequest request interface for post
type PostRequest interface {
	Url() string
	Encode() []byte
}

// GetRequest request interface for get
type GetRequest interface {
	Url() string
	Encode() string
}

// UploadField multipart/form-data post request field struct
type UploadField struct {
	// Reader upload file reader
	Reader io.Reader
	// Key field key
	Key string
	// Value field value
	Value string
}

// UploadRequest multipart/form-data reqeust interface
type UploadRequest interface {
	Url() string
	// Encode encode request to UploadFields
	Encode() []UploadField
}
