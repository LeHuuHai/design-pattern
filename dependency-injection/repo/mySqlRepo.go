package db

import "design-pattern/dependency-injection/model"

type MySQLRepo struct{}

func (r *MySQLRepo) Save(item model.Item) {}
