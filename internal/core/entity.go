package core

import "time"

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type UserInfo struct {
	ID       uint64
	Login    string
	UserName string
}

type MessageInfo struct {
	ID          uint64
	Text        string
	SendingTime time.Time
	SenderId    uint64
	UserName    string
	ChatId      uint64
}
