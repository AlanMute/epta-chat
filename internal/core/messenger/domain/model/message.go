package model

type MessageReceived struct {
	Text string `json:"text"`
}

type MessageSent struct {
	AuthorID int    `json:"author_id"`
	Text     string `json:"text"`
}
