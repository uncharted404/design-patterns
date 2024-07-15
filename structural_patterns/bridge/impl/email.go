package impl

type Email struct {
}

func (e Email) Send(msg string) {
	println("email发送：", msg)
}

func NewEmail() *Email {
	return &Email{}
}
