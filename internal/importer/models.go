package importer

import "time"

// BlogCategory ...
type BlogCategory struct {
	ID     int    `gorm:"column:id"`
	Title  string `gorm:"column:title"`
	Slug   string `gorm:"column:slug"`
	SiteID int    `gorm:"column:site_id"`
}

// BlogPost ...
type BlogPost struct {
	ID             int       `gorm:"column:id"`
	CommentsCount  int       `gorm:"column:comments_count"`
	KeywordsString string    `gorm:"column:keywords_string"`
	RatingCount    int       `gorm:"column:rating_count"`
	RatingSum      int       `gorm:"column:rating_sum"`
	RatingAverage  float64   `gorm:"column:rating_average"`
	Title          string    `gorm:"column:title"`
	Slug           string    `gorm:"column:slug"`
	MetaTitle      string    `gorm:"column:_meta_title"`
	Description    string    `gorm:"column:description"`
	GenDescription uint      `gorm:"column:gen_description"`
	Created        time.Time `gorm:"column:created"`
	Updated        time.Time `gorm:"column:updated"`
	Status         uint      `gorm:"column:status"`
	PublishDate    time.Time `gorm:"column:publish_date"`
	ExpiryDate     time.Time `gorm:"column:expiry_date"`
	ShortURL       string    `gorm:"column:short_url"`
	InSitemap      uint      `gorm:"column:in_sitemap"`
	Content        string    `gorm:"column:content"`
	AllowComments  uint      `gorm:"column:allow_comments"`
	FeaturedImage  string    `gorm:"column:featured_image"`
	SiteID         int       `gorm:"column:site_id"`
	UserID         int       `gorm:"column:user_id"`
}

// BlogPostCategory ...
type BlogPostCategory struct {
	ID           int `gorm:"column:id"`
	BlogPost     int `gorm:"column:blogpost_id"`
	BlogCategory int `gorm:"column:blogcategory_id"`
}
