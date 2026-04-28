package db

import "design-pattern/dependency-injection/model"

type PostgresRepo struct{}

func (r *PostgresRepo) Save(item model.Item) {}
