package main

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

var allCategories []Category
var categoryMap map[int]Category

func loadCategories() {
	err := dbx.Select(&allCategories, "SELECT * FROM categories")
	if err != nil {
		log.Fatal(err)
	}

	categoryMap = make(map[int]Category, len(allCategories))

	for _, category := range allCategories {
		if category.ParentID > 0 {
			category.ParentCategoryName = categoryMap[category.ParentID].CategoryName
		}
		categoryMap[category.ID] = category
	}
}

func getCategoryByID(q sqlx.Queryer, categoryID int) (category Category, err error) {
	category, ok := categoryMap[categoryID]
	if !ok {
		return Category{}, sql.ErrNoRows
	}
	return category, nil
}
