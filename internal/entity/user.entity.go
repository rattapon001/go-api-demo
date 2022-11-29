package entity

type User struct {
	Id        uint   `gorm:"primary_key" json:"id" gorm:"autoIncrement"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}
