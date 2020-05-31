package common

type Encoder interface {
	GetId() int
	DataCallback(data []byte)
}
