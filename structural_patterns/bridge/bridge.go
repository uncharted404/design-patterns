package bridge

import "design-patterns/structural_patterns/bridge/impl"

type Notification struct {
	sender Sender
}

func (n *Notification) Notice(msg string) {
	n.sender.Send(msg)
}

type SmsNotification struct {
	Notification
}

func NewSmsNotification() *SmsNotification {
	return &SmsNotification{
		Notification{
			sender: impl.NewSms(),
		}}
}

type EmailNotification struct {
	Notification
}

func NewEmailNotification() *EmailNotification {
	return &EmailNotification{
		Notification{
			sender: impl.NewEmail(),
		}}
}
