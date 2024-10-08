package handler

type AddChat struct {
	Name     string   `json:"name"`
	IsDirect bool     `json:"is_direct"`
	Members  []uint64 `json:"course_id"`
}

type AddContact struct {
	ContactId uint64 `json:"contact_id"`
}

type Sign struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Refresh struct {
	Token string `json:"token"`
}
