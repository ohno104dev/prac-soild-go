package bad

type EmailSender struct{}

func (e *EmailSender) SendEmail(address, message string) {
	// Logic to send an email
}

type SMSSender struct{}

func (s *SMSSender) SendSMS(number, message string) {
	// Logic to send an SMS
}

type Notification struct{}

// Notification struct depends directly on the concrete implementations of EmailSender and SMSSender
// If we need to add a new communication channel, we would have to change the Notification struct
func (n *Notification) Send(address, message string, sender string) {
	if sender == "email" {
		emailSender := EmailSender{}
		emailSender.SendEmail(address, message)
	} else if sender == "sms" {
		smsSender := SMSSender{}
		smsSender.SendSMS(address, message)
	}
}
