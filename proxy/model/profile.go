package model

import db "design-pattern/proxy/config/database"

type Profile struct {
	ID   uint
	Data []byte
}

func (p *Profile) Load() {
	db.DB.Select("profile").
		Where("id = ?", p.ID).
		Find(&p.Data)
}

func (p *Profile) Get() []byte {
	return p.Data
}
