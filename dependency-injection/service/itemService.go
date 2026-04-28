package service

import (
	"design-pattern/dependency-injection/model"
	db "design-pattern/dependency-injection/repo"
)

type ItemService struct {
	Repo db.IRepo
}

func (s *ItemService) SaveItem(item model.Item) {
	s.Repo.Save(item)
}
