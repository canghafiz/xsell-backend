package member

import "be/models/resources"

type PageLayoutServMember interface {
	GetHomeLayouts(sectionLimits map[string]int) (resources.PageResource, error)
	GetProductDetailLayouts(sectionLimits map[string]int) (resources.PageResource, error)
}
