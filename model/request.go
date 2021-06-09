package model

type PostRequest interface {
	Url() string
	Encode() []byte
}

type GetRequest interface {
	Url() string
	Encode() string
}
