package handler

type AddChat struct {
	Name     string   `json:"name"`
	IsDirect bool     `json:"is_direct"`
	Members  []uint64 `json:"course_id"`
}
