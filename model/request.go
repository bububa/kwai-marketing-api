package model

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
