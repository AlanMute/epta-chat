package core

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type UserInfo struct {
	ID       uint64
	Login    string
	UserName string
}
