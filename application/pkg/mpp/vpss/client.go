package vpss

type Client interface {
    Name() string
    ReceiveData(Frame)
}
