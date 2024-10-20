package model

type MessageReceived struct {
	Text string `json:"text"`
}

type MessageSent struct {
	ID          ID     `json:"id"`
	ChatID      ID     `json:"chatId"`
	SenderID    ID     `json:"senderId"`
	UserName    string `json:"userName"`
	SendingTime string `json:"sendingTime"`
	Text        string `json:"text"`
}
