package models

type Category struct {
	ID          int
	Title, Slug string
}

var Categories map[int]*Category

func init() {
	Categories = make(map[int]*Category)
}

func (c *Category) GetAll() []Category {
	cats := []Category{}
	DB.Find(&cats)

	for id, cat := range cats {
		Categories[id] = &cat
	}
	return cats
}

func GetCategoryByID(id int) *Category {
	cat := Category{}
	query := DB.Where("id = ?", id).First(&cat)

	if query.RowsAffected == 0 {
		return nil
	}

	return &cat
}
