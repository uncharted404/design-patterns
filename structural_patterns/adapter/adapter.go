package adapter

type HttpClient struct {
}

func (c *HttpClient) HttpRequest() {
	println("发送http请求")
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

type HttpClientAdapter struct {
	hc IHttpClient
}

func (a *HttpClientAdapter) Request() {
	a.hc.HttpRequest()
}

func NewHttpClientAdapter(hc *HttpClient) *HttpClientAdapter {
	return &HttpClientAdapter{
		hc: hc,
	}
}
