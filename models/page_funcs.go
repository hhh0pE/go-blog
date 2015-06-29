package models

func GetPageByUrl(url_to_find string) (*Page, bool) {
	p := Page{}
	query := Connection.Where("url = ?", url_to_find).First(&p)

	if query.Error != nil {
		return nil, false
	}
	return &p, true
}

func GetPageByID(id_to_find int) (*Page, bool) {
	p := Page{}
	query := Connection.First(&p, id_to_find)

	if query.Error != nil {
		return nil, false
	}

	return &p, true
}

func AllPostsInCategory(id int) []Page {
	posts := []Page{}
	Connection.Table("pages").Where("parent_id = ?", id).Order("updated_at desc").Find(&posts)
	return posts
}

func OtherPostsInThisCategory(p Page) []Page {
	other_in_cat := []Page{}
	Connection.Table("pages").Where("parent_id = ? and id <> ?", p.ParentID, p.ID).Order("viewed_count DESC").Find(&other_in_cat)
	return other_in_cat
}

func AllCategories() []Page {
	categories := []Page{}
	Connection.Table("pages").Where("template_id = 3").Find(&categories)
	return categories
}

func AllPosts() []Page {
	all_posts := []Page{}
	Connection.Table("pages").Where("template_id = 2").Order("updated_at desc").Order("id").Find(&all_posts)
	return all_posts
}
