package models

type ToDo struct {
	Id        int    `json:"id"`
	Content   string `json:"content"`
	IsChecked bool   `json:"is_checked"`
	IsDeleted bool   `json:"is_deleted"`
}
