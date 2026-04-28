package db

import "design-pattern/dependency-injection/model"

type IRepo interface {
	Save(item model.Item)
}
