package models

type UserRead struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	UserId   uint   `json:"userid"`
	BookId   string `json:"bookid"`
	Password string `json:"password"`
}
