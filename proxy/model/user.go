package model

type User struct {
	ID      uint `gorm:"primarykey"`
	Profile IProfile
}

func NewUser(id uint) *User {
	return &User{
		ID:      id,
		Profile: &ProxyProfile{ID: id},
	}
}
