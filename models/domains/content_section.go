package domains

import "time"

type ContentSection struct {
	SectionID   int           `gorm:"primaryKey;column:section_id;autoIncrement"`
	SectionKey  string        `gorm:"column:section_key;not null;uniqueIndex"`
	Title       string        `gorm:"column:title;not null"`
	Subtitle    string        `gorm:"column:subtitle"`
	SectionType string        `gorm:"column:section_type;not null"` // dynamic, fixed, predefined
	Config      ContentConfig `gorm:"column:config;type:jsonb;not null"`
	IsActive    bool          `gorm:"column:is_active;default:true"`
	SortOrder   int           `gorm:"column:sort_order;default:0"`
	CreatedAt   time.Time     `gorm:"column:created_at"`
	UpdatedAt   time.Time     `gorm:"column:updated_at"`
}

type ContentConfig struct {
	Limit      int                    `json:"limit,omitempty"`
	Filters    map[string]interface{} `json:"filters,omitempty"`
	ProductIDs []int                  `json:"product_ids,omitempty"`
	Strategy   string                 `json:"strategy,omitempty"`
}

func (ContentSection) TableName() string {
	return "content_sections"
}
