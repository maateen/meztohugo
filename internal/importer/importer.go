package importer

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/maateen/meztohugo/config"
)

var (
	db  *gorm.DB
	err error
)

// Post ...
type Post struct {
	Title         string
	Slug          string
	Category      []string
	Description   string
	Created       time.Time
	Updated       time.Time
	Status        uint
	PublishDate   time.Time
	ExpiryDate    time.Time
	Content       string
	FeaturedImage string
}

// TableName ...
func (BlogCategory) TableName() string {
	return "blog_blogcategory"
}

// TableName ...
func (BlogPost) TableName() string {
	return "blog_blogpost"
}

// TableName ...
func (BlogPostCategory) TableName() string {
	return "blog_blogpost_categories"
}

// ImportBlogPosts exports
func ImportBlogPosts() []Post {
	connectToDB()
	defer db.Close()

	posts := []Post{}

	blogposts := []BlogPost{}
	db.Find(&blogposts)

	for _, blogpost := range blogposts {
		categories := []string{}

		blogpostcategories := []BlogPostCategory{}
		db.Where("blogpost_id = ?", blogpost.ID).Find(&blogpostcategories)

		for _, blogpostcategory := range blogpostcategories {
			blogcategories := []BlogCategory{}
			db.Where("id = ?", blogpostcategory.BlogCategory).Find(&blogcategories)

			for _, blogcategory := range blogcategories {
				categories = append(categories, blogcategory.Title)
			}
		}
		posts = append(posts, Post{Title: blogpost.Title, Slug: blogpost.Slug, Category: categories, Description: blogpost.Description, Created: blogpost.Created, Updated: blogpost.Updated, Status: blogpost.Status, PublishDate: blogpost.PublishDate, ExpiryDate: blogpost.ExpiryDate, Content: blogpost.Content, FeaturedImage: blogpost.FeaturedImage})
	}

	return posts
}

// connectToDB ...
func connectToDB() {
	conf := config.GetConfig()
	dbURI := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.Database.Username, conf.Database.Password, conf.Database.Hostname, conf.Database.DBName)
	db, err = gorm.Open(conf.Database.Type, dbURI)

	if err != nil {
		panic(err)
	}
}
