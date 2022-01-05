package models


type WebsiteSoftware struct {
	ID                   int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Title                string `gorm:"column:title" db:"title" json:"title" form:"title"`
	CategoryId           int64  `gorm:"column:category_id" db:"category_id" json:"category_id" form:"category_id"`
	SeoId                int64  `gorm:"column:seo_id" db:"seo_id" json:"seo_id" form:"seo_id"`
	Status               int64  `gorm:"column:status" db:"status" json:"status" form:"status"`
	//	//Displayorder         int64  `gorm:"column:displayorder" db:"displayorder" json:"displayorder" form:"displayorder"`
	//	CreateTime           string `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	//	UpdateTime           string `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`
	//	DeleteTime           string `gorm:"column:delete_time" db:"delete_time" json:"delete_time" form:"delete_time"`
	//	SoftwarePlatform     string `gorm:"column:software_platform" db:"software_platform" json:"software_platform" form:"software_platform"`
	//	SoftwareLanguage     string `gorm:"column:software_language" db:"software_language" json:"software_language" form:"software_language"`
	//	SoftwareManufacture  string `gorm:"column:software_manufacture" db:"software_manufacture" json:"software_manufacture" form:"software_manufacture"`
	//	OfficialHomepage     string `gorm:"column:official_homepage" db:"official_homepage" json:"official_homepage" form:"official_homepage"`
	//	SoftwareStar         int64  `gorm:"column:software_star" db:"software_star" json:"software_star" form:"software_star"`
	//	IsDeleted            int64  `gorm:"column:is_deleted" db:"is_deleted" json:"is_deleted" form:"is_deleted"`
	//	IsRecommend          int64  `gorm:"column:is_recommend" db:"is_recommend" json:"is_recommend" form:"is_recommend"`
	//	IsPopular            int64  `gorm:"column:is_popular" db:"is_popular" json:"is_popular" form:"is_popular"`
	//	//RecommendReson       string `gorm:"column:recommend_reson" db:"recommend_reson" json:"recommend_reson" form:"recommend_reson"`
	//	SoftwareVersion      string `gorm:"column:software_version" db:"software_version" json:"software_version" form:"software_version"`
	//	SoftwarePlatformType int64  `gorm:"column:software_platform_type" db:"software_platform_type" json:"software_platform_type" form:"software_platform_type"`
	//	SoftwareShowSize     string `gorm:"column:software_show_size" db:"software_show_size" json:"software_show_size" form:"software_show_size"`
	//	SoftwareLogo         string `gorm:"column:software_logo" db:"software_logo" json:"software_logo" form:"software_logo"`
}

type WebsiteSoftwareContent struct {
	ID         int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	SoftwareId int64  `gorm:"column:software_id" db:"software_id" json:"software_id" form:"software_id"`
	Status     int64  `gorm:"column:status" db:"status" json:"status" form:"status"`
	CreateTime string `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	UpdateTime string `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`
	DeleteTime string `gorm:"column:delete_time" db:"delete_time" json:"delete_time" form:"delete_time"`
	IsDeleted  int64  `gorm:"column:is_deleted" db:"is_deleted" json:"is_deleted" form:"is_deleted"`
	Content    string `gorm:"column:content" db:"content" json:"content" form:"content"`
}