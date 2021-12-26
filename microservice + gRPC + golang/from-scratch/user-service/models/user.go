package models

type Tabler interface {
	TableName() string
}

type User struct {
	ID    int    `json:"id" gorm:"column:id" desc:"Identify"`
	Token string `json:"name" gorm:"column:name"`
}

type UserRepository interface {
	FindByToken(token string) (*User, error)
}
