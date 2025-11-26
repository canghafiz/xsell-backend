package member

import "be/models/resources"

type PageLayoutServMember interface {
	GetHomeLayouts(overrideLimit int) (resources.PageResource, error)
	GetProductDetailLayouts(overrideLimit int) (resources.PageResource, error)
}
