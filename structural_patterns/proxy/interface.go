package proxy

type IRpcClient interface {
	Request() string
}
