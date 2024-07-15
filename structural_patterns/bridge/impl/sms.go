package impl

type Sms struct {
}

func (s *Sms) Send(msg string) {
	println("sms发送：", msg)
}

func NewSms() *Sms {
	return &Sms{}
}
