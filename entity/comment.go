package entity



type Comment struct {
	ID int64 `json:"id"`
	Email string `json:"email"`
	Comment string `json:"comment"`
}