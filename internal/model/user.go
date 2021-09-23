package model

type User struct{
	ID uint32	`gorm:"primary_key" json:"id"`
	Name string	`json:"name"`
	Password string	`json:"password"`
	Privilege int	`json:"privilege"`
}

func (user *User)TableName() string{
	return "user"
}

func (u *User)