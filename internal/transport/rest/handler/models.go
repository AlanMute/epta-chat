package handler

type AddChat struct {
	Name     string   `json:"name"`
	IsDirect bool     `json:"is_direct"`
	Members  []uint64 `json:"members_ids"`
}

type AddMember struct {
	ChatId  uint64   `json:"chat_id"`
	Members []uint64 `json:"members_ids"`
}

type AddContact struct {
	ContactLogin string `json:"contact_login"`
}

type Sign struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Refresh struct {
	UserId uint64 `json:"user_id"`
	Token  string `json:"token"`
}

type UserName struct {
	Username string `json:"username"`
}

type RefreshRes struct {
	AccessToken string `json:"access_token"`
}

type ChatIdResponse struct {
	ChatId uint64 `json:"chat_id"`
}
