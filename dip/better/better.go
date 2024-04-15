package better

// Dependency Inversion Principle (DIP)
// 應該依賴於抽象,而不是具體實現

// Advantages
// 靈活性, 通過依賴抽象而不是具體實現, 方便修改或擴展組件, 而不影響其他代碼
// 可測試性, 通過依賴抽象, 可以在測試期間使用模擬實現或測試替身, 將正在測試的組件與系統的其餘部分隔離

// Disadvantages
// 複雜性增加, 依賴抽象會引入額外的間接層, 增加開發人員理解和導航抽象的複雜性
// 過早抽象, 可能導致過度工程, 並非所有系統部分都需要解耦合

type MessageSender interface {
	SendMessage(address, message string)
}

type EmailSender struct{}

func (e *EmailSender) SendMessage(address, message string) {
	// Logic to send an email
}

type SMSSender struct{}

func (s *SMSSender) SendMessage(number, message string) {
	// Logic to send an SMS
}

// If we need to add a new communication channel, we only need to create a new struct that implements the MessageSender interface

type Notification struct {
	send MessageSender
}

func (n *Notification) Send(address, message string) {
	n.send.SendMessage(address, message)
}
