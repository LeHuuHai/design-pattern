package main

import (
	"design-pattern/dependency-injection/model"
	db "design-pattern/dependency-injection/repo"
	"design-pattern/dependency-injection/service"
)

func main() {
	MySQLRepo := db.MySQLRepo{}
	ItemService := service.ItemService{Repo: &MySQLRepo}
	ItemService.SaveItem(model.Item{})
}
