package proxy

type RpcClient struct {
}

func NewRpcClient() *RpcClient {
	return &RpcClient{}
}

func (rc *RpcClient) Request() string {
	return "rpc response"
}

type RpcClientProxy struct {
	rc IRpcClient
}

func NewRpcClientProxy() *RpcClientProxy {
	return &RpcClientProxy{
		rc: NewRpcClient(),
	}
}

func (p *RpcClientProxy) Request() string {
	// 鉴权、加密参数
	println("鉴权、加密参数中")
	resp := p.rc.Request()
	// 打印日志
	println("打印日志中")
	return resp
}
