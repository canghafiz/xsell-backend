package category

import "be/models/domains"

type Resource struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	CategorySlug string `json:"category_slug"`
	Description  string `json:"description"`
	Icon         string `json:"icon"`
}

func ToResource(category domains.Categories) Resource {
	return Resource{
		CategoryId:   category.CategoryId,
		CategoryName: category.CategoryName,
		CategorySlug: category.CategorySlug,
		Description:  category.Description,
		Icon:         category.Icon,
	}
}

func ToResources(categories []domains.Categories) []Resource {
	var resources []Resource
	for _, category := range categories {
		resources = append(resources, ToResource(category))
	}
	return resources
}
