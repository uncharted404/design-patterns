package adapter

type IClient interface {
	Request()
}

type IHttpClient interface {
	HttpRequest()
}
